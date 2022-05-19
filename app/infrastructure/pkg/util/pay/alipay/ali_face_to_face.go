package alipay

import (
	"errors"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"net/url"
	"strconv"
	"vending/app/domain/obj"
	"vending/app/infrastructure/pkg/log"
	"vending/app/types"
	"vending/app/types/constants"
)

type AliFaceToFace struct {
	client          *alipay.Client
	notifyUrl       string // 支付回调通知
	AppId           string // 商户应用id
	AppPrivateKey   string // 应用私钥
	AliPayPublicKey string // 支付公钥
}

func New(appId, appPrivateKey, aliPayPublicKey, notifyUrl string) (*AliFaceToFace, error) {
	var (
		err    error
		client *alipay.Client
	)
	// 普通公钥方式
	// appId 和 应用私钥
	if client, err = alipay.New(appId, appPrivateKey, false); err != nil {
		return nil, err
	} else {
		// 加载支付宝公钥
		if err = client.LoadAliPayPublicKey(aliPayPublicKey); err != nil {
			return nil, err
		}
	}
	return &AliFaceToFace{
		client:          client,
		notifyUrl:       notifyUrl,
		AppId:           appId,
		AppPrivateKey:   appPrivateKey,
		AliPayPublicKey: aliPayPublicKey,
	}, nil
}

func (a *AliFaceToFace) Start(faceBody any) (any, error) {
	var (
		ok    bool
		param alipay.TradePreCreate
		body  obj.OrderAlipayFaceBody
	)
	// 参数类型校验
	if body, ok = faceBody.(obj.OrderAlipayFaceBody); !ok {
		return nil, errors.New("面对面支付参数错误")
	} else {
		// 参数合法性校验
		if body.OutTradeNo == constants.EmptyStr || len(body.OutTradeNo) != 19 ||
			body.Subject == constants.EmptyStr || body.TotalAmount >= 1000 || body.TotalAmount < 0 {
			log.Logger().Errorf("【Alipay 面对面支付】面对面支付参数错误, body: [ %v ]", body)
			return nil, errors.New("面对面支付参数错误")
		}
	}

	// 面对面支付默认
	param.ProductCode = "FACE_TO_FACE_PAYMENT"

	param.OutTradeNo = body.OutTradeNo
	param.Subject = body.Subject
	param.NotifyURL = a.notifyUrl
	param.TotalAmount = strconv.FormatFloat(body.TotalAmount, 'f', 2, 64)

	if rsp, err := a.client.TradePreCreate(param); err != nil {
		return nil, errors.New("面对面支付错误")
	} else {
		if !rsp.IsSuccess() {
			log.Logger().Errorf("【Alipay 面对面支付】创建订单 %v 信息发生错误: %s-%s", body, rsp.Content.Msg, rsp.Content.SubMsg)
			return nil, errors.New("创建面对面支付错误")
		}
		return rsp.Content.QRCode, nil
	}
}

func (a *AliFaceToFace) Status(outTradeNo string) (types.PayStatus, error) {
	var (
		q       alipay.TradeQuery
		rStatus types.PayStatus
	)
	if outTradeNo == constants.EmptyStr || len(outTradeNo) != 19 {
		return types.PayWaitBuyerPay, errors.New("面对面支付查询支付状态参数错误")
	}

	q.OutTradeNo = outTradeNo

	if rsp, err := a.client.TradeQuery(q); err != nil {
		return types.PayWaitBuyerPay, err
	} else {
		if !rsp.IsSuccess() {
			errMsg := fmt.Sprintf("【Alipay 面对面支付】验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
			log.Logger().Error(errMsg)
			return types.PayWaitBuyerPay, errors.New(errMsg)
		} else {
			switch rsp.Content.TradeStatus {
			case alipay.TradeStatusWaitBuyerPay: // "WAIT_BUYER_PAY" 交易创建，等待买家付款
				rStatus = types.PayWaitBuyerPay
			case alipay.TradeStatusClosed: // "TRADE_CLOSED" 未付款交易超时关闭，或支付完成后全额退款
				rStatus = types.PayClosed
			case alipay.TradeStatusSuccess: // "TRADE_SUCCESS" 交易支付成功
				rStatus = types.PaySuccess
			case alipay.TradeStatusFinished: // "TRADE_FINISHED" 交易结束，不可退款
				rStatus = types.PayFinished
			default:
				rStatus = types.PayWaitBuyerPay
				err = errors.New("面对面支付查询支付状态返回错误")
			}
			return rStatus, err
		}
	}
}

func (a *AliFaceToFace) Notify(qDta any) (bool, error) {
	var (
		ok         bool
		errs       error
		outTradeNo string
	)
	if body, ok := qDta.(url.Values); !ok {
		errs = errors.New("面对面支付参数错误")
	} else {
		if ok, err := a.client.VerifySign(body); err != nil || !ok {
			errs = errors.New("面对面支付异步通知验签错误")
		} else {
			outTradeNo = body.Get("out_trade_no")

			if status, err := a.Status(outTradeNo); err != nil {
				errs = err
			} else {
				if status == types.PaySuccess || status == types.PayFinished {
					ok = true
				}
			}
		}
	}
	log.Logger().Infof("【Alipay 面对面支付】回调信息 %v 是否成功: %t  失败参数: %v", qDta, ok, errs)
	return ok, errs
}
