package aggregate

import (
	"errors"
	"vending/app/domain/entity"
	"vending/app/domain/repo"
	"vending/app/infrastructure/pkg/util"
	"vending/app/types"
)

// 创建订单
// 取消订单
// 修改订单
type orderAggregateRepo struct {
	orderRepo       repo.OrderRepo
	orderTempRepo   repo.OrderTempRepo
	beneficiaryRepo repo.BeneficiaryRepo
}

type OrderAggregate struct {
	orderAggregateRepo
	OrderId string
	OrderEn entity.OrderEn // 订单基础数据
}

func NewOrderAggregate(orderRepo repo.OrderRepo, orderTempRepo repo.OrderTempRepo,
	beneficiaryRepo repo.BeneficiaryRepo) *OrderAggregate {
	return &OrderAggregate{
		orderAggregateRepo: orderAggregateRepo{
			orderRepo:       orderRepo,
			orderTempRepo:   orderTempRepo,
			beneficiaryRepo: beneficiaryRepo,
		},
	}
}

func (o *OrderAggregate) Instance(orderId ...string) (*OrderAggregate, error) {
	if len(orderId) < 1 {
		return o, nil
	}
	o.OrderId = orderId[0]

	if ca, err := o.orderRepo.GetOrderById(o.OrderId); err != nil {
		return o, err
	} else {
		util.StructCopy(&o.OrderEn, ca)
	}
	return o, nil
}

// Cancel 取消订单
func (o *OrderAggregate) Cancel() error {
	if temOrderDo, err := o.orderTempRepo.GetOrderById(o.OrderId); err != nil {
		return err
	} else {
		if temOrderDo.OrderStatus == types.OrderPayPending {
			// 更新订单状态为取消
			filter := types.B{"_id": o.OrderId}
			update := types.B{"orderStatus": types.OrderCancel}
			if err := o.orderTempRepo.UpdateOrder(filter, update); err != nil {
				return err
			}
		} else if temOrderDo.OrderStatus == types.OrderFinish {
			return errors.New("订单已完成,无法取消")
		}
	}
	return nil
}
