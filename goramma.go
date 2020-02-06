package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lyderic/tools"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	termwidth int
	dico      string
	redon     bool
)

func init() {
	termwidth, _, _ = terminal.GetSize(0)
	dico = filepath.Join(os.Getenv("HOME"), "cyranoplain", "fr.__personal__.json")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	flag.BoolVar(&redon, "r", false, "détecte les répétitions")
	flag.Usage = usage
	flag.Parse()
	files := flag.Args()
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			tools.PrintYellowf("%q: file not found. Skipping.\n", file)
			continue
		}
		report, err := analyze(file)
		if err != nil {
			if err != nil {
				panic(err)
			}
		}
		tools.Less(report)
	}
}

func analyze(file string) (report string, err error) {
	var buffer strings.Builder
	buffer.WriteString(strings.Repeat("-", termwidth) + "\n")
	buffer.WriteString(file + "\n")
	buffer.WriteString(strings.Repeat("-", termwidth) + "\n")
	var output []byte
	output, err = exec.Command("grammalecte", "-f", file,
		"--textformatter", "--only_when_errors",
		"-on", "mapos", "mc", "neg", "poncfin", "tab",
		tools.Ternary(redon, "redon1", "").(string),
		tools.Ternary(redon, "redon2", "").(string),
		"-pdi", dico).CombinedOutput()
	buffer.Write(output)
	report = buffer.String()
	return
}

func usage() {
	fmt.Println("Usage: goramma <file> <file> ...")
	flag.PrintDefaults()
}
