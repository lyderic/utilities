package main

import (
	"fmt"
	"os"

	"github.com/lyderic/tools"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	paths := os.Args[1:]
	verbose := true
	for _, path := range paths {
		tools.Sanitize(path, verbose)
	}

}

func usage() {
	fmt.Println("Usage: sanitize <files>")
	fmt.Println("set correct French typography for <files>")
}
