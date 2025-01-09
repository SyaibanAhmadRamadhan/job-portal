package presentations

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidate() *Validator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		if name == "" {
			name = strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
			if name == "-" {
				return ""
			}
		}

		return name
	})
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return &Validator{
		Validator:  v,
		Translator: trans,
	}
}
