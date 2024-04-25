package main

import (
	"flag"
	"fmt"
	"os"
	"service/app"
	g "service/global"
	"service/pkg/translator"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage of " + g.Name + ":")
		fmt.Println("\nCommands:")
		fmt.Println("  command  description")
	}

	if len(os.Args) < 2 {
		app.API()
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "translate":
			translator.GenerateCode(g.CFG.Language.Path, g.CFG.Language.DefaultLanguage)
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}
