package initialize

import (
	"backend/global"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:52
//

func SetUpK8s() {
	kubeConfig := ".kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	// create the client set
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	global.KubeConfigSet = clientSet
}
