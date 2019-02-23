package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const (
	FOUND    = 0
	NOTFOUND = 1
)

var (
	debug   bool
	verbose bool
)

func main() {
	var retcode int
	flag.BoolVar(&debug, "debug", false, "show output of rclone command")
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.Parse()
	cmd := exec.Command("rclone", "--ask-password=false", "listremotes")
	if debug {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		display("Rclone password not in memory!")
		retcode = NOTFOUND
	} else {
		display("Rclone password found.")
		retcode = FOUND
	}
	os.Exit(retcode)
}

func display(message string) {
	if verbose {
		fmt.Println(message)
	}
}
