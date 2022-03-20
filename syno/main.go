package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	. "github.com/lyderic/tools"

	"github.com/PuerkitoBio/goquery"
	"github.com/jedib0t/go-pretty/v6/table"
)

// flags
var (
	debug bool
	nocol bool
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	flag.BoolVar(&debug, "d", false, "show debugging information")
	flag.BoolVar(&nocol, "n", false, "display synonyms without using columns")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
		return
	}
	synonyms := getSynonyms(flag.Args())
	ln := len(synonyms)
	fmt.Printf("found %d synonym%s:\n", ln, Ternary(ln > 1, "s", ""))
	if debug {
		for idx, synonym := range synonyms {
			Cyan("[%02d:%s]", idx, synonym)
		}
		fmt.Println()
	}
	if nocol {
		for _, synonym := range synonyms {
			fmt.Println(synonym)
		}
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
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

func usage() {
	fmt.Println("syno [-d] [-n] <word(s)>")
	flag.PrintDefaults()
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
