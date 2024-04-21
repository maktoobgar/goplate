package main

import (
	"flag"
	"fmt"
	"os"
	"service/app"
)

func main() {
	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()
		fmt.Println()
		fmt.Println("Commands:")
		// fmt.Println("  command\n        description")
	}

	if len(os.Args) < 2 {
		app.API()
		return
	} else {
		switch os.Args[1] {
		default:
			fmt.Printf("expected 'fake' but '%s' received\n", os.Args[1])
			flag.Usage()
			fmt.Println()
		}
	}
}
