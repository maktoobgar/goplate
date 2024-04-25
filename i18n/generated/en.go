package generated

type TranslatorEn struct{}

func (t *TranslatorEn) Key() string {
	return "title"
}

func (t *TranslatorEn) Key2() string {
	return "something {message:number}"
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

func (t *TranslatorEnNew) S() string {
	return ""
}
