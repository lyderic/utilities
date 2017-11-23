package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	doc, err := goquery.NewDocument("http://www.synonymo.fr/synonyme/pouvoir")
	if err != nil {
		log.Fatal(err)
	}

	//findLiWord(doc)

	ulSelection := doc.Find("ul")
  current := ulSelection.Eq(0)
  liSelection := current.Find("li")
  for idx := range liSelection.Nodes {
    word := liSelection.Eq(idx)
    fmt.Printf("%s\n", strings.TrimSpace(word.Text()))
  }
}

func findLiWord(doc *goquery.Document) {
	doc.Find("li .word").Each(func(index int, item *goquery.Selection) {
		text := item.Text()
		fmt.Printf("%03d: %s\n", index, text)
	})

}
