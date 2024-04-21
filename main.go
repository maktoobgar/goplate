package main

import (
	"flag"
	"fmt"
	"os"
	"service/app"
	g "service/global"
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
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}
