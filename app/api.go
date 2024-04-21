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
	routes.HTTP(app)

	runCronJobs()

	// RunClonesAndServer(app)
	app.Listen(g.CFG.Gateway.IP + ":" + g.CFG.Gateway.Port)
}

func JustReturnAPI() *iris.Application {
	// Print Info
	info()

	app := iris.New()
	app.Configure(iris.WithoutStartupLog)

	// Router Settings
	g.App = app
	routes.HTTP(app)

	runCronJobs()

	// return app
	return app
}
