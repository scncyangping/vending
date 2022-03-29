package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"time"
	"vending/data/common/config/log"
)

func ZapLoggingMiddleWare() endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.ZapLogger.Debug("ZapLoggingMiddleWare")
			}(time.Now())
			return e(ctx, request)
		}
	}
}

func JsonWebTokenMiddleWare() endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.ZapLogger.Debugf("JsonWebTokenMiddleWare")
			}(time.Now())
			return e(ctx, request)
		}
	}
}
