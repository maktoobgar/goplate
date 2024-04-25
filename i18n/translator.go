package i18n

import "service/i18n/generated"

// Attribute 'lang' can be en, fa
func NewTranslator(lang string) generated.TranslatorI {
	if len(lang) >= 2 {
		lang = lang[:2]
	} else {
		lang = "fa"
	}

	if lang == "fa" {
		return &generated.Translator{}
	} else if lang == "en" {
		return &generated.TranslatorEn{}
	}

	return nil
}
