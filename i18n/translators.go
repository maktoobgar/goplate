package i18n

import "service/i18n/interfaces"

type translator struct {
	dictionary interfaces.Words
}

func (t *translator) Other() string {
	return t.dictionary.Other
}

func (t *translator) Voice() string {
	return t.dictionary.Voice
}

func (t *translator) New() string {
	return t.dictionary.New
}

func (t *translator) Key() string {
	return t.dictionary.Key
}

func (t *translator) Key2() string {
	return t.dictionary.Key2
}
