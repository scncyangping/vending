package main

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	// create config
	config, err := clientcmd.BuildConfigFromFlags("",
		"/Users/yapi/WorkSpace/GolandWorkSpace/study/client-go/config/demo.config")
	if err != nil {
		panic(err)
	}
	// create set
	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}
	// get informer
	factory := informers.NewSharedInformerFactory(clientSet, 0)

	informer := factory.Core().V1().Pods().Informer()

	rateLimitQueue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller")

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add")
			key, er := cache.MetaNamespaceKeyFunc(obj)
			if er != nil {
				fmt.Println("can not get key")
			}
			rateLimitQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete")
		},
	})

	stopCh := make(chan struct{})

	factory.Start(stopCh)

	factory.WaitForCacheSync(stopCh)

	<-stopCh
}
