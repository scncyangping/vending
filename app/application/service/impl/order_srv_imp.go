package impl

import (
	"errors"
	"fmt"
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
	"vending/app/infrastructure/pkg/log"
	"vending/app/types/constants"
)

type OrderSrvImp struct {
}

func NewOrderSrvImp() *OrderSrvImp {
	return &OrderSrvImp{}
}

func (o *OrderSrvImp) CreateOrder(cmd *cmd.CreateOrderCmd) (string, error) {
	commodityIds := make([]string, 0)
	for key, _ := range cmd.Items {
		commodityIds = append(commodityIds, key)
	}
	// 下单流程创建
	// 构建商品数据
	commodityEns, err := factory.Instance.CommodityAggregate.BuildCommodityObj(commodityIds)
	if err != nil {
		return constants.EmptyStr, err
	}
	// 构建预处理订单数据
	var orderId string
	var amount float64

	if items, err := factory.Instance.OrderAggregate.BuildOrderItem(cmd.Items, commodityEns); err != nil {
		return constants.EmptyStr, err
	} else {
		if o, a, err := factory.Instance.OrderAggregate.CreateTempOrderOne(items, cmd.PayDes); err != nil {
			return constants.EmptyStr, err
		} else {
			orderId = o
			amount = a
		}
	}

	// 库存校验及预锁定
	for _, v := range commodityEns {
		if ig, err := factory.Instance.InventoryAggregate.Instance(v.CategoryId); err != nil {
			return constants.EmptyStr, err
		} else {
			// 库存校验
			if ig.OutOfStock() {
				errMsg := fmt.Sprintf("商品: {%s} 库存不足,请联系客服", v.Name)
				log.Logger().Errorf("{%s} , 订单信息: {%v}", errMsg, cmd)
				return constants.EmptyStr, errors.New(errMsg)
			}
			// 预锁定
			if _, err := ig.OutStock(orderId, cmd.Items[v.Id]); err != nil {
				return constants.EmptyStr, err
			}
		}
	}

	// 订单Id及金额生成支付Url
	if payUrl, err := factory.Instance.PayAggregate.Pay(orderId, amount); err != nil {
		return constants.EmptyStr, err
	} else {
		return payUrl, nil
	}
}

func (o *OrderSrvImp) OrderCallBack(orderId string) error {
	if og, err := factory.Instance.OrderAggregateInstance(orderId); err != nil {
		return err
	} else {
		return og.SaveOrder(orderId)
	}
}

func (o *OrderSrvImp) Cancel(orderId string) error {
	return factory.Instance.OrderAggregate.Cancel(orderId)
}

func (o *OrderSrvImp) Get(s string) (*dto.OrderDto, error) {
	return nil, nil
}

func (o *OrderSrvImp) Query(query query.OrderPageQuery) ([]*dto.OrderListDto, error) {
	return nil, nil
}
