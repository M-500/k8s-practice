package validators

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:53
//

import (
	"backend/pkg/comm/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var InstanceTrans ut.Translator

func SetUpTrans(locale string) ut.Translator {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, zhT)
		InstanceTrans, ok = uni.GetTranslator(locale)
		if !ok {
			//return fmt.Errorf("uni.GetTranslator(%s)", locale)
			panic(any(fmt.Errorf("uni.GetTranslator(%s)", locale)))
		}

		switch locale {
		case "en":
			entranslations.RegisterDefaultTranslations(v, InstanceTrans)
		case "zh":
			zhtranslations.RegisterDefaultTranslations(v, InstanceTrans)
		default:
			entranslations.RegisterDefaultTranslations(v, InstanceTrans)
		}
		return InstanceTrans
	}
	return nil
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandlerValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		response.JsonFailMsg(c, err.Error())
		return
	}
	response.JsonFailData(c, "表单校验错误", removeTopStruct(errs.Translate(InstanceTrans)))
}
