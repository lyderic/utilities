package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lyderic/tools"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const VERSION = "0.0.0"
const BAK = "bak"

var dbg bool
var known_hosts string

type Input struct {
	File    string
	Bak     string
	Numbers []string
}

func init() {
	known_hosts = filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")
}

func main() {
	var file, bak string
	var input Input
	var err error
	flag.BoolVar(&dbg, "debug", false, "debug mode")
	flag.StringVar(&file, "f", known_hosts, "file to remove line(s) from")
	flag.StringVar(&bak, "b", BAK, "extension of backup file")
	flag.Usage = usage
	flag.Parse()
	if dbg {
		tools.PrintRedln("[DEBUG MODE ON]")
	}
	var abspath string
	if abspath, err = filepath.Abs(file); err != nil {
		log.Fatal(err)
	}
	input.File = abspath
	input.Bak = bak
	for _, arg := range flag.Args() {
		input.Numbers = append(input.Numbers, arg)
	}
	debug("input: %#v\n", input)
	if err = process(input); err != nil {
		tools.PrintRedln("*** FAILED ***")
		fmt.Println(err)
	}
}

func process(input Input) error {
	var err error
	var numbers []int
	if len(input.Numbers) == 0 {
		return fmt.Errorf("Please specify which line(s) to remove")
	}
	if err = backup(input); err != nil {
		return err
	}
	var lines []string
	if lines, err = linesFromFile(input.File); err != nil {
		return err
	}
	ln := len(lines)
	fmt.Printf("%s: %d line%s read\n", input.File, ln, tools.Ternary(ln > 1, "s", ""))
	for _, number := range input.Numbers {
		var n int
		if n, err = strconv.Atoi(number); err != nil {
			return fmt.Errorf("%s: invalid line number", number)
		}
		if n < 1 {
			return fmt.Errorf("cannot remove line %d! line count start at 1", n)
		}
		if n > ln {
			return fmt.Errorf("cannot remove line %d! file has only %d line%s",
				n, ln, tools.Ternary(ln > 1, "s", ""))
		}
		numbers = append(numbers, n)
	}
	debug("numbers: %#v\n", numbers)
	return nil
}

func backup(input Input) error {
	fmt.Println("backing up", input.File)
	var err error
	parent := filepath.Dir(input.File)
	debug("Parent directory: %s\n", parent)
	backupFile := fmt.Sprintf("%s.%s", input.File, input.Bak)
	debug("Backup file: %s\n", backupFile)
	backupPath := filepath.Join(parent, backupFile)
	debug("Backup path: %s\n", backupPath)
	if err = tools.Copy(input.File, backupPath); err != nil {
		return err
	}
	return nil
}

func linesFromFile(file string) ([]string, error) {
	var lines []string
	var err error
	var f *os.File
	if f, err = os.Open(file); err != nil {
		return lines, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func debug(format string, args ...interface{}) {
	format = "[DEBUG] " + format
	if dbg {
		fmt.Printf(format, args...)
	}
}

func usage() {
	appname := filepath.Base(os.Args[0])
	fmt.Printf("%s - v.%s (c) Lyderic Landry, London 2018\n", appname, VERSION)
	fmt.Printf("Usage: %s [--debug] [-f <file>] [-b <ext>] x y z...\n", appname)
	fmt.Println("Remove line numbers x, y and z from <file>")
	flag.PrintDefaults()
}
