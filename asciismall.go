package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/lyderic/tools"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	dryrun bool
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	flag.BoolVar(&dryrun, "n", false, "dry-run: simulate action only")
	flag.Usage = usage
	flag.Parse()
	inputs := flag.Args()
	if len(inputs) == 0 {
		usage()
		os.Exit(3)
	}
	if dryrun {
		tools.PrintYellowln("DRY RUN! No change, simulating only!")
	}
	for _, input := range inputs {
		finfo, err := os.Stat(input)
		if err != nil {
			tools.PrintRedf("%q: invalid input. Skipping...\n", input)
			continue
		}
		if finfo.IsDir() {
			err = processDir(input)
		} else {
			err = processPath(input)
		}
		if err != nil {
			tools.PrintRedf("Error processing input %q: %v\n", input, err)
			continue
		}
	}
}

func processDir(dirpath string) (err error) {
	err = filepath.Walk(dirpath, func(path string, finfo os.FileInfo, err error) error {
		if finfo.IsDir() {
			fmt.Printf("%q is a directory: conversion skipped...\n", path)
			return nil
		}
		return processPath(path)
	})
	return
}

func processPath(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%q: file not found", path)
	}
	name := filepath.Base(path)
	parent := filepath.Dir(path)
	newname := strings.ToLower(toAscii(name))
	newpath := filepath.Join(parent, newname)
	if newname == name {
		fmt.Printf("%q: no change\n", name)
		return
	}
	if _, err = os.Stat(newpath); os.IsNotExist(err) {
		if dryrun {
			tools.PrintYellowf("[DRY-RUN] %s -> %s\n", path, newpath)
			return nil
		} else {
			fmt.Println(path, "->", newpath)
			return os.Rename(path, newpath)
		}
	} else {
		fmt.Printf("%q: file already exists. Not overwritten!\n", name)
	}
	return
}

func toAscii(word string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, word)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func usage() {
	fmt.Println("Usage: [-r] [-n] asciismall <files...> <directories...>")
	flag.PrintDefaults()
	fmt.Println("This program converts file names to lower case and no non-ascii characters")
	fmt.Println("Directories are processed recursively. Directory names are NOT converted")
}
