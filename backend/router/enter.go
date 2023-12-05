package router

import (
	"backend/router/example"
	"backend/router/k8s"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:09
//

type routerGroup struct {
	ExampleRouterGroups example.ExampleRouter
	K8SRouterGroups     k8s.K8SRouter
}

var RouterGroupApp = new(routerGroup)
