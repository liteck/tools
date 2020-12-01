package tools

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
	"time"
)

func Valid(ptrInput interface{}) error {
	//if _, err := json.Marshal(ptrInput); err != nil {
	//	return err
	//}
	zh_ch := zh.New()
	validate := validator.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = validate.RegisterValidation("format_time_ymdhms", validateFormatTime)

	if err := validate.Struct(ptrInput); err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}

func validateFormatTime(fl validator.FieldLevel) bool {
	if _, err := time.Parse("2006-01-02 15:04:05", fl.Field().String()); err != nil {
		return false
	} else {
		return true
	}
}
