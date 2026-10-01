package common

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleZH Locale = "zh"
)

type translation struct {
	lang  language.Tag
	key   string
	value string
}

var translations = []translation{
	{language.Chinese, "vault.title", "标题"},
	{language.English, "vault.title", "Title"},

	{language.Chinese, "vault.password", "密码"},
	{language.English, "vault.password", "Password"},

	{language.Chinese, "vault.username", "用户名"},
	{language.English, "vault.username", "Username"},
}

var Printer *message.Printer

func InitI18n(locale Locale) {
	switch locale {
	case LocaleEN:
		Printer = message.NewPrinter(language.English)
	default:
		Printer = message.NewPrinter(language.Chinese)
	}
	for _, t := range translations {
		if err := message.SetString(t.lang, t.key, t.value); err != nil {
			panic(fmt.Errorf(
				"register translation %q for %s: %w",
				t.key,
				t.lang,
				err,
			))
		}
	}
}
