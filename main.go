package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"service/app"
	g "service/global"
	"service/pkg/colors"
	"service/pkg/repositories"
	"service/pkg/translator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of " + g.Name + ":")
		fmt.Println("\nCommands:")
		fmt.Println("  translate  translates whole yml files")
		fmt.Println("  migrate    migrates to the latest changes")
		fmt.Println("  demigrate  demigrates one migration back")
		fmt.Println()
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
			app.MigrateLatestChanges()
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sMigration Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "demigrate":
			app.DemigrateOneChange()
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sDemigration Finished Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		case "sqlc":
			repositories.GenerateRepositories(filepath.Join(g.CFG.PWD, "repositories"))
			fmt.Println(colors.Cyan, fmt.Sprintf("\n==%sRepositories Generated Successfully%s==%s\n", colors.Green, colors.Cyan, colors.Reset))
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}
