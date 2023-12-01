package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 18:56
//

type exampleController struct {
}

func (*Controller) TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "哈哈哈",
	})
}
