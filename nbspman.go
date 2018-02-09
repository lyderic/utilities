package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	appname = "nbspman"
	version = "0.0.1"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s v. %s - HTML &nbsp; manager (c) Lyderic Landry\n", appname, version)
		fmt.Printf("Usage: %s <file> [<outfile>]\n", appname)
		fmt.Println("This program add unbreakable spaces (&nbsp;) to HTML text, where it is needed")
		fmt.Println("in French (e.g. before !, ?, : or after «)")
		return
	}
	infile, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	outfile := infile + ".out"
	if len(os.Args) == 3 {
		outfile = os.Args[2]
	}
	process(infile, outfile)
}

func process(infile, outfile string) {
	fmt.Printf("Output file: %s\n", outfile)
	buffer, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes - ", len(buffer))
	content := string(buffer)
	fmt.Printf("Content length: %d\n", len(content))
	fmt.Println("Processing...")
	fmt.Print("Normalizing whitespaces... ")
	normalizedWhitespaces := strings.Replace(content, "\u00a0", " ", -1)
	fmt.Println("done.")
	fmt.Print("Adding '&nbsp;'... ")
	r := strings.NewReplacer(
		" :", "&nbsp;:",
		" ?", "&nbsp;?",
		" !", "&nbsp;!",
		"« ", "«&nbsp;",
		" »", "&nbsp;»",
		" %", "&nbsp;%")
	handler, herr := os.Create(outfile)
	if herr != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(handler)
	n, werr := r.WriteString(writer, normalizedWhitespaces)
	if werr != nil {
		log.Fatal(err)
	}
	writer.Flush()
	fmt.Println("done.")
	fmt.Printf("Byte count after replacements: %d\n", n)
}
