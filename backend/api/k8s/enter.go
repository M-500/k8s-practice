package k8s

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:58
//
import (
	"backend/service"
)

type Controller struct {
	PodApi
	NamespaceAPI
}

// var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService

//var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
//var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService
//var secretService = service.ServiceGroupApp.SecretServiceGroup.SecretService
