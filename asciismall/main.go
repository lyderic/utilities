package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	. "github.com/lyderic/tools"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	if len(os.Args) > 1 {
		for _, argument := range os.Args[1:] {
			process(argument)
		}
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text())
	}
}

func process(input string) {
	fmt.Println(ToAsciiSmall(input))
}
