package i18n_interfaces

import "fmt"

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

func (t *TranslatorEnAuth) EmailIsAlreadyVerified() string {
	return "email is already verified"
}

func (t *TranslatorEnAuth) EmailNotFound() string {
	return "email not found"
}

func (t *TranslatorEnAuth) EmailVerified() string {
	return "email verified"
}

func (t *TranslatorEnAuth) EmailVerifyCodeCoolDown(seconds int64) string {
	return fmt.Sprintf("can't generate code Now, try %v seconds later", seconds)
}

func (t *TranslatorEnAuth) EmailVerifyCodeExpired() string {
	return "provided email verification code is expired"
}

func (t *TranslatorEnAuth) EmailVerifyCodeSent() string {
	return "email verification code sent"
}

func (t *TranslatorEnAuth) EmailVerifyCodeTooManyRequests() string {
	return "you have reached the limit for number of allowed requests"
}

func (t *TranslatorEnAuth) FirstRequestForVerifyCode() string {
	return "please first request for a verification code"
}

func (t *TranslatorEnAuth) InvalidToken() string {
	return "invalid token"
}

func (t *TranslatorEnAuth) PhoneNumberIsAlreadyVerified() string {
	return "phone number is already verified"
}

func (t *TranslatorEnAuth) PhoneNumberVerified() string {
	return "phone number verified"
}

func (t *TranslatorEnAuth) PhoneNumberVerifyCodeCoolDown(seconds int64) string {
	return fmt.Sprintf("can't generate code Now, try %v seconds later", seconds)
}

func (t *TranslatorEnAuth) PhoneNumberVerifyCodeExpired() string {
	return "provided phone number verification code is expired"
}

func (t *TranslatorEnAuth) PhoneNumberVerifyCodeSent() string {
	return "phone number verification code sent"
}

func (t *TranslatorEnAuth) PhoneNumberVerifyCodeTooManyRequests() string {
	return "you have reached the limit for number of allowed requests"
}

func (t *TranslatorEnAuth) Unauthorized() string {
	return "unauthorized"
}

func (t *TranslatorEnAuth) UserWithEmailNotFound() string {
	return "no user found with provided email"
}

func (t *TranslatorEnAuth) UserWithPhoneNumberNotFound() string {
	return "no user found with provided phone number"
}

func (t *TranslatorEnAuth) WrongCode(attempts int) string {
	return fmt.Sprintf("provided code doesn't match, you are allowed %v more attempts", attempts)
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

func (t *TranslatorEnGalidator) Choices() string {
	return "acceptable choices are $choices"
}

func (t *TranslatorEnGalidator) ImageType() string {
	return "only png and jpg image formats are allowed"
}

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

func (t *TranslatorEnUsers) InvalidAvatar() string {
	return "invalid avatar"
}

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
