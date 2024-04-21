package pkg

import (
	"log"
	"os"
)

func ReCreateDir(path string) {
	// Directory does exist or not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If not, create it
		if err = os.MkdirAll(path, 0777); err != nil {
			log.Panicf("Error creating directory: %s\n", err)
		}
	} else {
		// If yes, remove it and create it again
		os.RemoveAll(path)
		if err = os.MkdirAll(path, 0777); err != nil {
			log.Panicf("Error creating directory: %s\n", err)
		}
	}
}
