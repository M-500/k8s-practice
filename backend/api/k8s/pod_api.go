package k8s

import (
	"backend/global"
	rsp "backend/pkg/comm/response"
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

// GetPodListOrDetail 获取pod列表或者详情
func (*PodApi) GetPodListOrDetail(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		detail, err := podService.GetPodDetail(namespace, name)
		if err != nil {
			rsp.JsonFailMsg(c, err.Error())
			return
		}
		rsp.JsonSuccessData(c, "获取Pod详情成功", detail)
	} else {
		err, items := podService.GetPodList(namespace, keyword, c.Query("nodeName"))
		if err != nil {
			rsp.JsonFailMsg(c, err.Error())
			return
		}
		rsp.JsonSuccessData(c, "获取Pod列表成功", items)
	}
}
