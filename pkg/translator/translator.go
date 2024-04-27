package translator

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"gopkg.in/yaml.v2"
)

var translatorContent = `package i18n

import "service/i18n/i18n_interfaces"

// Attribute 'lang' can be %s
func NewTranslator(lang string) i18n_interfaces.TranslatorI {
	if len(lang) >= 2 {
		lang = lang[:2]
	} else {
		lang = "%s"
	}

	if lang == "%s" {
		return &i18n_interfaces.Translator{}
	}%s

	return nil
}
`

var generatedContent = `package i18n_interfaces%s

%s`

var translateWithKeyContent = `func translate(instance any, key string, optionalInputs ...[]any) string {
	structType := reflect.TypeOf(instance)
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
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
}`

var translateWithKeyContentInStruct = `func (t *%s) Translate(key string, optionalInputs ...[]any) string {
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}
	return translate(t, key, inputs)
}`

func getLanguages(address, mainLang string) []string {
	langsPath := filepath.Join(address, "languages")
	if _, err := os.Stat(langsPath); err != nil {
		if err = os.Mkdir(langsPath, 509); err != nil {
			log.Fatalf("translator: can't create folder '%s', err: %v", langsPath, err)
		}
	}

	files, err := os.ReadDir(langsPath)
	if err != nil {
		log.Fatalf("translator: can't access folder '%s', err: %v", langsPath, err)
	}

	if len(files) == 0 {
		enFileAddr := filepath.Join(langsPath, mainLang+".yml")
		if _, err = os.Create(enFileAddr); err != nil {
			log.Fatalf("translator: can't read file '%s', err: %v", enFileAddr, err)
		}
	}

	languages := []string{}
	for _, file := range files {
		languageName, _ := strings.CutSuffix(file.Name(), ".yml")
		languages = append(languages, languageName)
	}
	return languages
}

func _getInOrder(words map[any]any) []any {
	keys := make([]any, 0, len(words))
	for key, value := range words {
		if val, ok := value.(map[any]any); ok {
			keys = append(keys, map[any]any{key: _getInOrder(val)})
		} else {
			keys = append(keys, key)
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		if first, ok := keys[i].(string); ok {
			if second, ok := keys[j].(string); ok {
				return first < second
			} else if second, ok := keys[j].(map[any]any); ok {
				for k := range second {
					return first < k.(string)
				}
			}
		} else if first, ok := keys[i].(map[any]any); ok {
			if second, ok := keys[j].(string); ok {
				for k := range first {
					return k.(string) < second
				}
			} else if second, ok := keys[j].(map[any]any); ok {
				for k1 := range first {
					for k2 := range second {
						return k1.(string) < k2.(string)
					}
				}
			}
		}
		return false
	})

	return keys
}

func _holdUpperCase(words map[any]any) map[any]any {
	for k, v := range words {
		key := k.(string)
		if cases.Title(language.English).String(key)[0] != key[0] {
			delete(words, key)
		} else {
			if value, ok := v.(map[any]any); ok {
				words[key] = _holdUpperCase(value)
			}
		}
	}

	return words
}

func getWords(address, mainLang string) (map[any]any, []any) {
	mainLangPath := filepath.Join(address, "languages/"+mainLang+".yml")
	words := make(map[any]any)

	yamlFile, err := os.ReadFile(mainLangPath)
	if err != nil {
		fmt.Printf("getWords: failed to read '%s', err: %v ", mainLangPath, err)
	}
	err = yaml.Unmarshal(yamlFile, words)
	if err != nil {
		fmt.Printf("getWords: failed to unmarshal yaml '%s', err: %v ", mainLangPath, err)
	}
	words = _holdUpperCase(words)

	return words, _getInOrder(words)
}

func _hasDeepSameKeys(words1, words2 map[any]any, lang string, beforeKeys string) {
	for word, v := range words1 {
		if _, ok := words2[word]; !ok {
			log.Fatalf("_hasDeepSameKeys: word '%s' doesn't exist in '%s' language in '%s' keys deep", word, lang, beforeKeys)
		}
		if v1, ok := v.(map[any]any); ok {
			if v2, ok := words2[word].(map[any]any); ok {
				_hasDeepSameKeys(v1, v2, lang, beforeKeys+"."+word.(string))
			} else {
				log.Fatalf("_hasDeepSameKeys: word '%s' must have a group of words but in '%s' language in '%s' keys deep it is just a string", word, lang, beforeKeys)
			}
		} else if _, ok := v.(string); ok {
			if _, ok := words2[word].(map[any]any); ok {
				log.Fatalf("_hasDeepSameKeys: word '%s' must be a string but in '%s' language in '%s' keys deep it has a group of words", word, lang, beforeKeys)
			}
		}
	}
}

