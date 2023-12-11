//@Author: wulinlin
//@Description:
//@File:  pod
//@Version: 1.0.0
//@Date: 2023/12/10 15:54

package request

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

type Pod struct {
	// base 基础定义信息
	Base Base `json:"base"`
	// 卷
	Volumes []Volume `json:"volumes"`
	// 网络相关
	NetWorking string `json:"netWorking"`
	// init containers
	InitContainers []string `json:"initContainers"`
	//
	Containers []string `json:"containers"`
}
