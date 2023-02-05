package gosdk

import (
	"github.com/go-playground/locales/zh_Hans_SG"
	ut "github.com/go-playground/universal-translator"
	zhs "github.com/go-playground/validator/v10/translations/zh"

	"github.com/go-playground/validator/v10"
	gbox "github.com/mvity/go-box"
	"log"
	"reflect"
)

// 是否打开调试模式
var debug = true

var Validate *validator.Validate
var Translator ut.Translator

func init() {
	gbox.SetDebug(true)

	Validate = validator.New()
	chinese := zh_Hans_SG.New()
	uti := ut.New(chinese, chinese)
	Translator, _ = uti.GetTranslator("zh")
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tag := field.Tag.Get("label")
		if tag == "" {
			tag = "[" + field.Name + "]"
		}
		return tag
	})
	_ = zhs.RegisterDefaultTranslations(Validate, Translator)

}

// SetDebug 设置GoBox是否为调试模式
func SetDebug(_debug bool) {
	debug = _debug
	gbox.SetDebug(debug)
}

// IsDebug 是否调试模式
func IsDebug() bool {
	return debug
}

// WARN 输出警告信息，仅在调试模式下有效
func WARN(format string, v ...any) {
	if debug {
		log.Printf("WARN: "+format, v...)
	}
}
