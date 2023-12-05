package main

import (
	"backend/global"
	"backend/initialize"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 18:50
//

func main() {
	r := initialize.SetUpRouters()
	initialize.Viper()
	panic(r.Run(global.CONF.System.Addr))
}
