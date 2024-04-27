package i18n_interfaces

type TranslatorEn struct{}

func (t *TranslatorEn) Galidator() TranslatorGalidatorI {
	return &TranslatorEnGalidator{}
}

func (t *TranslatorEn) HelloWorld() string {
	return "Hello World"
}

func (t *TranslatorEn) StatusCodes() TranslatorStatusCodesI {
	return &TranslatorEnStatusCodes{}
}

func (t *TranslatorEn) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorEnGalidator struct{}

func (t *TranslatorEnGalidator) Example() string {
	return "example"
}

func (t *TranslatorEnGalidator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorEnStatusCodes struct{}

func (t *TranslatorEnStatusCodes) InternalServerError() string {
	return "Internal server error"
}

func (t *TranslatorEnStatusCodes) PageNotFound() string {
	return "Page not found"
}

func (t *TranslatorEnStatusCodes) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}
