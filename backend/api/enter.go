package api

import (
	"backend/api/example"
	"backend/api/k8s"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 18:58
//

type ApiGroups struct {
	ExampleApiGroup example.Controller
	K8SApiGroup     k8s.Controller
}

func NewApiGroups() *ApiGroups {
	return &ApiGroups{}
}

var APIGroupApp = NewApiGroups()
