package main

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("ao3api-rod: ")
	log.SetOutput(os.Stderr)

	// add additional code here
}
