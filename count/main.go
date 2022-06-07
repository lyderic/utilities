package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/lyderic/tools"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type Data struct {
	Name  string
	Bytes int
	Chars int
	Words int
}

/* Globals switches */
var verbose, paginate, hidden bool

func main() {
	var pattern string
	flag.BoolVar(&verbose, "v", false, "be verbose")
	dirPtr := flag.String("d", ".", "directory")
	flag.BoolVar(&hidden, "h", false, "count hidden files")
	flag.BoolVar(&paginate, "l", false, "paginate output")
	flag.Parse()
	dir, err := filepath.Abs(*dirPtr)
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		fmt.Println("Directory:", dir)
	}
	if len(flag.Args()) > 0 {
		pattern = flag.Args()[0]
	}
	var all []Data
	listing, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range listing {
		name := item.Name()
		if name[0] == '.' {
			if !hidden {
				skip(name, "hidden file")
				continue
			}
		}
		abspath := filepath.Join(dir, name)
		if !strings.Contains(name, pattern) {
			skip(name, "pattern not matched")
			continue
		}
		if info, err := os.Stat(abspath); err == nil && info.IsDir() {
			skip(name, "is a directory")
			continue
		}
		data := Data{}
		data.Name = name
		b, err := ioutil.ReadFile(abspath)
		if err != nil {
			log.Fatal(err)
		}
		if !utf8.Valid(b) {
			skip(name, "not UTF8 readable")
			continue
		}
		content := string(b)
		data.Bytes = len(b)
		data.Chars = utf8.RuneCount(b)
		data.Words = len(strings.Fields(content))
		all = append(all, data)
	}
	if len(all) == 0 {
		fmt.Printf("file not found, not readable or not matching pattern (%s)\n", pattern)
		return
	}
	display(all)
}

func display(all []Data) {
	buffer := new(bytes.Buffer)
	lines := [][]string{}
	for _, data := range all {
		n := data.Name
		c := tools.ThousandSeparator(data.Chars)
		w := tools.ThousandSeparator(data.Words)
		b := tools.ThousandSeparator(data.Bytes)
		lines = append(lines, []string{n, c, w, b})
	}
	table := tablewriter.NewWriter(buffer)
	table.SetHeader([]string{"Name", "Chars", "Words", "Bytes"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT,
		tablewriter.ALIGN_RIGHT})
	table.AppendBulk(lines)
	totals := getTotals(all)
	table.SetFooter([]string{
		totals.Name,
		tools.ThousandSeparator(totals.Chars),
		tools.ThousandSeparator(totals.Words),
		tools.ThousandSeparator(totals.Bytes)})
	table.Render()
	if paginate {
		tools.Less(buffer.String())
	} else {
		fmt.Print(buffer.String())
	}
}

func getTotals(all []Data) Data {
	totals := Data{}
	var c, w, b int
	for _, data := range all {
		c = c + data.Chars
		w = w + data.Words
		b = b + data.Bytes
	}
	totals.Name = "TOTALS"
	totals.Chars = c
	totals.Words = w
	totals.Bytes = b
	return totals
}

func skip(file, reason string) {
	if verbose {
		fmt.Printf("> skipped %s: %s\n", file, reason)
	}
}
