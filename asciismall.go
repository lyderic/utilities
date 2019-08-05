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

type File struct {
	Name    string
	Path    string
	Parent  string
	Newname string
	Newpath string
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	flag.BoolVar(&dryrun, "n", false, "dry-run: simulate action only")
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()

	if len(paths) == 0 {
		usage()
		os.Exit(3)
	}

	if dryrun {
		tools.PrintYellowln("DRY RUN! No change, simulating only!")
	}

	for _, path := range paths {
		var file File
		file.Path = path
		file.Name = filepath.Base(file.Path)
		file.Parent = filepath.Dir(file.Path)
		file.Newname = strings.ToLower(toAscii(file.Name))
		file.Newpath = filepath.Join(file.Parent, file.Newname)
		err := process(file)
		if err != nil {
			tools.PrintRedf("[FAIL] %v\n", err)
		}
	}
}

func process(file File) (err error) {
	if _, err = os.Stat(file.Path); os.IsNotExist(err) {
		return fmt.Errorf("%q: file not found", file.Path)
	}
	if file.Newname == file.Name {
		fmt.Printf("%q: no change\n", file.Name)
		return
	}
	if _, err = os.Stat(file.Newpath); os.IsNotExist(err) {
		if dryrun {
			tools.PrintYellowf("[DRY-RUN] %s -> %s\n", file.Path, file.Newpath)
			return nil
		} else {
			fmt.Println(file.Path, "->", file.Newpath)
			return os.Rename(file.Path, file.Newpath)
		}
	} else {
		fmt.Printf("%q: file already exists. Not overwritten!\n", file.Name)
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
	fmt.Println("Usage: [-n] asciismall <files...>")
	flag.PrintDefaults()
	fmt.Println("This program converts file names to lower case and no non-ascii characters")
}
