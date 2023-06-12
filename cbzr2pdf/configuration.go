package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/lyderic/tools"
)

type Configuration struct {
	Path string
	Dir  string
	Ext  string
	Pdf  string
	Temp string
}

func (c Configuration) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Path: %q\n", c.Path)
	fmt.Fprintf(&b, "Dir:  %q\n", c.Dir)
	fmt.Fprintf(&b, "Ext:  %q\n", c.Ext)
	fmt.Fprintf(&b, "Pdf:  %q\n", c.Pdf)
	fmt.Fprintf(&b, "Temp: %q\n", c.Temp)
	return b.String()
}

func initConfig(path string) (c Configuration) {
	var err error
	if c.Path, err = filepath.Abs(path); err != nil {
		E(err)
	}
	c.Dir = filepath.Dir(c.Path)
	c.Ext = filepath.Ext(c.Path)
	c.Pdf = strings.Replace(c.Path, c.Ext, ".pdf", 1)
	var tempdir string
	if tempdir, err = os.MkdirTemp(TEMPBASE, os.Args[0]); err != nil {
		E(err)
	}
	c.Temp = tempdir
	return
}
