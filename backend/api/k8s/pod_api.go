package k8s

import (
	"backend/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:57
//

type PodApi struct {
}

func (*PodApi) GetPodList(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metaV1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "操了个DJ",
	})
}