func getWordsForEachLang(address string, words map[any]any, languages []string, mainLang string) (map[string]map[any]any, map[string][]any) {
	wordsForEachLangs := make(map[string]map[any]any, 0)
	wordsForEachLangsInOrder := make(map[string][]any, 0)
	for _, lang := range languages {
		if lang == mainLang {
			continue
		}
		mainLangPath := filepath.Join(address, "languages/"+lang+".yml")
		_words := make(map[any]any, 0)

		yamlFile, err := os.ReadFile(mainLangPath)
		if err != nil {
			fmt.Printf("yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, _words)
		if err != nil {
			fmt.Printf("Unmarshal: %v", err)
		}

		_words = _holdUpperCase(_words)
		_wordsKeysInOrder := _getInOrder(_words)

		_hasDeepSameKeys(words, _words, lang, "current_level")

		wordsForEachLangs[lang] = _words
		wordsForEachLangsInOrder[lang] = _wordsKeysInOrder
	}

	return wordsForEachLangs, wordsForEachLangsInOrder
}

func getInterfacesStructs(words map[any]any, wordsKeysInOrder []any, inputs map[any]any, structKey, structKeySimple string) (string, string) {
	re, _ := regexp.Compile(`{(\w+):(int|string)}`)
	reWithoutType, _ := regexp.Compile(`{(\w+)}`)
	oneKeyValueI := "	%s(%s) %s\n"
	oneKeyValueFunc := "\nfunc (t *%s) %s(%s) %s {\n\treturn %s\n}\n"
	singleInterface := ""
	singleStruct := ""
	interfaces := []string{}
	structs := []string{}

	for _, complexOrValue := range wordsKeysInOrder {
		if key, ok := complexOrValue.(string); ok {
			nextFinalInputs, ok := inputs[key].(map[any]any)
			if !ok {
				nextFinalInputs = map[any]any{}
			}

			arrs := ""
			fmtArrs := ""
			keysInOrder, ok := nextFinalInputs["in_order_keys"].([]string)
			if !ok {
				keysInOrder = []string{}
			}

			for _, key := range keysInOrder {
				if arrs == "" {
					arrs = fmt.Sprintf("%s %s", key, nextFinalInputs[key])
				} else {
					arrs += fmt.Sprintf(", %s %s", key, nextFinalInputs[key])
				}
			}

			matches := re.FindAllStringSubmatch(words[key].(string), int(math.Inf(1)))
			for _, match := range matches {
				if _, ok := nextFinalInputs[match[1]]; ok {
					if fmtArrs == "" {
						fmtArrs = match[1]
					} else {
						fmtArrs += fmt.Sprintf(", %s", match[1])
					}
				}
			}
			matches = reWithoutType.FindAllStringSubmatch(words[key].(string), int(math.Inf(1)))
			for _, match := range matches {
				if _, ok := nextFinalInputs[match[1]]; ok {
					if fmtArrs == "" {
						fmtArrs = match[1]
					} else {
						fmtArrs += fmt.Sprintf(", %s", match[1])
					}
				}
			}

			valueFormatted := reWithoutType.ReplaceAllString(re.ReplaceAllString(words[key].(string), "%v"), "%v")
			if fmtArrs != "" {
				valueFormatted = fmt.Sprintf("fmt.Sprintf(\"%s\", %s)", valueFormatted, fmtArrs)
			} else {
				valueFormatted = fmt.Sprintf("\"%s\"", valueFormatted)
			}

			singleInterface += fmt.Sprintf(oneKeyValueI, complexOrValue, arrs, "string")
			singleStruct += fmt.Sprintf(oneKeyValueFunc, structKey, complexOrValue, arrs, "string", valueFormatted)
		} else if complex, ok := complexOrValue.(map[any]any); ok {
			for k, v := range complex {
				var justKey = structKey + k.(string)
				var justSimpleKey = structKeySimple + k.(string)
				singleInterface += fmt.Sprintf(oneKeyValueI, k.(string), "", justKey+"I")
				singleStruct += fmt.Sprintf(oneKeyValueFunc, structKey, k.(string), "", justSimpleKey+"I", fmt.Sprintf("&%s{}", justKey))
				if words[k.(string)] != nil {
					nextInputs, ok := inputs[k.(string)].(map[any]any)
					if !ok {
						nextInputs = map[any]any{}
					}
					oneInterface, oneStruct := getInterfacesStructs(words[k.(string)].(map[any]any), v.([]any), nextInputs, justKey, justSimpleKey)
					interfaces = append(interfaces, oneInterface)
					structs = append(structs, oneStruct)
				}
			}
		}
	}

	singleInterface += fmt.Sprintf(oneKeyValueI, "Translate", "key string, optionalInputs ...[]any", "string")
	singleStruct = "type " + structKey + " struct{}\n" + singleStruct + "\n" + fmt.Sprintf(translateWithKeyContentInStruct, structKey) + "\n"
	singleInterface = "type " + structKey + "I" + " interface {\n" + singleInterface + "}\n"
	for i := 0; i < len(structs); i++ {
		singleStruct += "\n" + structs[i]
	}
	for i := 0; i < len(interfaces); i++ {
		singleInterface += "\n" + interfaces[i]
	}

	return singleInterface, singleStruct
}

func createInterfaces(address, interfaces string) {
	interfacePath := filepath.Join(address, "i18n_interfaces/interfaces.go")
	content := fmt.Sprintf(generatedContent, "\n\nimport \"reflect\"", interfaces+"\n"+translateWithKeyContent)

	file, err := os.Create(interfacePath)
	if err != nil {
		log.Fatalf("translator: error creating file '%s', err: %s\n", interfacePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("translator: failed to write to file '%s', err: %s\n", interfacePath, err)
	}
}

func createStructs(interfacePath, structs string, addImportFmt bool) {
	content := ""
	if addImportFmt {
		content = fmt.Sprintf(generatedContent, "\n\nimport \"fmt\"", structs)
	} else {
		content = fmt.Sprintf(generatedContent, "", structs)
	}

	file, err := os.Create(interfacePath)
	if err != nil {
		log.Fatalf("translator: error creating file '%s', err: %s\n", interfacePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("translator: failed to write to file '%s', err: %s\n", interfacePath, err)
	}
}

func createTranslator(address string, languages []string, mainLang string) {
	interfacePath := filepath.Join(address, "translator.go")
	elseIfI := " else if lang == \"%s\" {\n\t\treturn &i18n_interfaces.%s{}\n\t}"

	langsString := ""
	elseIfBlock := ""
	for i, lang := range languages {
		if mainLang != lang {
			elseIfBlock += fmt.Sprintf(elseIfI, lang, fmt.Sprintf("Translator%s", cases.Title(language.English).String(lang)))
		}
		if i == 0 {
			langsString = lang
			continue
		}
		langsString += fmt.Sprintf(", %s", lang)
	}

	content := fmt.Sprintf(translatorContent, langsString, mainLang, mainLang, elseIfBlock)
	file, err := os.Create(interfacePath)
	if err != nil {
		log.Fatalf("translator: error creating file '%s', err: %s\n", interfacePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("translator: failed to write to file '%s', err: %s\n", interfacePath, err)
	}
}

func returnMethodInputs(words map[any]any) map[any]any {
	output := make(map[any]any)
	re, _ := regexp.Compile(`{(\w+):(int|float|string)}`)
	for word, value := range words {
		inputs := make(map[any]any)
		inputsInOrder := make([]string, 0)
		if v, ok := value.(string); ok {
			matches := re.FindAllStringSubmatch(v, int(math.Inf(1)))
			for _, match := range matches {
				inputs[match[1]] = match[2]
				inputsInOrder = append(inputsInOrder, match[1])
			}
			if len(inputsInOrder) > 0 {
				inputs["in_order_keys"] = inputsInOrder
			}
		} else if v, ok := value.(map[any]any); ok {
			inputs = returnMethodInputs(v)
		}
		if len(inputs) > 0 {
			output[word] = inputs
		}
	}

	return output
}

func GenerateCode(address, mainLang string) {
	if _, err := os.Stat(address); err != nil {
		if err = os.Mkdir(address, 509); err != nil {
			log.Fatalf("translator: can't create folder '%s', err: %v", address, err)
		}
	}

	generatedFolder := filepath.Join(address, "i18n_interfaces")
	if _, err := os.Stat(generatedFolder); err != nil {
		if err = os.Mkdir(generatedFolder, 509); err != nil {
			log.Fatalf("translator: can't create folder '%s', err: %v", generatedFolder, err)
		}
	}

	languages := getLanguages(address, mainLang)
	words, wordsKeysInOrder := getWords(address, mainLang)
	wordsForEachLangs, wordsForEachLangsInOrder := getWordsForEachLang(address, words, languages, mainLang)

	inputs := returnMethodInputs(words)
	interfaces, structs := getInterfacesStructs(words, wordsKeysInOrder, inputs, "Translator", "Translator")
	createStructs(filepath.Join(address, fmt.Sprintf("i18n_interfaces/%s.go", mainLang)), structs, len(inputs) > 0)
	createInterfaces(address, interfaces)
	for lang := range wordsForEachLangs {
		_, structs := getInterfacesStructs(wordsForEachLangs[lang], wordsForEachLangsInOrder[lang], inputs, "Translator"+cases.Title(language.English).String(lang), "Translator")
		createStructs(filepath.Join(address, fmt.Sprintf("i18n_interfaces/%s.go", lang)), structs, len(inputs) > 0)
	}
	createTranslator(address, languages, mainLang)
}
