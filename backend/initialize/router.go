package initialize

import (
	"backend/middlewares"
	"backend/router"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:13
//

func SetUpRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleWare())
	exampleRouterG := router.RouterGroupApp.ExampleRouterGroups
	exampleRouterG.InitExample(r)
	k8sRouterG := router.RouterGroupApp.K8SRouterGroups
	k8sRouterG.InitK8S(r)
	return r
}
