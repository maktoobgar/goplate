package generated

type TranslatorI interface {
	Key() string
	Key2(message int, new string) string
	New() TranslatorNewI
	Other() string
	Voice() string
}

type TranslatorNewI interface {
	S(parameter int) string
}
