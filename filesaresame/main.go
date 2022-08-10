package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	file1 := os.Args[1]
	file2 := os.Args[2]
	os.Exit(process(file1, file2))
}

func process(file1, file2 string) (rcode int) {
	var err error
	for _, file := range []string{file1, file2} {
		if _, err = os.Stat(file); os.IsNotExist(err) {
			panic(err)
		}
	}
	content1, err := os.ReadFile(file1)
	if err != nil {
		panic(err)
	}
	content2, err := os.ReadFile(file2)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(content1, content2) {
		return 1
	}
	return
}

func usage() {
	fmt.Printf("%s <file1> <file2>\n", filepath.Base(os.Args[0]))
	fmt.Println("compare two files. Return 0 if files are identical, otherwise return 1")
	os.Exit(23)
}
