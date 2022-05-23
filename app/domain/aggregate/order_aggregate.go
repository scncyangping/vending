package aggregate

import (
	"errors"
	"vending/app/domain/aggregate/factory"
	"vending/app/domain/entity"
	"vending/app/domain/obj"
	"vending/app/domain/repo"
	"vending/app/infrastructure/do"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types"
	"vending/app/types/constants"
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
	OrderId       string
	OrderEn       entity.OrderEn       // 订单基础数据
	Items         []obj.OrderItemObj   // 订单明细
	PayDesObj     obj.PayDesObj        // 支付明细
	BeneficiaryEn entity.BeneficiaryEn // 收款信息镜像
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
func (o *OrderAggregate) CreateTempOrderOne(commodityId string, num int,
	payType types.BeneficiaryType, payDes obj.PayDesObj) (string, error) {
	var (
		orderEn entity.OrderEn
		orderId = snowflake.NextId()
		items   = make([]obj.OrderItemObj, 0)
	)
	// 组装商品明细
	if i, err := o.buildOrderItem(payType, commodityId, num); err != nil {
		return constants.EmptyStr, err
	} else {
		items = append(items, i)

		originalAmount, amount := o.orderBuild(&items)
		orderEn.OriginalAmount = originalAmount
		orderEn.Amount = amount
		// TODO 默认取第一个支付方式
		orderEn.BfObj = items[0].Payment
	}
	orderEn.Id = orderId                        // 预定义订单id
	orderEn.Items = items                       // 商品即商品支付明细
	orderEn.PayDesObj = payDes                  // 订单描述
	orderEn.OrderStatus = types.OrderPayPending // 订单状态创建为待支付

	// 临时订单
	if _, err := o.orderTempRepo.SaveOrder(&orderEn); err != nil {
		return constants.EmptyStr, err
	}

	pg, _ := factory.Instance.PayAggregateInstance()
	if payUrl, err := pg.Pay(orderId, orderEn.Amount); err != nil {
		return constants.EmptyStr, err
	} else {
		return payUrl, nil
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

func (o *OrderAggregate) buildOrderItem(payType types.BeneficiaryType, commodityId string, num int) (obj.OrderItemObj, error) {
	var (
		item obj.OrderItemObj
	)
	// 商品聚合
	// 确认商品是否存在,取其金额
	if cg, err := factory.Instance.CommodityAggregateInstance(commodityId); err != nil {
		return item, err
	} else {
		item.CommodityId = cg.commodityEn.Id
		item.Amount = cg.commodityEn.Amount * float64(num)
		item.CommodityName = cg.commodityEn.Name
		item.OriginalAmount = cg.commodityEn.Amount * float64(num) // 原金额 TOTO 可加优惠策略
		item.OwnerId = cg.commodityEn.OwnerId
		// 配置收款信息
		if bf, err := o.beneficiaryRepo.GetBeneficiaryByOwnerIdAndType(cg.commodityEn.OwnerId, payType); err != nil {
			return item, err
		} else {
			pObj := obj.BeneficiaryObj{}
			util.StructCopy(pObj, bf)
			item.Payment = pObj
		}
		return item, nil
	}
}

func (o *OrderAggregate) preOutStack(orderId, categoryId string, num int) ([]*entity.StockEn, error) {
	// 校验库存并出货预锁定
	if ig, err := factory.Instance.InventoryAggregateInstance(categoryId); err != nil {
		return nil, err
	} else {
		if ig.OutOfStock() {
			return nil, errors.New("该商品缺货,请联系客服")
		} else {
			// 出货,预锁定
			if es, err := ig.OutStock(orderId, num); err != nil {
				return nil, err
			} else {
				return es, err
			}
		}
	}
}

func (o *OrderAggregate) orderBuild(item *[]obj.OrderItemObj) (float64, float64) {
	var (
		originalAmount float64
		amount         float64
	)
	if item == nil || len(*item) < 1 {
		log.Logger().Errorf("[orderBuild] 参数异常 %v", *item)
	} else {
		for _, v := range *item {
			originalAmount += v.OriginalAmount
			amount += v.Amount
		}
	}
	return originalAmount, amount
}
