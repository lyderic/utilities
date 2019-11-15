/*
  This utility display the pdf passed as argument
	Needs 'evince' binary
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lyderic/tools"
)

const (
	pdfreader = "evince"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	path := sanitize(os.Args[1])
	extension := filepath.Ext(path)
	switch extension {
	case ".pdf", ".PDF":
		display(path)
	default:
		fmt.Printf("%s: this doesn't look to be a PDF!\n", extension)
	}
}

func sanitize(file string) (path string) {
	var err error
	path, err = filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		tools.PrintRedf("%s: file not found\n", path)
		os.Exit(23)
	}
	return
}

func display(path string) {
	err := tools.CheckBinaries(pdfreader)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(pdfreader, path)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Println("Usage: di <file.pdf>")
}
