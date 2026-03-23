package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: rss-publisher [output-file]")
		os.Exit(1)
	}

	// TODO
}
