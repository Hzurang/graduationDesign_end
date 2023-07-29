package config

import (
	"fmt"
	"ginStudy/model"
	"ginStudy/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

// 定义一个全局翻译器T
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		//注册自定义校验方法
		v.RegisterStructValidation(utils.SignUpParamStructLevelValidation, model.ParamMobilePasswordSignUp{})
		if err = v.RegisterValidation("mobile", utils.VerifyMobileParam); err != nil {
			zap.L().Error(utils.GetCodeMsg(90013), zap.Error(err))
			return err
		}
		if err = v.RegisterValidation("id_card", utils.VerifyIDCardParam); err != nil {
			zap.L().Error(utils.GetCodeMsg(90013), zap.Error(err))
			return err
		}
		if err = v.RegisterValidation("checkDate", utils.VerifyDate); err != nil {
			zap.L().Error(utils.GetCodeMsg(90013), zap.Error(err))
			return err
		}
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找, 获取实例
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		if err = v.RegisterTranslation("mobile", Trans, registerTranslator("mobile", "{0}格式错误"), translate); err != nil {
			zap.L().Error(utils.GetCodeMsg(90014), zap.Error(err))
			return err
		}
		if err = v.RegisterTranslation("id_card", Trans, registerTranslator("id_card", "{0}格式错误"), translate); err != nil {
			zap.L().Error(utils.GetCodeMsg(90014), zap.Error(err))
			return err
		}
		if err = v.RegisterTranslation("checkDate", Trans, registerTranslator("checkDate", "{0}必须要晚于当前日期"), translate); err != nil {
			zap.L().Error(utils.GetCodeMsg(90014), zap.Error(err))
			return err
		}
		return
	}
	return
}

/*
registerTranslator
@author: LJR
@Description: 为自定义字段添加翻译功能
@param tag
@param msg
@return validator.RegisterTranslationsFunc
*/
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

/*
translate
@author: LJR
@Description: 自定义字段的翻译方法
@param trans
@param fe
@return string
*/
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
