package extra_middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
	"sync"

	"service/dto"
	g "service/global"

	"service/pkg/errors"
	"service/pkg/translator"

	"github.com/kataras/iris/v12"
)

func Panic(ctx iris.Context) {
	translate := ctx.Values().Get(g.TranslateKey).(translator.TranslatorFunc)

	defer func() {
		errInterface := recover()
		if errInterface == nil {
			return
		}

		writerLockInterface := ctx.Values().Get(g.WriterLock)
		var writerLock *sync.Mutex = nil
		if writerLockInterface != nil {
			writerLock = writerLockInterface.(*sync.Mutex)
			writerLock.Lock()
			defer writerLock.Unlock()
		}

		closedWriter := false
		closedWriterInterface := ctx.Values().Get(g.ClosedWriter)
		if closedWriterInterface != nil {
			closedWriter = closedWriterInterface.(bool)
		}
		if !closedWriter {
			if err, ok := errInterface.(error); ok && errors.IsServerError(err) {
				castedError := errors.CastError(err)
				code, message, _, errors := errors.HttpError(err)
				if castedError.HasStackError() {
					stack := castedError.GetStack()
					if code == 500 {
						g.Logger.Panic(errInterface, ctx.Request(), stack)
					}
					res := dto.PanicResponse{
						Message: translate(message),
						Code:    code,
						Errors:  errors,
					}
					if g.CFG.Debug {
						log.Println(err)
					}
					ctx.StopWithJSON(res.Code, res)
				} else {
					res := dto.PanicResponse{
						Message: translate(message),
						Code:    code,
						Errors:  errors,
					}
					if g.CFG.Debug {
						log.Println(err)
					}
					ctx.StopWithJSON(code, res)
				}
			} else {
				stack := string(debug.Stack())
				g.Logger.Panic(errInterface, ctx.Request(), stack)
				res := dto.PanicResponse{
					Message: translate("InternalServerError"),
					Code:    http.StatusInternalServerError,
					Errors:  nil,
				}
				ctx.StopWithJSON(res.Code, res)
			}
		}
		closedWriter = true
		ctx.Values().Set(g.ClosedWriter, closedWriter)
	}()

	ctx.Next()
}
