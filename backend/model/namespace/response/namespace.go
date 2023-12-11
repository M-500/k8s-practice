//@Author: wulinlin
//@Description:
//@File:  namespace
//@Version: 1.0.0
//@Date: 2023/12/10 15:46

package response

type Namespace struct {
	Name              string `json:"name"`
	CreationTimestamp int64  `json:"creationTimestamp"`
	Status            string `json:"status"`
}
