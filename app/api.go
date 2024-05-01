package app

import (
	g "service/global"
	"service/routes"

	"github.com/kataras/iris/v12"
)

func API() {
	// Initialize all
	InitializeService()
	// Print Info
	Info()

	app := iris.New()
	app.Configure(iris.WithoutStartupLog)

	// Router Settings
	g.App = app
	routes.HTTP(app)

	runCronJobs()

	app.Listen(g.CFG.Gateway.IP + ":" + g.CFG.Gateway.Port)
}
