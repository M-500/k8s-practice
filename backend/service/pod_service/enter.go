package pod_service

import "backend/convert"

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 21:18
type PodServiceGroup struct {
	PodService
}

var podConvert = convert.ConvertGroupApp.PodConvert
