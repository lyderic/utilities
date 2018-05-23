package main

import (
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

func main() {
	var pattern string
	dir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-h":
			usage()
			return
		default:
			pattern = os.Args[1]
		}
	}
	var all []Data
	listing, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range listing {
		name := item.Name()
		abspath := filepath.Join(dir, name)
		if !strings.Contains(name, pattern) {
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
	lines := [][]string{}
	for _, data := range all {
		n := data.Name
		c := tools.ThousandSeparator(data.Chars)
		w := tools.ThousandSeparator(data.Words)
		b := tools.ThousandSeparator(data.Bytes)
		lines = append(lines, []string{n, c, w, b})
	}
	table := tablewriter.NewWriter(os.Stdout)
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
	fmt.Printf("> skipped %s: %s\n", file, reason)
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "<pattern>")
	fmt.Println("count chars, words and bytes of files in current directory")
	fmt.Println("optionnally filtered by pattern")
}
