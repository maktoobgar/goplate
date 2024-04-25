package i18n

import "service/i18n/interfaces"

// Attribute 'lang' can be en, fa
func NewTranslator(lang string) i18nInterfaces.TranslatorI {
	if len(lang) >= 2 {
		lang = lang[:2]
	} else {
		lang = "fa"
	}

	if lang == "fa" {
		return &i18nInterfaces.Translator{}
	} else if lang == "en" {
		return &i18nInterfaces.TranslatorEn{}
	}

	return nil
}
