package i18n_interfaces

type Translator struct{}

func (t *Translator) Auth() TranslatorAuthI {
	return &TranslatorAuth{}
}

func (t *Translator) Galidator() TranslatorGalidatorI {
	return &TranslatorGalidator{}
}

func (t *Translator) HelloWorld() string {
	return "درود"
}

func (t *Translator) StatusCodes() TranslatorStatusCodesI {
	return &TranslatorStatusCodes{}
}

func (t *Translator) Users() TranslatorUsersI {
	return &TranslatorUsers{}
}

func (t *Translator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorAuth struct{}

func (t *TranslatorAuth) InvalidToken() string {
	return "توکن نامعتبر"
}

func (t *TranslatorAuth) Unauthorized() string {
	return "احراز هویت نشده"
}

func (t *TranslatorAuth) UserWithEmailNotFound() string {
	return "کاربری با ایمیل وارده یافت نشد"
}

func (t *TranslatorAuth) UserWithPhoneNumberNotFound() string {
	return "کاربری با شماره تماس وارده یافت نشد"
}

func (t *TranslatorAuth) WrongPasswordWithEmailPassword() string {
	return "ایمیل یا پسورد وارد شده صحیح نمیباشد"
}

func (t *TranslatorAuth) WrongPasswordWithPhoneNumberPassword() string {
	return "شماره تماس یا پسورد وارد شده صحیح نمیباشد"
}

func (t *TranslatorAuth) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorGalidator struct{}

func (t *TranslatorGalidator) MaxLength() string {
	return "حداکثر باید دارای طول $max کاراکتر باشد"
}

func (t *TranslatorGalidator) MinLength() string {
	return "حداقل باید دارای طول $min کاراکتر باشد"
}

func (t *TranslatorGalidator) Phone() string {
	return "شماره تماس ارسالی معتبر نمیباشد"
}

func (t *TranslatorGalidator) Required() string {
	return "اجباری"
}

func (t *TranslatorGalidator) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}

type TranslatorStatusCodes struct{}

func (t *TranslatorStatusCodes) BodyNotProvidedProperly() string {
	return "درخواست دارای خطاست"
}

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

type TranslatorUsers struct{}

func (t *TranslatorUsers) InvalidAvatar() string {
	return "محتوای آواتار صحیح ارسال نشده است"
}

func (t *TranslatorUsers) UserNotFound() string {
	return "کاربر یافت نشد"
}

func (t *TranslatorUsers) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}
