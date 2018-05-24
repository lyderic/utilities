package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ch <word>")
		return
	}
	word := os.Args[1]
	grepCmd := exec.Command("grep", "-siIrn", "--color", word)
  lessCmd := exec.Command("less")
	lessCmd.Stdin = grepCmd.Stdout
  lessCmd.Stdout = os.Stdout
	grepCmd.Stderr = os.Stderr
	lessCmd.Stderr = os.Stderr
	err := grepCmd.Run()
	if err != nil {
		fmt.Println("not found")
	}

}
