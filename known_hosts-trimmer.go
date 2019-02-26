package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

type Configuration struct {
	File   string
	Lineno int
}

type Line struct {
	Content string
	Number  int
}

const (
	VERSION = "0.0.0"
)

var (
	c            Configuration
	debug, quiet bool
)

func init() {
	c.File = filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
}

func main() {
	var err error
	flag.BoolVar(&debug, "debug", false, "debugging mode")
	flag.BoolVar(&quiet, "q", false, "be quiet")
	flag.StringVar(&c.File, "f", c.File, "`file` to remove lines from")
	flag.Usage = usage
	flag.Parse()
	if _, err = os.Stat(c.File); os.IsNotExist(err) {
		fmt.Println("File not found!")
		return

	}
	if len(flag.Args()) == 0 {
		usage()
	}
	c.Lineno, err = strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println("Invalid parameter:", c.Lineno)
		usage()
	}
	if c.Lineno == 0 {
		fmt.Println("lineno cannot be zero!")
		return
	}
	if debug {
		fmt.Printf("DEBUG: %#v\n", c)
	}
	removeplus()
}

func removeplus() {
	var lines []Line
	fh, err := os.Open(c.File)
	if err != nil {
		panic(err)
	}
	idx := 1
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		var line Line
		line.Content = scanner.Text()
		line.Number = idx
		lines = append(lines, line)
		idx++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fh.Close()
	max := len(lines)
	if debug {
		fmt.Printf("DEBUG: %d line%s found in %s\n", max, tools.Ternary(max > 1, "s",  ""),
		c.File)
	}
	if c.Lineno > max {
		fmt.Printf("File %q has only %d line%s!\n", c.File, max,
			tools.Ternary(max > 1, "s", ""))
		return
	}
	var buffer strings.Builder
	for _, line := range lines {
		if line.Number == c.Lineno || line.Number == c.Lineno+1 {
			if !quiet {
				fmt.Println("Removing line", line)
			}
			continue
		}
		buffer.WriteString(line.Content)
		buffer.WriteString("\n")
	}
	ioutil.WriteFile(c.File, []byte(buffer.String()), 0600)
}

func usage() {
	fmt.Println("known_hosts-trimmer, v.", VERSION, "(c) Lyderic Landry, London 2019")
	fmt.Println("Usage: known_hosts-trimmer <lineno>")
	fmt.Println("removes two consecutive lines from ssh known_host file starting at <lineno>")
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("For example, to remove line 95 and 96 from ~/.ssh/known_hosts:")
	fmt.Println("known_hosts-trimmer 95")
	os.Exit(23)
}

func (line Line) String() string {
	return fmt.Sprintf("[%02d] %s", line.Number, line.Content)
}
