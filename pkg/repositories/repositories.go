package repositories

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const decoderContent = `package repositories

import (
	"encoding/json"
	"service/pkg/copier"
	"time"
)
%s
`

func reformatFile(address string) {
	execOutput := exec.Command("gofmt", "-w", address)
	err := execOutput.Run()
	if execOutput.Err != nil || err != nil {
		if execOutput.Err != nil {
			err = execOutput.Err
		}
		log.Panicf("repositories: can't execute '%v', err: %v", strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(execOutput.Args), "[", ""), "]", ""), err)
	}
}

func GenerateRepositories(address string) {
	// Execute sqlc
	execOutput := exec.Command("sqlc", "generate", fmt.Sprintf("-f=%s", filepath.Join(address, "sqlc.yml")))
	err := execOutput.Run()
	if execOutput.Err != nil || err != nil {
		if execOutput.Err != nil {
			err = execOutput.Err
		}
		log.Panicf("repositories: can't execute '%v', err: %v", strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(execOutput.Args), "[", ""), "]", ""), err)
	}

	{ // query.sql.go file
		// Read query.sql.go file
		dir, err := os.Open(address)
		if err != nil {
			log.Panicf("repositories: can't open folder '%s', err: %v", address, err)
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			log.Panicf("repositories: can't read folder content '%s', err: %v", address, err)
		}
		// Iterate over all files which has generated queries and mutate their functions
		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".sql.go") {
				continue
			}
			queriesAddress := filepath.Join(address, file.Name())
			content, err := os.ReadFile(queriesAddress)
			if err != nil {
				log.Panicf("repositories: can't read file '%s', err: %v", queriesAddress, err)
			}

			// Import adding part
			re, _ := regexp.Compile(`import\s*\((\s*((\w+ \"(\w|\/)*\")|(\"(\w|\/)*\")))+\s*\)`)
			importContent := strings.Replace(re.FindString(string(content)), ")", "", 1)
			if !strings.Contains(importContent, "\"service/pkg/errors\"") {
				importContent += "\"service/pkg/errors\"\n"
			}
			if !strings.Contains(importContent, "\"service/global\"") {
				importContent += "\"service/global\"\n"
			}
			if !strings.Contains(importContent, "\"service/i18n/i18n_interfaces\"") {
				importContent += "\"service/i18n/i18n_interfaces\"\n"
			}
			if !strings.Contains(importContent, "\"database/sql\"") {
				importContent += "\"database/sql\"\n"
			}
			importContent += ")"
			output := re.ReplaceAllString(string(content), importContent)
			file, err := os.Create(queriesAddress)
			if err != nil {
				log.Panicf("repositories: error opening file '%s', err: %s\n", queriesAddress, err)
			}
			defer file.Close()

			// Error panic before return statement
			output = strings.ReplaceAll(output, ")\n\treturn i, err", ")\nif err != nil && err != sql.ErrNoRows {panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))}\nreturn i, err")
			output = strings.ReplaceAll(output, "return nil, err", "panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))")

			// Add translator to the start of it
			output = strings.ReplaceAll(output, "{\n\trow", "{translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI);row")

			_, err = file.WriteString(output)
			if err != nil {
				log.Panicf("repositories: failed to write to file '%s', err: %s\n", queriesAddress, err)
			}

			reformatFile(queriesAddress)
		}
	}

	{ // create decoder.go file
		decoderPath := filepath.Join(address, "decoder.go")
		content, err := os.ReadFile(filepath.Join(address, "models.go"))
		if err != nil {
			log.Panicf("repositories: can't read file '%s', err: %v", content, err)
		}

		// `type \w* struct {(\s|\d|\w|\:|\"|`|\.)*}`
		re, _ := regexp.Compile(`type (\w*) struct {((\s|\d|\w|\:|\"|` + "`" + `|\.)*)}`)
		const structureStringStructure = "\ntype %s struct {%s}\n%s"
		const functionStructure = "\nfunc (u %s) MarshalJSON() ([]byte, error) {\n\treturn json.Marshal(copier.Copy(&%s{}, &u))\n}"
		structures := re.FindAllStringSubmatch(string(content), int(math.Inf(1)))
		structuresString := ""
		for i := range structures {
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullString", "string")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullBool", "bool")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullByte", "byte")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullFloat64", "float64")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullInt16", "int16")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullInt32", "int32")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullInt64", "int64")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullString", "string")
			structures[i][2] = strings.ReplaceAll(structures[i][2], "sql.NullTime", "time.Time")
			originalStructName := structures[i][1]
			structures[i][1] = strings.ToLower(structures[i][1][0:1]) + structures[i][1][1:]

			if i == 0 {
				structuresString += fmt.Sprintf(structureStringStructure, structures[i][1], structures[i][2], fmt.Sprintf(functionStructure, originalStructName, structures[i][1]))
			} else {
				structuresString += fmt.Sprintf("\n"+structureStringStructure, structures[i][1], structures[i][2], fmt.Sprintf(functionStructure, originalStructName, structures[i][1]))
			}
		}

		output := fmt.Sprintf(decoderContent, structuresString)

		file, err := os.Create(decoderPath)
		if err != nil {
			log.Panicf("repositories: error opening file '%s', err: %s\n", decoderPath, err)
		}
		defer file.Close()

		_, err = file.WriteString(output)
		if err != nil {
			log.Panicf("repositories: failed to write to file '%s', err: %s\n", decoderPath, err)
		}

		// Reformat the file
		reformatFile(decoderPath)
	}
}
