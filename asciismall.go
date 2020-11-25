package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if len(os.Args) > 1 {
		for _, args := range os.Args[1:] {
			process(args)
		}
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text())
	}
}

func process(input string) {
	fmt.Println(strings.ToLower(toAscii(input)))
}

func toAscii(word string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, word)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
