package interfaces

type Words struct {
	Key string
	Key2 string
	Other string
	Voice string
	New string
}

type WordsI interface {
	Key() string
	Key2() string
	Other() string
	Voice() string
	New() string
}

type I18N interface {
	EN() WordsI
	FA() WordsI
}
