package common

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type LocaleFactory struct {
	new      func() locales.Translator
	register func(*validator.Validate, ut.Translator) error
}

var localeFactoryMap = map[Locale]LocaleFactory{
	LocaleEN: {
		new:      en.New,
		register: enTranslations.RegisterDefaultTranslations,
	},
	LocaleZH: {
		new:      zh.New,
		register: zhTranslations.RegisterDefaultTranslations,
	},
}

var (
	Validate              *validator.Validate
	Trans                 ut.Translator
	defaultFallbackLocale = LocaleZH
)

func InitValidator(l Locale) {
	factory, ok := localeFactoryMap[l]
	if !ok {
		factory, _ = localeFactoryMap[defaultFallbackLocale]
	}
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label != "" {
			return label
		}
		return Printer.Sprint(label)
	})
	locale := factory.new()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator(locale.Locale())
	_ = factory.register(Validate, Trans)
}

func ValidateStruct(v any) error {
	if err := Validate.Struct(v); err != nil {
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			return err
		}

		translations := errs.Translate(Trans)

		messages := make([]string, 0, len(translations))
		for _, msg := range translations {
			messages = append(messages, msg)
		}

		return errors.New(strings.Join(messages, "\n"))
	}

	return nil
}
