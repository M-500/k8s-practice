package pod

// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 21:11
type PodConvertGroup struct {
	K8s2ReqConvert
	Req2K8sConvert
}

const (
	probe_http = "http"
	probe_tcp  = "tcp"
	probe_exec = "exec"
)
const volume_type_emptydir = "emptyDir"

const (
	volume_emptyDir = "emptyDir"
)

const (
	scheduling_nodename     = "nodeName"
	scheduling_nodeselector = "nodeSelector"
	scheduling_nodeaffinity = "nodeAffinity"
	scheduling_nodeany      = "nodeAny"
)

const (
	ref_type_configMap = "configMap"
	ref_type_secret    = "secret"
)
