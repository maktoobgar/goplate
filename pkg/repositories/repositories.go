package repositories

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateRepositories(address string) {
	// Execute sqlc
	execOutput := exec.Command("sqlc", "generate")
	if execOutput.Err == nil {
		execOutput.Err = execOutput.Wait()
	} else {
		log.Fatalf("repositories: can't execute '%v', err: %v", strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(execOutput.Args), "[", ""), "]", ""), execOutput.Err)
	}

	// Read query.sql.go file
	queriesAddress := filepath.Join(address, "query.sql.go")
	content, err := os.ReadFile(queriesAddress)
	if err != nil {
		log.Fatalf("repositories: can't read file '%s', err: %v", queriesAddress, err)
	}

	// Import adding part
	re, _ := regexp.Compile(`import\s*\((\s*((\w+ \"(\w|\/)*\")|(\"(\w|\/)*\")))+\s*\)`)
	importContent := strings.Replace(re.FindString(string(content)), ")", "", 1)
	if !strings.Contains(importContent, "\"service/pkg/errors\"") {
		importContent += "\t\"service/pkg/errors\"\n"
	}
	if !strings.Contains(importContent, "\"service/global\"") {
		importContent += "\t\"service/global\"\n"
	}
	if !strings.Contains(importContent, "\"service/i18n/i18n_interfaces\"") {
		importContent += "\t\"service/i18n/i18n_interfaces\"\n"
	}
	importContent += ")"
	output := re.ReplaceAllString(string(content), importContent)
	file, err := os.Create(queriesAddress)
	if err != nil {
		log.Fatalf("repositories: error opening file '%s', err: %s\n", queriesAddress, err)
	}
	defer file.Close()

	// Error panic before return statement
	output = strings.ReplaceAll(output, ")\n\treturn i, err", ")\n\tif err != nil {\n\t\tpanic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))\n\t}\n\treturn i, err")

	// Add translator to the start of it
	output = strings.ReplaceAll(output, "{\n\trow", "{\n\ttranslator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)\n\trow")

	_, err = file.WriteString(output)
	if err != nil {
		log.Fatalf("repositories: failed to write to file '%s', err: %s\n", queriesAddress, err)
	}
}
