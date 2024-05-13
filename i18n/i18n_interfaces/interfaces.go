package i18n_interfaces

import "reflect"

type TranslatorI interface {
	Auth() TranslatorAuthI
	Galidator() TranslatorGalidatorI
	HelloWorld() string
	StatusCodes() TranslatorStatusCodesI
	Users() TranslatorUsersI
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorAuthI interface {
	EmailIsAlreadyVerified() string
	EmailNotFound() string
	EmailVerified() string
	EmailVerifyCodeCoolDown(seconds int64) string
	EmailVerifyCodeExpired() string
	EmailVerifyCodeSent() string
	EmailVerifyCodeTooManyRequests() string
	FirstRequestForVerifyCode() string
	InvalidToken() string
	PhoneNumberIsAlreadyVerified() string
	PhoneNumberVerified() string
	PhoneNumberVerifyCodeCoolDown(seconds int64) string
	PhoneNumberVerifyCodeExpired() string
	PhoneNumberVerifyCodeSent() string
	PhoneNumberVerifyCodeTooManyRequests() string
	Unauthorized() string
	UserWithEmailNotFound() string
	UserWithPhoneNumberNotFound() string
	WrongCode(attempts int) string
	WrongPasswordWithEmailPassword() string
	WrongPasswordWithPhoneNumberPassword() string
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorGalidatorI interface {
	Choices() string
	ImageType() string
	MaxLength() string
	MinLength() string
	Phone() string
	Required() string
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorStatusCodesI interface {
	BodyNotProvidedProperly() string
	InternalServerError() string
	PageNotFound() string
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorUsersI interface {
	InvalidAvatar() string
	UserNotFound() string
	Translate(key string, optionalInputs ...[]any) string
}

func translate(instance any, key string, optionalInputs ...[]any) string {
	structType := reflect.TypeOf(instance)
	inputs := []any{instance}
	if len(optionalInputs) > 0 {
		inputs = append(inputs, optionalInputs[0]...)
	}

	// Iterate over all methods of the struct
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		if method.Name == key {
			reflectValues := make([]reflect.Value, len(inputs))
			for i, v := range inputs {
				reflectValues[i] = reflect.ValueOf(v)
			}
			return method.Func.Call(reflectValues)[0].String()
		}
	}
	return key
}