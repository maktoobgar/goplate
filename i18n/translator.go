package i18n

import "service/i18n/i18n_interfaces"

// Attribute 'lang' can be en, fa
func NewTranslator(lang string) i18n_interfaces.TranslatorI {
	if len(lang) >= 2 {
		lang = lang[:2]
	} else {
		lang = "fa"
	}

	if lang == "fa" {
		return &i18n_interfaces.Translator{}
	} else if lang == "en" {
		return &i18n_interfaces.TranslatorEn{}
	}

	return nil
}
