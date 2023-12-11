//@Author: wulinlin
//@Description:
//@File:  namespace_api
//@Version: 1.0.0
//@Date: 2023/12/10 15:33

package k8s

import (
	"backend/global"
	"backend/model/namespace/response"
	rsp "backend/pkg/comm/response"
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceAPI struct {
}

func (*NamespaceAPI) GetNamespaceList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	namespaceList := make([]response.Namespace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, response.Namespace{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.Phase),
		})
	}
	rsp.JsonSuccessData(c, "获取命名空间列表成功", namespaceList)
	return
}
