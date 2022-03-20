package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/lyderic/tools"

	"github.com/PuerkitoBio/goquery"
	"github.com/jedib0t/go-pretty/v6/table"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Please provide a French word as argument!")
	}
	synonyms := getSynonyms(os.Args[1:])
	ln := len(synonyms)
	fmt.Printf("found %d synonym%s:\n", ln, Ternary(ln > 1, "s", ""))
	/*
		for idx, synonym := range synonyms {
			Green("[%02d:%s]", idx, synonym)
		}
		fmt.Println()
	*/
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.SetStyle(table.StyleLight)
	t.SetStyle(NoStyle)
	for i := 0; i < len(synonyms); i++ {
		one := synonyms[i]
		two := ""
		if i < ln-1 {
			two = synonyms[i+1]
		}
		t.AppendRow(table.Row{one, two})
		i++
	}
	t.Render()
}

func getSynonyms(args []string) (synonyms []string) {
	word := strings.Join(args, "%20")
	url := fmt.Sprintf("https://www.cnrtl.fr/synonymie/%s", word)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".syno_format").Each(func(i int, s *goquery.Selection) {
		synonyms = append(synonyms, s.Text())
	})
	return
}

var NoStyle = table.Style{
	Name:    "NoStyle",
	Box:     table.StyleBoxDefault,
	Color:   table.ColorOptionsDefault,
	Format:  table.FormatOptionsDefault,
	HTML:    table.DefaultHTMLOptions,
	Options: table.OptionsNoBordersAndSeparators,
	Title:   table.TitleOptionsDefault,
}
