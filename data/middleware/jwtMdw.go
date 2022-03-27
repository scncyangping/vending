package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"time"
	"vending/data/common/config/log"
)

func JsonWebTokenMiddleWare() endpoint.Middleware {
	fmt.Println("JsonWebTokenMiddleWare")
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.ZapLogger.Debugf("JsonWebTokenMiddleWare")
			}(time.Now())
			return e(ctx, request)
		}
	}
}
