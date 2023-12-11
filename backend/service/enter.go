package service

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 21:25
//
import "backend/service/pod_service"

// @Author: morris
type ServiceGroup struct {
	PodServiceGroup pod_service.PodServiceGroup
	//NodeServiceGroup      node.Group
	//ConfigMapServiceGroup configmap.ServiceGroup
	//SecretServiceGroup    secret.SeviceGroup
}

var ServiceGroupApp = new(ServiceGroup)
