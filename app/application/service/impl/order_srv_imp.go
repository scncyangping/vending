package impl

import (
	"vending/app/application/cqe/cmd"
	"vending/app/application/cqe/query"
	"vending/app/application/dto"
	"vending/app/domain/aggregate/factory"
)

type OrderSrvImp struct {
}

func NewOrderSrvImp() *OrderSrvImp {
	return &OrderSrvImp{}
}

func (o *OrderSrvImp) CreateOrder(cmd *cmd.CreateOrderCmd) (string, error) {
	return factory.Instance.OrderAggregate.CreateTempOrderOne(cmd)
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
