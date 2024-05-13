package routes

import (
	g "service/global"
	"service/handlers"
	"service/handlers/error_handlers"
	"service/middlewares/extra_middlewares"
	"service/pkg/api"
	admin_users_routes "service/routes/admin"
	"strings"
	"time"

	"github.com/kataras/iris/v12/context"

	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Applies all necessary middlewares
func addBasicMiddlewares(app *iris.Application) {
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

	api.Get(basicApi, "/", []context.Handler{handlers.Hello}, &handlers.HelloRes{}, api.Setting{Summary: "Simple hello world"})

	// Activate Swagger
	basicApi.Get("/swagger/{any}", func(ctx *context.Context) {
		httpSwagger.Handler(
			httpSwagger.URL("doc.json"), //The url pointing to API definition
		)(ctx.ResponseWriter(), ctx.Request())
	})

	{ //* serving static files
		statics := staticFiles.Party("/")
		statics.HandleDir(g.MediaServePath, iris.Dir(g.Media.GetAddress()))
	}

	AuthHTTP(basicApi)
	UsersHTTP(basicApi)
	admin_users_routes.AdminHTTP(basicApi)
}
