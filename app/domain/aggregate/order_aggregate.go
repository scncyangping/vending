package aggregate

import (
	"errors"
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/snowflake"
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

// CreateTempOrderOne 下单
// 创建订单仅需支付信息金额
func (o *OrderAggregate) CreateTempOrderOne(items []obj.OrderItemObj, desObj obj.PayDesObj) (string, float64, error) {
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
func (o *OrderAggregate) SaveOrder(orderId string) error {
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

// Cancel 取消订单
func (o *OrderAggregate) Cancel(orderId string) error {
	if temOrderDo, err := o.orderTempRepo.GetOrderById(orderId); err != nil {
		return err
	} else {
		if temOrderDo.OrderStatus == types.OrderPayPending {
			// 更新订单状态为取消
			filter := types.B{"_id": orderId}
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

func (o *OrderAggregate) BuildOrderItem(m map[string]int, commodityEns []*entity.CommodityEn) ([]obj.OrderItemObj, error) {
	objs := make([]obj.OrderItemObj, 0)
	for _, v := range commodityEns {
		if obj, err := o.buildOrderItem(v, m[v.Id]); err != nil {
			return objs, err
		} else {
			objs = append(objs, obj)
		}
	}
	return objs, nil
}

func (o *OrderAggregate) buildOrderItem(commodityEn *entity.CommodityEn, num int) (obj.OrderItemObj, error) {
	var (
		item obj.OrderItemObj
	)
	// 商品聚合
	item.CommodityId = commodityEn.Id
	item.Amount = commodityEn.Amount * float64(num)
	item.CommodityName = commodityEn.Name
	item.OriginalAmount = commodityEn.Amount * float64(num) // 原金额 TOTO 可加优惠策略
	item.OwnerId = commodityEn.OwnerId
	return item, nil
}

func (o *OrderAggregate) preOutStack(orderId, categoryId string, num int) ([]*entity.StockEn, error) {
	// 校验库存并出货预锁定
	//if ig, err := factory.Instance.InventoryAggregateInstance(categoryId); err != nil {
	//	return nil, err
	//} else {
	//	if ig.OutOfStock() {
	//		return nil, errors.New("该商品缺货,请联系客服")
	//	} else {
	//		// 出货,预锁定
	//		if es, err := ig.OutStock(orderId, num); err != nil {
	//			return nil, err
	//		} else {
	//			return es, err
	//		}
	//	}
	//}
	return nil, nil
}
