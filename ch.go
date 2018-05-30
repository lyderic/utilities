package main

import (
	"fmt"
	"github.com/lyderic/tools"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	if os.Args[1][0] == '-' {
		usage()
		return
	}
	word := os.Args[1]
	cmd := exec.Command("grep", "-siIrn", "--color=always", word)
	b, err := cmd.CombinedOutput()
	if err != nil {
		tools.PrintColorf(tools.RED, "'%s' not found\n", word)
		return
	}
	output := strings.TrimSpace(string(b))
	lines := strings.Split(output, "\n")
	sort.Strings(lines)
	tools.Less(strings.Join(lines, "\n"))
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "<word>")
	fmt.Println("look for <word> in all files in dir")
}
