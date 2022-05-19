package types

type AuthenticationType string

type B map[string]any

const (
	JWT AuthenticationType = "JWT"
)

type ResultCode string // 返回代码
type ResultMsg string  // 返回信息
type UserStatus uint8  // 用户状态
const (
	UserNormal UserStatus = 1 << iota
	UserFrozen

	PayWaitBuyerPay PayStatus = 1 << iota
	PayClosed
	PaySuccess
	PayFinished
)

type PayStatus uint8 // 支付状态

var PayStatusM = map[PayStatus]string{
	PayWaitBuyerPay: "待支付",
	PaySuccess:      "已支付",
	PayClosed:       "已取消",
	PayFinished:     "已完成",
}

type BeneficiaryType uint8 // 收款方式

const (
	BfAlipayFace BeneficiaryType = 1 << iota // 支付宝当面付
	BfWechat
)

type BeneficiaryStatus uint8 // 收款方式状态

const (
	BfUse BeneficiaryStatus = 1 << iota
	BfStop
	BfFrozen
)

var BeneficiaryStatusM = map[BeneficiaryStatus]string{
	BfUse:    "使用中",
	BfStop:   "停止",
	BfFrozen: "冻结",
}

type CommodityStatus uint8 // 商品状态

const (
	CommodityPending CommodityStatus = 1 << iota
	CommodityUp
	CommodityDown
)

var CommodityStatusM = map[CommodityStatus]string{
	CommodityPending: "审核中",
	CommodityUp:      "已上架",
	CommodityDown:    "已下架",
}

type OrderStatus uint8 // 订单状态

const (
	OrderPayPending OrderStatus = 1 << iota
	OrderFinish
	OrderCancel
)

var OrderStatusM = map[OrderStatus]string{
	OrderPayPending: "订单待支付",
	OrderFinish:     "订单已完成",
	OrderCancel:     "订单已取消",
}

type StockStatus uint8 // 库存状态

const (
	StockNormal StockStatus = 1 << iota
	StockUsed
)

var StockStatusM = map[StockStatus]string{
	StockNormal: "待使用",
	StockUsed:   "已使用",
}

type SellType uint8

const (
	Once SellType = 1 << iota
	Repeat
)

var SellTypeM = map[SellType]string{
	Once:   "一次性",
	Repeat: "可重复使用",
}
