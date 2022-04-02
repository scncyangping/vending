/*
@date : 2019/11/18
@author : YaPi
@desc :
*/
package static

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestDoSomth(ctx *gin.Context) {
	fmt.Println(ctx.Request.URL)
	fmt.Println(ctx.Request.RequestURI)

	ctx.HTML(200, "403.html", gin.H{"hello": "world"})
}
