package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		processFileArguments()
	} else {
		processPipedInputFromStdin()
	}
}

func processFileArguments() {
	var b strings.Builder
	for _, file := range os.Args[1:] {
		if raw, err := os.ReadFile(file); err != nil {
			fmt.Printf("Error processing %q: %v\n", file, err)
			continue
		} else {
			fmt.Fprintf(&b, "%s\n", string(raw))
		}
	}
	countOccurences(&b)
}

func processPipedInputFromStdin() {
	var b strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(&b, "%s\n", line)
	}
	if err := scanner.Err(); err != nil {
		return
	}
	countOccurences(&b)
}

func countOccurences(buffer *strings.Builder) {
	content := buffer.String()
	var cleaner = regexp.MustCompile(`[^\p{L}\p{N}' ]+`)
	dict := make(map[string]int)
	for _, word := range strings.Fields(string(content)) {
		word = cleaner.ReplaceAllString(word, "")
		word = strings.ToLower(word)
		if len(word) > 0 {
			dict[word] = dict[word] + 1
		}
	}
	keys := make([]string, 0, len(dict))
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return dict[keys[i]] > dict[keys[j]]
	})
	for _, key := range keys {
		fmt.Printf("% 5d %s\n", dict[key], key)
	}
}
