package app

import (
	g "service/global"
	"service/routes"

	"github.com/kataras/iris/v12"
)

func API() {
	// Print Info
	info()

	app := iris.New()
	app.Configure(iris.WithoutStartupLog)

	// Router Settings
	g.App = app
	apiRouter := routes.HTTP(app)
	routes.AuthHTTP(apiRouter)
	routes.UsersHTTP(apiRouter)

	runCronJobs()

	app.Listen(g.CFG.Gateway.IP + ":" + g.CFG.Gateway.Port)
}
