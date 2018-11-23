package main

import (
	"log"
	"os"
	"os/exec"
)

const (
	spellfile = "~/.vim/spell/fr.utf-8.add"
)

func main() {
	cmd := exec.Command("vim", "--cmd", ":mkspell! "+spellfile, "--cmd", ":q!")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
