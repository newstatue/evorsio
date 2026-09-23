package common

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

type Locale string

const (
	LocaleEN = "en"
	LocaleZH = "zh"
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
	defaultFallbackLocale Locale = LocaleZH
)

func InitValidator(l Locale) {
	factory, ok := localeFactoryMap[l]
	if !ok {
		factory, _ = localeFactoryMap[defaultFallbackLocale]
	}
	Validate = validator.New()
	locale := factory.new()
	uni := ut.New(locale, locale)
	Trans, _ = uni.GetTranslator(locale.Locale())
	_ = factory.register(Validate, Trans)
}
