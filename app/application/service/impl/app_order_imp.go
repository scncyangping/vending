package impl

import (
	"errors"
	"fmt"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/service"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
	"vending/app/types/constants"
)

// 流程编排
// 不要处理业务逻辑
// 参数验证、构建、错误处理、监控日志、事务处理、认证与授权

type OrderSrvImp struct {
	factory            *factory.AgFactory
	doCommoditySrv     service.DoCommodityService
	doInventoryService service.DoInventoryService
	doOrderService     service.DoOrderService
	doPayService       service.DoPayService
}

func NewOrderSrvImp(factory *factory.AgFactory, doCommoditySrv service.DoCommodityService,
	doInventoryService service.DoInventoryService,
	doOrderService service.DoOrderService) *OrderSrvImp {
	return &OrderSrvImp{
		doCommoditySrv:     doCommoditySrv,
		doInventoryService: doInventoryService,
		doOrderService:     doOrderService,
		factory:            factory}
}

func (o *OrderSrvImp) CreateOrder(cmd *cmd.CreateOrderCmd) (string, error) {
	// 构建预处理订单数据
	var (
		orderId string
		amount  float64
		result  = constants.EmptyStr
	)
	// 下单流程创建
	// step1. 构建商品数据
	commodityEns, err := o.doCommoditySrv.QueryCommoditiesByIds(cmd.GetCommodityIds())
	if err != nil {
		return result, err
	}
	// step2. 构建订单明细
	if items, err := o.buildOrderItems(cmd.Items, commodityEns); err != nil {
		return constants.EmptyStr, err
	} else {
		orderId, amount, err = o.doOrderService.CreateTempOrderOne(items, cmd.PayDes)
		if err != nil {
			return result, err
		}
	}
	// step3. 库存校验及预锁定
	err = o.preOutStock(cmd.Items, commodityEns, orderId)
	if err != nil {
		return result, err
	}

	// step4. 支付Url
	if payUrl, err := o.doPayService.PayUrl(orderId, amount); err != nil {
		return constants.EmptyStr, err
	} else {
		return payUrl, nil
	}
}

func (o *OrderSrvImp) OrderCallBack(orderId string) error {
	return o.doOrderService.SaveOrder(orderId)
}

func (o *OrderSrvImp) Cancel(orderId string) error {
	if ig, err := o.factory.OrderAggregateInstance(orderId); err != nil {
		return err
	} else {
		return ig.Cancel()
	}
}

func (o *OrderSrvImp) GetTempOrderById(s string) (*dto.OrderDto, error) {
	var (
		dto *dto.OrderDto
	)
	if do, err := o.doOrderService.GetTempOrderById(s); err != nil {
		return nil, err
	} else {
		util.StructCopy(dto, do)
		return dto, nil
	}
}

func (o *OrderSrvImp) GetOrderById(s string) (*dto.OrderDto, error) {
	var (
		dto *dto.OrderDto
	)
	if do, err := o.doOrderService.GetOrderById(s); err != nil {
		return nil, err
	} else {
		util.StructCopy(dto, do)
		return dto, nil
	}
}

func (o *OrderSrvImp) Query(query query.OrderPageQuery) ([]*dto.OrderListDto, error) {
	qb := query.PageQuery.QBase()

	filter := types.B{}

	if query.CommodityName != constants.EmptyStr {
		filter["items.commodityName"] = types.B{"$reg": query.CommodityName}
	}
	if query.CommodityId != constants.EmptyStr {
		filter["items.commodityId"] = query.CommodityId
	}
	if query.PaymentPhone != constants.EmptyStr {
		filter["payment.phone"] = query.PaymentPhone
	}

	if query.PaymentEmail != constants.EmptyStr {
		filter["payment.email"] = query.PaymentEmail
	}

	if dos, err := o.doOrderService.QueryOrderPageBy(qb.Skip, qb.Limit, qb.Sort, filter); err != nil {
		return nil, err
	} else {
		var dtoList []*dto.OrderListDto
		for _, v := range dos {
			dto := dto.OrderListDto{}
			util.StructCopy(dto, v)
			dtoList = append(dtoList, &dto)
		}
		return dtoList, nil
	}
}

func (o *OrderSrvImp) preOutStock(m map[string]int, commodityEns []*entity.CommodityEn, orderId string) error {
	for _, v := range commodityEns {
		if ig, err := o.factory.InventoryAggregateInstance(v.CategoryId); err != nil {
			return err
		} else {
			// 库存校验
			if ig.OutOfStock() {
				errMsg := fmt.Sprintf("商品: {%s} 库存不足,请联系客服", v.Name)
				log.Logger().Errorf("{%s} , 订单信息: {%v}", errMsg, m)
				return errors.New(errMsg)
			}
			// 预锁定
			if _, err := ig.OutStock(orderId, m[v.Id]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *OrderSrvImp) buildOrderItems(m map[string]int, commodityEns []*entity.CommodityEn) ([]obj.OrderItemObj, error) {
	objs := make([]obj.OrderItemObj, 0)
	for _, v := range commodityEns {
		objs = append(objs, o.buildOrderItem(v, m[v.Id]))
	}
	return objs, nil
}

func (o *OrderSrvImp) buildOrderItem(commodityEn *entity.CommodityEn, num int) obj.OrderItemObj {
	var (
		item obj.OrderItemObj
	)
	// 商品聚合
	item.CommodityId = commodityEn.Id
	item.Amount = commodityEn.Amount * float64(num)
	item.CommodityName = commodityEn.Name
	item.OriginalAmount = commodityEn.Amount * float64(num) // 原金额 TOTO 可加优惠策略
	item.OwnerId = commodityEn.OwnerId
	return item
}
