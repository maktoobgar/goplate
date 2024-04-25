package i18nInterfaces

import "fmt"

type Translator struct{}

func (t *Translator) Key() string {
	return "title"
}

func (t *Translator) Key2(message int, new string) string {
	return fmt.Sprintf("something %v %v", message, new)
}

func (t *Translator) New() TranslatorNewI {
	return &TranslatorNew{}
}

func (t *Translator) Other() string {
	return "ی سری چیزا به فارسی"
}

func (t *Translator) Voice() string {
	return "h"
}

type TranslatorNew struct{}

func (t *TranslatorNew) S(parameter int) string {
	return fmt.Sprintf("empty %v", parameter)
}
