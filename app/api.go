package app

import (
	"path/filepath"
	g "service/global"
	"service/pkg/api"
	"service/routes"

	"github.com/kataras/iris/v12"
)

func API(justReturn ...bool) *iris.Application {
	run := len(justReturn) == 0 || !justReturn[0]
	// Initialize all
	InitializeService()
	// Print Info
	if run {
		Info()
	}

	app := iris.New()
	app.Configure(iris.WithoutStartupLog)

	// Router Settings
	g.App = app
	routes.HTTP(app)

	runCronJobs()

	return api.Run(app, g.CFG.IP, g.CFG.Port, filepath.Join(g.CFG.PWD, "docs"))
}
