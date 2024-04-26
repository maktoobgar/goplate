package routes

import (
	g "service/global"
	"service/handlers"
	"service/handlers/error_handlers"
	"service/middlewares/extra_middlewares"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
)

// Applies all necessary middlewares
func addBasicMiddlewares(app *iris.Application) {
	// Copression
	app.UseRouter(iris.Compression)

	// Add Cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(g.CFG.AllowOrigins, ","),
		AllowedHeaders:   strings.Split(g.CFG.AllowHeaders, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	})
	app.WrapRouter(c.ServeHTTP)

	// Translator
	app.Use(extra_middlewares.Translator)

	// Panic
	app.Use(extra_middlewares.Panic)

	// Creates a db for every db operation
	app.Use(extra_middlewares.CreateDbInstance)
}

func addBasicErrorRoutes(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, extra_middlewares.Translator, extra_middlewares.Panic, error_handlers.NotFound)
}

func HTTP(app *iris.Application) {
	addBasicMiddlewares(app)
	addBasicErrorRoutes(app)

	// Timeout + Rate Limiter middlewares
	basicApi := app.Party("/", extra_middlewares.Timeout(time.Second*time.Duration(g.CFG.Timeout)), extra_middlewares.ConcurrentLimiter(g.CFG.MaxConcurrentRequests))

	// More basic
	staticFiles := app.Party("/")

	basicApi.Get("/", handlers.Hello)

	{ //* serving static files
		statics := staticFiles.Party("/")
		statics.HandleDir(g.MediaServePath, iris.Dir(g.Media.GetAddress()))
	}

	{ //* /api
		api := basicApi.Party("/api")
		api.Get("/", handlers.Hello)
	}
}
