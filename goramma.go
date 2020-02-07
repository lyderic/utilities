package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	flag.BoolVar(&redon, "r", false, "détecte les répétitions")
	flag.Usage = usage
	flag.Parse()
	if len(os.Args) == 1 {
		usage()
		return
	}
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
	max := getMaxWidth()
	var buffer strings.Builder
	buffer.WriteString(strings.Repeat("-", termwidth) + "\n")
	buffer.WriteString(fmt.Sprintf("%s [%d|%s]\n", file, termwidth, max))
	buffer.WriteString(strings.Repeat("-", termwidth) + "\n")
	var output []byte
	cmd := exec.Command("grammalecte", "-f", file,
		"--textformatter", "--only_when_errors", "--width", max,
		"-on", "mapos", "mc", "neg", "poncfin", "tab", "idrule", "liga",
		tools.Ternary(redon, "redon1", "").(string),
		tools.Ternary(redon, "redon2", "").(string),
		"-pdi", dico)
	output, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	buffer.Write(output)
	report = buffer.String()
	return
}

func getMaxWidth() (max string) {
	if termwidth < 40 {
		return "40"
	}
	if termwidth > 200 {
		return "200"
	}
	allowed := []int{40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}
	var index, value int
	for index, value = range allowed {
		if termwidth < value {
			break
		}
	}
	return strconv.Itoa(allowed[index-1])
}

func usage() {
	fmt.Println("Usage: [-r] goramma <file> <file> ...")
	flag.PrintDefaults()
}
