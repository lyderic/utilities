package main

import (
	"fmt"
	"github.com/lyderic/tools"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"log"
	"os"
	"path"
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
	if len(os.Args) == 1 {
		fmt.Println("Usage:", os.Args[0], "<file(s)>")
		os.Exit(42)
	}
	var all []Data
	for _, file := range os.Args[1:] {
		if !tools.PathExists(file) {
			skip(file, "not found!")
			continue
		}
		if info, err := os.Stat(file); err == nil && info.IsDir() {
			skip(file, "is a directory")
			continue
		}
		data := Data{}
		data.Name = path.Base(file)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		if !utf8.Valid(b) {
			skip(file, "not UTF8 readable")
			continue
		}
		content := string(b)
		data.Bytes = len(b)
		data.Chars = utf8.RuneCount(b)
		data.Words = len(strings.Fields(content))
		all = append(all, data)
	}
	if len(all) == 0 {
		fmt.Println("No readable file found")
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
