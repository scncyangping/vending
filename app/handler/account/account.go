package account

import "github.com/gin-gonic/gin"

type BaseAccountService interface {
	CreateAccount(ctx *gin.Context)
	GetAccount(ctx *gin.Context)
	DeleteAccount(ctx *gin.Context)
}

type Acc struct {
	Name string
	bl   float64
}
