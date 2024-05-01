package i18n_interfaces

type TranslatorEn struct{}

func (t *TranslatorEn) Auth() TranslatorAuthI {
	return &TranslatorEnAuth{}
}

func (t *TranslatorEn) Galidator() TranslatorGalidatorI {
	return &TranslatorEnGalidator{}
}

func (t *TranslatorEn) HelloWorld() string {
	return "Hello World"
}

func (t *TranslatorEn) StatusCodes() TranslatorStatusCodesI {
	return &TranslatorEnStatusCodes{}
}

func (t *TranslatorEn) Users() TranslatorUsersI {
	return &TranslatorEnUsers{}
}

func (t *TranslatorEn) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorEnAuth struct{}

func (t *TranslatorEnAuth) Unauthorized() string {
	return "unauthorized"
}

func (t *TranslatorEnAuth) UserWithEmailNotFound() string {
	return "no user found with provided email"
}

func (t *TranslatorEnAuth) UserWithPhoneNumberNotFound() string {
	return "no user found with provided phone number"
}

func (t *TranslatorEnAuth) WrongPasswordWithEmailPassword() string {
	return "provided email or password is not correct"
}

func (t *TranslatorEnAuth) WrongPasswordWithPhoneNumberPassword() string {
	return "provided phone number or password is not correct"
}

func (t *TranslatorEnAuth) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorEnGalidator struct{}

func (t *TranslatorEnGalidator) MaxLength() string {
	return "must be at most $max characters in length"
}

func (t *TranslatorEnGalidator) MinLength() string {
	return "must be at least $min characters in length"
}

func (t *TranslatorEnGalidator) Phone() string {
	return "sent phone number is not valid"
}

func (t *TranslatorEnGalidator) Required() string {
	return "required"
}

func (t *TranslatorEnGalidator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorEnStatusCodes struct{}

func (t *TranslatorEnStatusCodes) BodyNotProvidedProperly() string {
	return "Body not provided properly"
}

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

type TranslatorEnUsers struct{}

func (t *TranslatorEnUsers) UserNotFound() string {
	return "user not found"
}

func (t *TranslatorEnUsers) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}
