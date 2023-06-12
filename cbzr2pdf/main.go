package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/lyderic/tools"
)

func main() {
	DEBUG = flag.Bool("d", false, "debug mode")
	SHOWCMD = flag.Bool("showcmd", false, "show output of unix commands")
	DRYRUN = flag.Bool("n", false, "dry run, don't execute")
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
	}
	var curdir string
	var err error
	if curdir, err = os.Getwd(); err != nil {
		E(err)
	}
	debug("Current directory: %q\n", curdir)
	for _, input := range flag.Args() {
		BlueB("Processing %q\n", input)
		c := initConfig(input) // configuration.go
		debug("Configuration:\n%s\n", c)
		if c.Ext != ".cbz" && c.Ext != ".cbr" {
			Yellow("%q: not a valid file!\n", c.Path)
			continue
		}
		extractcbzr(c) // unix.go
		images2pdf(c)  // unix.go
		createpdf(c)   // unix.go
	}
}

func usage() {
	progname := os.Args[0]
	fmt.Printf(`%s v%s (c) Lydéric Landry, Luxembourg 2023
Usage: %s <option> <file> <file>...
convert .cbz or .cbr comics files to pdf
Options:
`, progname, VERSION, progname)
	flag.PrintDefaults()
	os.Exit(23)
}
