package main

import (
	"log"
	"os"

	"github.com/ryota-sakamoto/topsy-turvy/pkg/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
