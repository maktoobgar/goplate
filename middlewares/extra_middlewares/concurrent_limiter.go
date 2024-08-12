package extra_middlewares

import (
	"github.com/kataras/iris/v12"
)

var sem chan struct{} = nil

func ConcurrentLimiter(maxConcurrentRequests int) iris.Handler {
	if sem == nil {
		sem = make(chan struct{}, maxConcurrentRequests)
	}

	return func(ctx iris.Context) {
		sem <- struct{}{} // Acquire a semaphore slot
		defer func() {
			<-sem // Release the semaphore slot
		}()

		ctx.Next()
	}
}
