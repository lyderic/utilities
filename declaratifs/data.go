package main

import (
	"fmt"
	"strings"

	. "github.com/lyderic/tools"
)

type Category struct {
	Description string   `yaml:"description"`
	Examples    []string `yaml:"exemples"`
}

func (c Category) String() string {
	var buffer strings.Builder
	ln := len(c.Examples)
	fmt.Fprintf(&buffer, "%s (%d):\n", c.Description, ln)
	for idx, example := range c.Examples {
		fmt.Fprintf(&buffer, "%s", example)
		if idx < ln-1 {
			fmt.Fprintf(&buffer, ", ")
		}
	}
	fmt.Fprintf(&buffer, "\n")
	return buffer.String()
}

type Entry struct {
	Verb string   `yaml:"verbe"`
	Tags []string `yaml:"identifiants"`
}

type Verb struct {
	Name       string     `yaml:"name"`
	Categories []Category `yaml:"categories"`
}

func (e Entry) String() string {
	var buffer strings.Builder
	fmt.Fprintf(&buffer, "%s: ", e.Verb)
	if dbg {
		Cyan("#:%d %v\n", len(e.Tags), e.Tags)
	}
	if len(e.Tags) > 0 {
		for _, tag := range e.Tags {
			fmt.Fprintf(&buffer, "%s ", tag)
		}
	} else {
		fmt.Fprintf(&buffer, "-- no tags --")
	}
	return buffer.String()
}
