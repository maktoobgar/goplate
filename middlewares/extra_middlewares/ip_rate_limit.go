package extra_middlewares

import (
	"fmt"
	g "service/global"
	"service/pkg/errors"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"golang.org/x/time/rate"
)

type Limiter struct {
	limiter    *rate.Limiter
	lastAccess time.Time
}

var (
	rateLimiterMap    = map[string]*Limiter{}
	rateLimiterMapKey = &sync.Mutex{}

	// Checks if `lastAccess` is 5 minutes behind current time and if so, delete it
	IpRateLimitGarbageCollector = &g.CronJob{
		Job: func() {
			past5min := time.Now().Add(-5 * time.Minute)
			for k, v := range rateLimiterMap {
				if v.lastAccess.Before(past5min) {
					rateLimiterMapKey.Lock()
					defer rateLimiterMapKey.Unlock()

					delete(rateLimiterMap, k)
				}
			}
		},
	}
)

func IpRateLimit(ctx iris.Context) {
	ip := ctx.RemoteAddr()
	limiter, exists := rateLimiterMap[ip]
	if !exists {
		rateLimiterMapKey.Lock()
		defer rateLimiterMapKey.Unlock()
		limiter = &Limiter{
			// 1.0/5.0 requests in every second gets accepted which means
			// 1 request in 5 seconds
			// What that 3 means? figure it out cause I have no freaking idea
			// and I do not care too at the time but this combination seems
			// fine too me and I like it
			limiter:    rate.NewLimiter(rate.Limit(1.0/5.0), 3),
			lastAccess: time.Now(),
		}
		rateLimiterMap[ip] = limiter
	}

	limiter.lastAccess = time.Now()
	if !limiter.limiter.Allow() {
		panic(errors.New(errors.TooManyRequests, "TooManyRequests", fmt.Sprintf("too many attempts for: %s", ip)))
	}

	ctx.Next()
}
