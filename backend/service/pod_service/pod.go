package pod_service

import (
	"backend/global"
	pod_req "backend/model/pod/request"
	pod_res "backend/model/pod/response"
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 21:09
//

type PodService struct {
}

func (*PodService) GetPodDetail(namespace string, name string) (podReq pod_req.Pod, err error) {
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(namespace)
	k8sGetPod, err := podApi.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("Pod[namespace=%s,name=%s]查询失败，detail：%s", namespace, name, err.Error())
		err = errors.New(errMsg)
		return
	}
	//将k8s pod 转为 pod request
	podReq = podConvert.K8s2ReqConvert.PodK8s2Req(*k8sGetPod)
	return
}
func (*PodService) GetPodList(namespace string, keyword string, nodeName string) (err error, _ []pod_res.PodListItem) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	podList := make([]pod_res.PodListItem, 0)
	for _, item := range list.Items {
		if nodeName != "" && item.Spec.NodeName != nodeName {
			continue
		}
		if strings.Contains(item.Name, keyword) {
			podItem := podConvert.PodK8s2ItemRes(item)
			podList = append(podList, podItem)
		}
	}
	return err, podList
}
