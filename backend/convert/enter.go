package convert

import "backend/convert/pod"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 21:18
//

type ConvertGroup struct {
	PodConvert pod.PodConvertGroup
	//NodeConvert      node.Group
	//ConfigMapConvert configmap.ConvertGroup
	//SecretConvert    secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
