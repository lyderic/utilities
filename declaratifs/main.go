package main

import (
	_ "embed"
	"fmt"

	. "github.com/lyderic/tools"

	"gopkg.in/yaml.v2"
)

//go:embed categories.yaml
var categoriesYaml []byte

//go:embed verbes.yaml
var verbesYaml []byte

var dbg = true

func main() {
	if dbg {
		Cyanln("*** DEBUG ON ***")
	}
	var err error
	var entries []Entry
	if entries, err = getEntries(); err != nil {
		return
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func getEntries() (entries []Entry, err error) {
	if err = yaml.Unmarshal(verbesYaml, &entries); err != nil {
		return
	}
	return
}

func listCategories() (err error) {
	var categories []Category
	if err = yaml.Unmarshal(categoriesYaml, &categories); err != nil {
		return
	}
	for _, c := range categories {
		fmt.Println(c)
	}
	return
}
