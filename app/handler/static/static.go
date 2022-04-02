/*
@date : 2019/11/18
@author : YaPi
@desc :
*/
package static

import "github.com/gin-gonic/gin"

func TestStatic(router *gin.RouterGroup) {

	router.GET("/test", TestDoSomth)
}
