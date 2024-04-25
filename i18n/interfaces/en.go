package i18nInterfaces

import "fmt"

type TranslatorEn struct{}

func (t *TranslatorEn) Key() string {
	return "title"
}

func (t *TranslatorEn) Key2(message int, new string) string {
	return fmt.Sprintf("something %v", message)
}

func (t *TranslatorEn) New() TranslatorNewI {
	return &TranslatorEnNew{}
}

func (t *TranslatorEn) Other() string {
	return "Hi"
}

func (t *TranslatorEn) Voice() string {
	return "English"
}

type TranslatorEnNew struct{}

func (t *TranslatorEnNew) K() string {
	return ""
}

func (t *TranslatorEnNew) S(parameter int) string {
	return fmt.Sprintf("%v", parameter)
}
