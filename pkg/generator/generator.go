package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golodash/godash/strings"
)

const handlerStructure = `package %s

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/repositories"

	"github.com/kataras/iris/v12"
)

type %sReq struct {
}

var %sValidator = g.Galidator.Validator(%sReq{})

func %s(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*%sReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	_, _, _, _ = db, req, user, translator
}
`

func GenerateNewHandler(name, packageName, address string) {
	var fileAddress string
	if len(packageName) > 0 {
		packageName = strings.SnakeCase(packageName) + "_handlers"
	} else {
		packageName = "handlers"
	}
	if packageName != "handlers" {
		fileAddress = filepath.Join(address, "handlers", packageName, strings.SnakeCase(name)+".go")
	} else {
		fileAddress = filepath.Join(address, packageName, strings.SnakeCase(name)+".go")
	}

	name = strings.PascalCase(name)
	packageName = strings.SnakeCase(packageName)
	content := fmt.Sprintf(handlerStructure, packageName, name, name, name, name, name)

	file, err := os.Create(fileAddress)
	if err != nil {
		log.Fatalf("generator: error opening file '%s', err: %s\n", fileAddress, err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("generator: failed to write to file '%s', err: %s\n", fileAddress, err)
	}
}
