package i18n_interfaces

import "fmt"

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

func (t *TranslatorAuth) EmailIsAlreadyVerified() string {
	return "ایمیل در حال حاضر فعال است"
}

func (t *TranslatorAuth) EmailNotFound() string {
	return "ایمیل یافت نشد"
}

func (t *TranslatorAuth) EmailVerified() string {
	return "ایمیل با موفقیت تایید شد"
}

func (t *TranslatorAuth) EmailVerifyCodeCoolDown(seconds int64) string {
	return fmt.Sprintf("برای درخواست تولید دوباره کد، %v ثانیه دیگر تلاش کنید", seconds)
}

func (t *TranslatorAuth) EmailVerifyCodeExpired() string {
	return "کد ایمیل منقضی شده است"
}

func (t *TranslatorAuth) EmailVerifyCodeSent() string {
	return "کد تایید ایمیل ارسال شد"
}

func (t *TranslatorAuth) EmailVerifyCodeTooManyRequests() string {
	return "تعداد درخواست های مجاز شما به انتها رسیده است"
}

func (t *TranslatorAuth) FirstRequestForVerifyCode() string {
	return "لطفا ابتدا درخواست ارسال کد را ارسال کنید"
}

func (t *TranslatorAuth) InvalidToken() string {
	return "توکن نامعتبر"
}

func (t *TranslatorAuth) PhoneNumberIsAlreadyVerified() string {
	return "شماره تماس در حال حاضر فعال است"
}

func (t *TranslatorAuth) PhoneNumberVerified() string {
	return "شماره تماس با موفقیت تایید شد"
}

func (t *TranslatorAuth) PhoneNumberVerifyCodeCoolDown(seconds int64) string {
	return fmt.Sprintf("برای درخواست تولید دوباره کد، %v ثانیه دیگر تلاش کنید", seconds)
}

func (t *TranslatorAuth) PhoneNumberVerifyCodeExpired() string {
	return "کد شماره تماس منقضی شده است"
}

func (t *TranslatorAuth) PhoneNumberVerifyCodeSent() string {
	return "کد تایید شماره تماس ارسال شد"
}

func (t *TranslatorAuth) PhoneNumberVerifyCodeTooManyRequests() string {
	return "تعداد درخواست های مجاز شما به انتها رسیده است"
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

func (t *TranslatorAuth) WrongCode(attempts int) string {
	return fmt.Sprintf("کد وارده تطبیق ندارد، %v تلاش دیگر مجاز است", attempts)
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

func (t *TranslatorGalidator) ImageType() string {
	return "تنها تصاویر png و jpg قابل قبول هستند"
}

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
