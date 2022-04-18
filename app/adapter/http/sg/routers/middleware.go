package routers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"io/ioutil"
	"time"
	"vending/app/infrastructure/config"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types/constants"
)

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			actionId string
		)
		actionId = snowflake.NextId()
		data, err := ctx.GetRawData()
		if err != nil {
			log.Logger().Errorf("Visit Param Init Error %v", err)
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		log.Logger().Infof(
			"%s request %s [%s] from [%s]: %v",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			actionId, ctx.ClientIP(),
			data)

		ctx.Set("ActionId", actionId)
		ctx.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			errFlag int
			session = &jwt.Session{}
		)
		token = c.Request.FormValue(config.Base.Jwt.JwtAuthKey)
		if token == constants.EmptyStr {
			token = c.GetHeader(config.Base.Jwt.JwtAuthKey)
		}
		if token == constants.EmptyStr {
			c.Abort()
			transport.SendFailure(c, transport.RequestTokenNotFound, transport.StatusText(transport.RequestTokenNotFound))
			return

		}
		if claims, err := jwt.ParseToken(token); err != nil {
			errFlag = transport.RequestCheckTokenError
		} else if time.Now().Unix() > claims.ExpiresAt {
			errFlag = transport.RequestCheckTokenTimeOut
		} else {
			// 设置登录信息到token里面
			session.UserName = claims.Username
			c.Set("Session", session)
		}
		if errFlag > constants.ZERO {
			c.Abort()
			transport.SendFailure(c, errFlag, transport.StatusText(errFlag))
			return
		}
		c.Next()
	}
}

func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.Abort()
			transport.SendFailure(c, transport.RateLimit, transport.StatusText(transport.RateLimit))
			return
		}
		c.Next()
	}
}
