package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	godashStrings "github.com/golodash/godash/strings"
)

const handlerStructure = `package %s

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/repositories"
	"service/validators"

	"github.com/kataras/iris/v12"
)

type %sReq struct {
}

type %sRes struct {
}

var %sValidator = validators.Generator.Validator(%sReq{})

func %s(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*%sReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries
}
`

func GenerateNewHandler(name, packageName, address string) {
	packageNames := strings.Split(packageName, ".")
	if len(packageNames) > 1 {
		firstAddress := godashStrings.SnakeCase(packageNames[0])
		address = filepath.Join(address, firstAddress+"_handlers")
		packageNames = packageNames[1:]
		if _, err := os.Stat(address); os.IsNotExist(err) {
			if err = os.MkdirAll(address, os.ModePerm); err != nil {
				log.Panicf("generator: error full creating folder '%s', err: %s\n", address, err)
			}
		}
		GenerateNewHandler(name, strings.Join(packageNames, "."), address)
		return
	}

	packageName = godashStrings.SnakeCase(packageName + "_handlers")
	address = filepath.Join(address, packageName)

	if _, err := os.Stat(address); os.IsNotExist(err) {
		if err = os.MkdirAll(address, os.ModePerm); err != nil {
			log.Panicf("generator: error full creating folder '%s', err: %s\n", address, err)
		}
	}

	fileAddress := filepath.Join(address, godashStrings.SnakeCase(name)+".go")

	name = godashStrings.PascalCase(name)
	content := fmt.Sprintf(handlerStructure, packageName, name, name, name, name, name, name)

	file, err := os.Create(fileAddress)
	if err != nil {
		log.Panicf("generator: error opening file '%s', err: %s\n", fileAddress, err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Panicf("generator: failed to write to file '%s', err: %s\n", fileAddress, err)
	}
}
