package pkg

import (
	"context"
	v17 "k8s.io/api/core/v1"
	v15 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v16 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	v14 "k8s.io/client-go/informers/core/v1"
	v13 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	v12 "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"time"
)

const (
	workerNum = 5
	maxRetry  = 10
)

type controller struct {
	client        kubernetes.Interface
	ingressLister v1.IngressLister
	serviceLister v12.ServiceLister
	queue         workqueue.RateLimitingInterface
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	// todo 比较annotation是否相同
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}

func (c *controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c *controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}

	c.queue.Add(key)
}
func (c *controller) deleteIngressFun(obj interface{}) {
	ingress := obj.(*v15.Ingress)
	service := v16.GetControllerOf(ingress)
	if service == nil {
		return
	}
	if service.Kind != "Service" {
		return
	}
	c.queue.Add(ingress.Namespace + "/" + ingress.Name)
}

func (c *controller) Run(stop chan struct{}) {
	for i := 0; i < workerNum; i++ {
		go wait.Until(c.worker, time.Minute, stop)
	}
	<-stop
}

func (c *controller) worker() {
	for c.processNextItem() {

	}
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)

	key := item.(string)

	err := c.syncService(key)

	if err != nil {
		c.handlerError(key, err)
	}

	return true
}

func (c *controller) syncService(key string) error {
	namespaceKey, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return nil
	}
	service, err := c.serviceLister.Services(namespaceKey).Get(name)
	if errors.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	_, ok := service.GetAnnotations()["ingress/http"]

	ingress, err := c.ingressLister.Ingresses(namespaceKey).Get(name)

	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if ok && errors.IsNotFound(err) {
		// 此service存在ingress注解，并且没有被创建，需要创建
		// create ingress
		ig := c.constructIngress(service)
		c.client.NetworkingV1().Ingresses(namespaceKey).Create(context.TODO(), ig, v16.CreateOptions{})
	} else if !ok && ingress != nil {
		// 不存在这个注解，并且已经被创建了，需要删除
		err := c.client.NetworkingV1().Ingresses(namespaceKey).Delete(context.TODO(), name, v16.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *controller) handlerError(key string, err error) {
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c *controller) constructIngress(service *v17.Service) *v15.Ingress {
	ingress := v15.Ingress{}

	ingress.ObjectMeta.OwnerReferences = []v16.OwnerReference{
		*v16.NewControllerRef(service, v17.SchemeGroupVersion.WithKind("Service")),
	}

	ingress.Name = service.Name
	ingress.Namespace = service.Namespace
	pathType := v15.PathTypePrefix
	icn := "nginx"
	ingress.Spec = v15.IngressSpec{
		IngressClassName: &icn,
		Rules: []v15.IngressRule{
			{
				Host: "scncys.cn",
				IngressRuleValue: v15.IngressRuleValue{
					HTTP: &v15.HTTPIngressRuleValue{
						Paths: []v15.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &pathType,
								Backend: v15.IngressBackend{
									Service: &v15.IngressServiceBackend{
										Name: service.Name,
										Port: v15.ServiceBackendPort{
											Number: 80,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return &ingress
}

func NewController(client kubernetes.Interface, ingressInformer v13.IngressInformer, serviceInformer v14.ServiceInformer) *controller {
	c := controller{
		client:        client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "IngressManager"),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngressFun,
	})

	return &c
}
