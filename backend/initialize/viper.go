package initialize

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/1 19:40
//

import (
	"backend/global"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("dev.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		panic(err.Error())
	}
}
