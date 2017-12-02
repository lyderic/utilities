package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		lcname := strings.ToLower(toAscii(name))
		if lcname == name {
			fmt.Println(name, ": no change")
			continue
		}
		fmt.Println(name, "->", lcname)
		err = os.Rename(name, lcname)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func toAscii(word string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, word)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
