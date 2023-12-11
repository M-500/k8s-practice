//@Author: wulinlin
//@Description:
//@File:  pod
//@Version: 1.0.0
//@Date: 2023/12/10 15:54

package request

import corev1 "k8s.io/api/core/v1"

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Base struct {
	Name          string        `json:"name"`          // 名字
	Labels        []ListMapItem `json:"labels"`        // 标签
	Namespace     string        `json:"namespace"`     // 命名空间
	RestartPolicy string        `json:"restartPolicy"` // 重启策略 Always | Never | On-Failure
}

type Volume struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//	type Pod struct {
//		// base 基础定义信息
//		Base Base `json:"base"`
//		// 卷
//		Volumes []Volume `json:"volumes"`
//		// 网络相关
//		NetWorking string `json:"netWorking"`
//		// init containers
//		InitContainers []string `json:"initContainers"`
//		//
//		Containers []string `json:"containers"`
//	}
type DnsConfig struct {
	Nameservers []string `json:"nameservers"`
}

type NetWorking struct {
	HostNetwork bool          `json:"hostNetwork"`
	HostName    string        `json:"hostName"`
	DnsPolicy   string        `json:"dnsPolicy"`
	DnsConfig   DnsConfig     `json:"dnsConfig"`
	HostAliases []ListMapItem `json:"hostAliases"`
}
type Resources struct {
	Enable     bool  `json:"enable"`     //是否配置容器的配额
	MemRequest int32 `json:"memRequest"` //内存 Mi
	MemLimit   int32 `json:"memLimit"`
	CpuRequest int32 `json:"cpuRequest"` //cpu m
	CpuLimit   int32 `json:"cpuLimit"`
}

type VolumeMount struct {
	MountName string `json:"mountName"` //挂载卷名称
	MountPath string `json:"mountPath"` //挂载卷->对应的容器内的路径
	ReadOnly  bool   `json:"readOnly"`  //是否只读
}

type ProbeHttpGet struct {
	Scheme      string        `json:"scheme"`      //请求协议http / https
	Host        string        `json:"host"`        //请求host 如果为空 那么就是Pod内请求
	Path        string        `json:"path"`        //请求路径
	Port        int32         `json:"port"`        //请求端口
	HttpHeaders []ListMapItem `json:"httpHeaders"` //请求的header
}
type ProbeCommand struct {
	Command []string `json:"command"` // cat /test/test.txt
}

type ProbeTcpSocket struct {
	Host string `json:"host"` //请求host 如果为空 那么就是Pod内请求
	Port int32  `json:"port"` //探测端口
}

type ProbeTime struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds"` //初始化时间 初始化若干秒之后才开始探针
	PeriodSeconds       int32 `json:"periodSeconds"`       //每隔若干秒之后 去探针
	TimeoutSeconds      int32 `json:"timeoutSeconds"`      //探针等待时间 等待若干秒之后还没有返回 那么就是探测失败
	SuccessThreshold    int32 `json:"successThreshold"`    //探针若干次成功了 才认为这次探针成功
	FailureThreshold    int32 `json:"failureThreshold"`    //探测若干次 失败了 才认为这次探针失败
}

type ContainerProbe struct {
	//是否打开探针
	Enable bool `json:"enable"`
	//探针类型 tcp / http / exec
	Type      string         `json:"type"`
	HttpGet   ProbeHttpGet   `json:"httpGet"`
	Exec      ProbeCommand   `json:"exec"`
	TcpSocket ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}

type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}

type EnvVar struct {
	Name    string `json:"name"`
	RefName string `json:"refName"`
	Value   string `json:"value"`
	Type    string `json:"type"` //configMap | secret | default(k/v形式)
}
type EnvVarFromResource struct {
	Name    string `json:"name"`    //资源名称
	RefType string `json:"refType"` //configMap | secret
	Prefix  string `json:"prefix"`  //用于表示环境变量前缀
}

type Container struct {
	Name            string               `json:"name"`            //容器名称
	Image           string               `json:"image"`           //容器镜像
	ImagePullPolicy string               `json:"imagePullPolicy"` //镜像拉取策略
	Tty             bool                 `json:"tty"`             //是否开启伪终端
	Ports           []ContainerPort      `json:"ports"`
	WorkingDir      string               `json:"workingDir"` //工作目录
	Command         []string             `json:"command"`    //执行命令
	Args            []string             `json:"args"`       //参数
	Envs            []EnvVar             `json:"envs"`       //环境变量 [{key:value}]
	EnvsFrom        []EnvVarFromResource `json:"envsFrom"`
	Privileged      bool                 `json:"privileged"`     //是否开启模式
	Resources       Resources            `json:"resources"`      //容器申请配额
	VolumeMounts    []VolumeMount        `json:"volumeMounts"`   //容器卷挂载
	StartupProbe    ContainerProbe       `json:"startupProbe"`   //启动探针
	LivenessProbe   ContainerProbe       `json:"livenessProbe"`  //存活探针
	ReadinessProbe  ContainerProbe       `json:"readinessProbe"` //就绪探针
}

type NodeSelectorTermExpressions struct {
	Key      string                      `json:"key"`
	Operator corev1.NodeSelectorOperator `json:"operator"`
	Value    string                      `json:"value"`
}

type NodeScheduling struct {
	//nodeName nodeSelector nodeAffinity
	Type         string                        `json:"type"`
	NodeName     string                        `json:"nodeName"`
	NodeSelector []ListMapItem                 `json:"nodeSelector"`
	NodeAffinity []NodeSelectorTermExpressions `json:"nodeAffinity"`
}

type Pod struct {
	Base           Base                `json:"base"`        //基础定义信息
	Tolerations    []corev1.Toleration `json:"tolerations"` //pod 容忍
	NodeScheduling NodeScheduling      `json:"nodeScheduling"`
	Volumes        []Volume            `json:"volumes"`        // 卷
	NetWorking     NetWorking          `json:"netWorking"`     //网络相关
	InitContainers []Container         `json:"initContainers"` ///init containers
	Containers     []Container         `json:"containers"`     //containers
}
