package i18nInterfaces

type Translator struct{}

func (t *Translator) Galidator() TranslatorGalidatorI {
	return &TranslatorGalidator{}
}

func (t *Translator) HelloWorld() string {
	return "درود"
}

func (t *Translator) StatusCodes() TranslatorStatusCodesI {
	return &TranslatorStatusCodes{}
}

func (t *Translator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorGalidator struct{}

func (t *TranslatorGalidator) Example() string {
	return "مثال"
}

func (t *TranslatorGalidator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorStatusCodes struct{}

func (t *TranslatorStatusCodes) InternalServerError() string {
	return "خطایی در سرور رخ داده است"
}

func (t *TranslatorStatusCodes) PageNotFound() string {
	return "صفحه مورد نظر یافت نشد"
}

func (t *TranslatorStatusCodes) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}
