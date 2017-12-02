package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	word := getWord()
	synonyms := getSynos(word)
	for idx, synonym := range synonyms {
		fmt.Printf("[%03d] %s ", idx, synonym)
	}
}

func getWord() string {
	if len(os.Args) == 1 {
		log.Fatalln("Please provide a French word as argument!")
	}
	return os.Args[1]
}

func getSynos(word string) []string {
	var synonyms []string
	url := fmt.Sprintf("http://www.synonymo.fr/synonyme/%s", word)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	ulSelection := doc.Find("ul")
	current := ulSelection.Eq(0)
	liSelection := current.Find("li")
	for idx := range liSelection.Nodes {
		word := liSelection.Eq(idx)
		synonyms = append(synonyms, strings.TrimSpace(word.Text()))
	}
	return synonyms
}
