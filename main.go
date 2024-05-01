package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"service/app"
	g "service/global"
	"service/pkg/colors"
	"service/pkg/generator"
	"service/pkg/repositories"
	"service/pkg/translator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of " + g.Name + ":")
		fmt.Println("Commands:")
		fmt.Println("  translate  translates whole yml files")
		fmt.Println("  migrate    migrates to the latest changes")
		fmt.Println("  demigrate  demigrates one migration back")
		fmt.Println("  sqlc       generates content of repositories folder")
		fmt.Println("  new        generates new files and content for fast development")
		fmt.Println("  -v         shows info about the service")
	}

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")

	handlerCmd := flag.NewFlagSet("handler", flag.ExitOnError)
	var functionName string
	handlerCmd.StringVar(&functionName, "name", "", "receives handler name in pascal or snake case")
	var packageName string
	handlerCmd.StringVar(&packageName, "package", "", "receives package name in pascal or snake case")

	flag.Parse()
	if showVersion {
		app.Info()
		return
	}

	if len(os.Args) < 2 {
		app.API()
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "translate":
			translator.GenerateCode(g.CFG.Language.Path, g.CFG.Language.DefaultLanguage)
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sTranslation Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "migrate":
			app.InitializeService()
			app.MigrateLatestChanges()
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sMigration Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "demigrate":
			app.InitializeService()
			app.DemigrateOneChange()
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sDemigration Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "sqlc":
			repositories.GenerateRepositories(filepath.Join(g.CFG.PWD, "repositories"))
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sRepositories Generated Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "handler":
			err := handlerCmd.Parse(os.Args[2:])
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(functionName) > 0 {
				generator.GenerateNewHandler(functionName, packageName, g.CFG.PWD)
				fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sGeneration Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
			} else {
				fmt.Println("flag required: -name")
				handlerCmd.Usage()
			}
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}
