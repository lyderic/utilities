package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lyderic/tools"
)

type Config struct {
	Root  string
	From  int
	Yes   bool
	Width int
	Debug bool
}

var c Config

type File struct {
	Finfo     os.FileInfo
	Extension string
	Basename  string
	Number    int
	Content   []byte
}

func main() {
	from := flag.Int("f", 0, "shift from this `number`")
	root := flag.String("r", "chapitre", "`root` of file name")
	yes := flag.Bool("y", false, "don't ask for confirmation")
	width := flag.Int("w", 2, "number format")
	dbg := flag.Bool("debug", false, "print debugging information")
	flag.Usage = usage
	flag.Parse()
	if len(os.Args) == 1 {
		usage()
		return
	}
	c = Config{*root, *from, *yes, *width, *dbg}
	for _, file := range fetch() {
		if file.Number >= c.From {
			process(file)
		}
	}
}

func fetch() (files []File) {
	listing, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, finfo := range listing {
		filename := finfo.Name()
		if finfo.IsDir() {
			debug("skipping %q as it is a directory\n", filename)
			continue
		}
		if strings.HasPrefix(filename, c.Root) {
			var f File
			f.Finfo = finfo
			f.Extension = filepath.Ext(filename)
			f.Basename = filename[0 : len(filename)-len(f.Extension)]
			f.Number, err = strconv.Atoi(strings.Replace(f.Basename, c.Root, "", 1))
			if err != nil {
				alert("Cannot parse name of %q. Skipping\n", filename)
				continue
			}
			f.Content, err = ioutil.ReadFile(filename)
			if err != nil {
				alert("Cannot read content of %q. Skipping\n", filename)
				continue
			}
			files = append(files, f)
		} else {
			debug("skipping %q as it doesn't have root %q\n", finfo.Name(), c.Root)
			continue
		}
	}
	return
}

func process(f File) {
	newname := fmt.Sprintf("%s%02d%s", c.Root, f.Number+1, f.Extension)
	fmt.Println(f.Finfo.Name(), ">", newname)
	ioutil.WriteFile(newname, f.Content, f.Finfo.Mode())
}

func alert(format string, message ...interface{}) {
	tools.PrintRedf(format, message...)
}

func debug(format string, message ...interface{}) {
	if c.Debug {
		tools.PrintYellowf(format, message...)
	}
}
func usage() {
	fmt.Printf(`Usage: %s <options>
Increment numbers at the end of <root> by <increment>
Example: if there are two files called 'chapitre01.lkl' and 'chapitre02.lkl',
         running '%s -f 2 -i 1 -r chapitre' will rename them resp. to:
         chapitre02.lkl and chapitre03.lkl
Options:
`, os.Args[0], os.Args[0])
	flag.PrintDefaults()
}
