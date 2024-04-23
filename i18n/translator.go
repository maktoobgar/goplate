package i18n

import "service/i18n/interfaces"

func NewTranslator(lang string) interfaces.WordsI {
	if len(lang) >= 2 {
		lang = lang[:2]
	} else {
		lang = "en"
	}

	return nil
}
