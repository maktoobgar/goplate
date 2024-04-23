package translator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var interfaceContent = `package interfaces

type Words struct {
%s
}

type WordsI interface {
%s
}

type I18N interface {
%s
}
`

var translatorsContent = `package i18n

import "service/i18n/interfaces"

type translator struct {
	dictionary interfaces.Words
}
`

func getLanguages(address, mainLang string) []string {
	langsPath := filepath.Join(address, "languages")
	if _, err := os.Stat(langsPath); err != nil {
		if err = os.Mkdir(langsPath, 509); err != nil {
			log.Fatalf("translator: can't create folder '%s', err: %v", langsPath, err)
		}
	}

	files, err := os.ReadDir(langsPath)
	if err != nil {
		log.Fatalf("translator: can't read folder '%s', err: %v", langsPath, err)
	}

	if len(files) == 0 {
		enFileAddr := filepath.Join(langsPath, mainLang+".go")
		if _, err = os.Create(enFileAddr); err != nil {
			log.Fatalf("translator: can't read folder '%s', err: %v", enFileAddr, err)
		}
	}

	languages := []string{}
	for _, file := range files {
		languageName, _ := strings.CutSuffix(file.Name(), ".go")
		languages = append(languages, languageName)
	}
	return languages
}

func getWords(address, mainLang string) map[string]string {
	mainLangPath := filepath.Join(address, "languages/"+mainLang+".go")
	body, err := os.ReadFile(mainLangPath)
	if err != nil {
		log.Fatalf("translator: can't read file '%s', err: %v", mainLangPath, err)
	}

	if err != nil {
		log.Fatalf("translator: can't read file '%s', err: %v", mainLangPath, err)
	}
	content := string(body)

	pattern := `interfaces\.Words\s*{(.|\n)*\n}.*`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(content)
	wordsRaw := ""
	if len(matches) > 0 {
		wordsRaw = regexp.MustCompile("\n}.*").ReplaceAllString(regexp.MustCompile(`interfaces\.Words\s*{[\n ]*`).ReplaceAllString(matches[0], ""), "")
	} else {
		return nil
	}

	var keyValues = map[string]string{}
	for _, rawKeyValue := range strings.Split(wordsRaw, "\n") {
		rawKeyValue = strings.TrimSpace(rawKeyValue)
		var keyValue = strings.Split(rawKeyValue, ":")
		if len(keyValue) == 2 {
			var key = cases.Title(language.English, cases.Compact).String(strings.Trim(strings.TrimSpace(keyValue[0]), "\""))
			var value = strings.Trim(strings.TrimSuffix(strings.TrimSpace(keyValue[1]), ","), "\"")

			keyValues[key] = value
		}
	}
	return keyValues
}

func getLangVariables(address string, languages []string) map[string]string {
	pattern := `var\s+\w+\s*=\s*interfaces.Words\s*{`
	re := regexp.MustCompile(pattern)
	varNames := map[string]string{}
	for _, lang := range languages {
		langAddr := filepath.Join(address, "languages/"+lang+".go")
		body, err := os.ReadFile(langAddr)
		if err != nil {
			log.Fatalf("translator: can't read file '%s', err: %v", langAddr, err)
		}
		content := string(body)
		matches := re.FindStringSubmatch(content)
		varNames[lang] = strings.TrimSpace(strings.TrimPrefix(strings.Split(matches[0], "=")[0], "var"))
	}
	return varNames
}

func createInterface(address string, languages []string, words map[string]string) string {
	interfacePath := filepath.Join(address, "interfaces/interface.go")
	_ = interfacePath
	oneKeyValue := "	%s string\n"
	oneKeyValueI := "	%s() string\n"
	rows := ""
	rowsI := ""
	for key := range words {
		var row = fmt.Sprintf(oneKeyValue, key)
		var rowI = fmt.Sprintf(oneKeyValueI, key)
		rows += row
		rowsI += rowI
	}
	rows = strings.TrimRightFunc(rows, unicode.IsSpace)
	rowsI = strings.TrimRightFunc(rowsI, unicode.IsSpace)

	oneKeyValueLang := "	%s() WordsI\n"
	langRows := ""
	for _, key := range languages {
		var row = fmt.Sprintf(oneKeyValueLang, strings.ToUpper(key))
		langRows += row
	}
	langRows = strings.TrimRightFunc(langRows, unicode.IsSpace)

	content := fmt.Sprintf(interfaceContent, rows, rowsI, langRows)
	file, err := os.Create(interfacePath)
	if err != nil {
		log.Fatalf("translator: error creating file '%s', err: %s\n", interfacePath, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("translator: failed to write to file '%s', err: %s\n", interfacePath, err)
	}
	return content
}

func generateActualFunctions(address string, words map[string]string) {
	translatorsPath := filepath.Join(address, "translators.go")
	oneFunc := "\nfunc (t *translator) %s() string {\n\treturn t.dictionary.%s\n}\n"
	functions := ""
	for key := range words {
		functions += fmt.Sprintf(oneFunc, key, key)
	}
	file, err := os.Create(translatorsPath)
	if err != nil {
		log.Fatalf("translator: error creating file '%s', err: %s\n", translatorsPath, err)
	}
	defer file.Close()

	_, err = file.WriteString(translatorsContent + functions)
	if err != nil {
		log.Fatalf("translator: failed to write to file '%s', err: %s\n", translatorsPath, err)
	}
}

func GenerateLanguages(address, mainLang string) {
	if _, err := os.Stat(address); err != nil {
		if err = os.Mkdir(address, 509); err != nil {
			log.Fatalf("translator: can't create folder '%s', err: %v", address, err)
		}
	}

	languages := getLanguages(address, mainLang)
	words := getWords(address, mainLang)
	langVariables := getLangVariables(address, languages)
	createInterface(address, languages, words)
	generateActualFunctions(address, words)
	_ = langVariables
}
