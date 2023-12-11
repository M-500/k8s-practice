package response

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/11 16:21
//

type PodListItem struct {
	Name     string `json:"name"`     // 名称
	Ready    string `json:"ready"`    // 状态
	Status   string `json:"status"`   // Running/Error
	Restarts int32  `json:"restarts"` // N次
	Age      int64  `json:"age"`      // 运行时间
	IP       string `json:"IP"`       // Pod id
	Node     string `json:"node"`     //  pod被调度到哪个node
}
