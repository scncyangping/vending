package order

import (
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types"
)

type DoOrderSrvImpl struct {
	beneficiaryRepo repo.BeneficiaryRepo
	orderTempRepo   repo.OrderTempRepo
	orderRepo       repo.OrderRepo
}

func NewDoOrderSrvImpl(beneficiaryRepo repo.BeneficiaryRepo, orderTempRepo repo.OrderTempRepo, orderRepo repo.OrderRepo) *DoOrderSrvImpl {
	return &DoOrderSrvImpl{beneficiaryRepo: beneficiaryRepo, orderTempRepo: orderTempRepo, orderRepo: orderRepo}
}

func (o *DoOrderSrvImpl) QueryOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error) {
	return o.orderRepo.ListOrderPageBy(skip, limit, sort, filter)
}

func (o *DoOrderSrvImpl) QueryTempOrderPageBy(skip, limit int64, sort, filter any) ([]*do.OrderDo, error) {
	return o.orderTempRepo.ListOrderPageBy(skip, limit, sort, filter)
}
func (o *DoOrderSrvImpl) GetTempOrderById(id string) (*do.OrderDo, error) {
	return o.orderTempRepo.GetOrderById(id)
}

func (o *DoOrderSrvImpl) GetOrderById(id string) (*do.OrderDo, error) {
	return o.orderRepo.GetOrderById(id)
}

// CreateTempOrderOne 下单
// 创建订单仅需支付信息金额
func (o *DoOrderSrvImpl) CreateTempOrderOne(items []obj.OrderItemObj, desObj obj.PayDesObj) (string, float64, error) {
	var (
		orderEn entity.OrderEn
		orderId = snowflake.NextId()
	)
	orderEn.Id = orderId                        // 预定义订单id
	orderEn.Items = items                       // 商品即商品支付明细
	orderEn.PayDesObj = desObj                  // 订单描述
	orderEn.OrderStatus = types.OrderPayPending // 订单状态创建为待支付

	// TODO 平台收款方式 暂时指定Id为admin
	bf, _ := o.beneficiaryRepo.GetBeneficiaryByOwnerIdOrTypeDefault("admin", types.BfAlipayFace)
	pObj := obj.BeneficiaryObj{}
	util.StructCopy(pObj, bf)

	orderEn.BfObj = pObj // 支付数据

	for _, v := range items {
		orderEn.OriginalAmount += v.OriginalAmount
		orderEn.Amount += v.Amount
	}

	// 创建临时订单
	if _, err := o.orderTempRepo.SaveOrder(&orderEn); err != nil {
		return orderId, orderEn.Amount, err
	} else {
		return orderId, orderEn.Amount, nil
	}

}

// SaveOrder 创建订单
// 在支付完成后回调
func (o *DoOrderSrvImpl) SaveOrder(orderId string) error {
	var (
		err        error
		temOrderDo *do.OrderDo
		orderEn    entity.OrderEn
	)
	if temOrderDo, err = o.orderTempRepo.GetOrderById(orderId); err != nil {
		return err
	} else {
		// 将当前数据拷贝到订单标中
		util.StructCopy(&orderEn, temOrderDo)
		if _, err = o.orderRepo.SaveOrder(&orderEn); err != nil {
			return err
		}
	}
	return nil
}
