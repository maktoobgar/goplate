package generated

type Translator struct{}

func (t *Translator) Key() string {
	return "title"
}

func (t *Translator) Key2() string {
	return "something {message:number}"
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

func (t *TranslatorNew) S() string {
	return "empty"
}
