package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var base string

func init() {
	switch runtime.GOOS {
	case "windows":
		base = "C:\\"
	default:
		base = "/"
	}
}

func main() {
	fmt.Println("Starting from:", base)
	if len(os.Args) == 1 {
		fmt.Println("Usage: findend <ending>")
		fmt.Println("find all files on / (Unix) or C: (Windows) which name ends in 'ending'")
		return
	}
	ending := os.Args[1]
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Warning with path %q: %v\n", path, err)
		} else {
			if strings.HasSuffix(info.Name(), ending) {
				fmt.Println(path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
