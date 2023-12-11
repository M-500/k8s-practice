package k8s

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 20:00
type K8SRouter struct {
}

func (*K8SRouter) InitK8S(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.NewApiGroups()
	group.GET("/get-pods", apiGroup.K8SApiGroup.GetPodList)
	group.GET("/pod/:namespace", apiGroup.K8SApiGroup.GetPodListOrDetail)
	group.GET("/namespace", apiGroup.K8SApiGroup.GetNamespaceList)
}
