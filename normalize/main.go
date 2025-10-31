package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	_, err := io.Copy(&buf, in)
	if err != nil {
		os.Exit(1)
	}
	out := removeDiacritics(buf.String())
	_, _ = os.Stdout.WriteString(out)
}

func removeDiacritics(src string) string {
	decomposed := norm.NFD.String(src)
	var b bytes.Buffer
	for _, r := range decomposed {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return norm.NFC.String(b.String())
}
