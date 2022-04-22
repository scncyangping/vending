package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
	"vending/app/infrastructure/config"
	"vending/app/infrastructure/pkg/log"
	"vending/app/infrastructure/pkg/util/jwt"
	"vending/app/infrastructure/pkg/util/snowflake"
	"vending/app/types/constants"
)

// Handler 具体业务服务聚合
type Handler struct {
	Logger *zap.SugaredLogger
}

func NewHandler() *Handler {
	return &Handler{
		Logger: log.Logger(),
	}
}

func (h *Handler) LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			actionId string
		)
		actionId = snowflake.NextId()
		data, err := ctx.GetRawData()
		if err != nil {
			h.Logger.Errorf("Visit Param Init Error %v", err)
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		h.Logger.Infof(
			"%s request %s [%s] from [%s]: %v",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			actionId, ctx.ClientIP(),
			string(data))

		ctx.Set("ActionId", actionId)
		ctx.Next()
	}
}

func (h *Handler) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			errFlag ResultCode
		)
		token = c.Request.FormValue(config.Base.Jwt.JwtAuthKey)
		if token == constants.EmptyStr {
			token = c.GetHeader(config.Base.Jwt.JwtAuthKey)
		}
		if token == constants.EmptyStr {
			c.Abort()
			h.SendFailure(c, RequestTokenNotFound, StatusText(RequestTokenNotFound))
			return

		}
		if claims, err := jwt.ParseToken(token); err != nil {
			errFlag = RequestCheckTokenError
		} else if time.Now().Unix() > claims.ExpiresAt {
			errFlag = RequestCheckTokenTimeOut
		} else {
			// 设置登录信息到token里面
			c.Set("username", claims.Username)
		}
		if errFlag > constants.ZERO {
			c.Abort()
			h.SendFailure(c, errFlag, StatusText(errFlag))
			return
		}
		c.Next()
	}
}

func (h *Handler) RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.Abort()
			h.SendFailure(c, RateLimit, StatusText(RateLimit))
			return
		}
		c.Next()
	}
}
