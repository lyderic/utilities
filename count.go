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
		data := Data{}
		data.Name = path.Base(file)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		content := string(b)
		data.Bytes = len(b)
		data.Chars = utf8.RuneCountInString(content)
		data.Words = len(strings.Fields(content))
		all = append(all, data)
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
