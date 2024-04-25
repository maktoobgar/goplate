package generated

type TranslatorI interface {
	Key() string
	Key2() string
	New() TranslatorNewI
	Other() string
	Voice() string
}

type TranslatorNewI interface {
	S() string
}
