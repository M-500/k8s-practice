package example

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:06
//

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	group := r.Group("/example")
	apiGroup := api.NewApiGroups()
	group.GET("/test", apiGroup.ExampleApiGroup.TestHandler)
}
