package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Synonym struct {
	Index int
	Word  string
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	word := getWord()
	synonyms := getSynos(word)
	for _, synonym := range synonyms {
		fmt.Printf("[%03d] %s ", synonym.Index, synonym.Word)
	}
	fmt.Println()
}

func getWord() string {
	if len(os.Args) == 1 {
		log.Fatalln("Please provide a French word as argument!")
	}
	return os.Args[1]
}

func getSynos(word string) (synonyms []Synonym) {
	url := fmt.Sprintf("https://www.cnrtl.fr/synonymie/%s", word)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".syno_format").Each(func(i int, s *goquery.Selection) {
		var synonym Synonym
		synonym.Index = i + 1
		synonym.Word = s.Text()
		synonyms = append(synonyms, synonym)
	})
	return
}
