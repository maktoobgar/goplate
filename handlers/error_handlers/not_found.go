package error_handlers

import (
	"service/pkg/errors"

	"github.com/kataras/iris/v12"
)

func NotFound(ctx iris.Context) {
	panic(errors.New(errors.NotFoundStatus, "PageNotFound", "page not found"))
}
