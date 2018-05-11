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
	cmd := exec.Command("grep", "-siIrn", "--color", word)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("not found")
	}

}
