package main

import (
	"os"

	"github.com/gsy/gddd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
