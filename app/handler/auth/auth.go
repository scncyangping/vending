package auth

import "github.com/gin-gonic/gin"

type BaseAuthService interface {
	Login(ctx *gin.Context)
	Registry(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

//
//type loginRequestBody struct {
//	Name string `json:name`
//	Pwd  string `json:pwd`
//}
//
//type AuthController struct {
//}
//
//func (a AuthController) Login(ctx *gin.Context) {
//	var (
//		requestBody loginRequestBody
//	)
//	err := ctx.ShouldBind(&requestBody)
//	if err != nil {
//		transport.SendFailure(ctx, err.Error())
//		return
//	}
//
//	if requestBody.Name == "aw" && requestBody.Pwd == "123" {
//		// 创建token
//		if token, err := jwt.GenerateToken(requestBody.Name); err != nil {
//			transport.SendFailure(ctx, "jwt generate error")
//			return
//		} else {
//			transport.SendSuccess(ctx, token)
//		}
//	} else {
//		transport.SendFailure(ctx, "password or name error")
//	}
//}
