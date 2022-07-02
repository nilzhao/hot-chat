package utils

import (
	"errors"
	"red-server/global"
	"strings"
	"sync"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	validate   *validator.Validate
	translator ut.Translator
	once       sync.Once
)

func init() {
	once.Do(func() {
		validate = validator.New()
		//创建消息国际化通用翻译器
		cn := zh.New()
		uni := ut.New(cn, cn)
		var found bool
		translator, found = uni.GetTranslator("zh")
		if found {
			err := zhTranslations.RegisterDefaultTranslations(validate, translator)
			if err != nil {
				global.Logger.Error(err)
			}
		} else {
			global.Logger.Error("Not found translator: zh")
		}
	})
}

func ValidateStruct(s interface{}) (err error) {
	errMsgs := []string{}
	err = validate.Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			global.Logger.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				errMsg := e.Translate(translator)
				errMsgs = append(errMsgs, errMsg)
			}
		}
		return errors.New(strings.Join(errMsgs, ";"))
	}
	return nil
}
