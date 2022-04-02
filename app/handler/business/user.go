package business

import (
	"github.com/gin-gonic/gin"
	"vending/app/transport"
)

type UserService interface {
	WhoIam(ctx *gin.Context)
}

type UserController struct {
}

type whoIamRequestBody struct {
	Name string `json:name`
	Age  uint8  `json:age`
}

func (u UserController) WhoIam(ctx *gin.Context) {
	var (
		requestBody whoIamRequestBody
	)
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		transport.SendFailure(ctx, err.Error())
		return
	}

	transport.SendSuccess(ctx, map[string]interface{}{
		"name": requestBody.Name + " hao shen shou",
		"age":  requestBody.Age,
	})
}
