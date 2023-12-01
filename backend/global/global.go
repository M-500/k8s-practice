package global

import (
	"k8s-imooc/config"
	"k8s.io/client-go/kubernetes"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:40
var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
