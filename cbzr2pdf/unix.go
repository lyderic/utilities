package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	. "github.com/lyderic/tools"
)

func init() {
	if err := CheckBinaries("7z", "img2pdf", "pdftk"); err != nil {
		E(err)
	}
}

func extractcbzr(c Configuration) {
	fmt.Printf("Extracting %q... ", c.Path)
	os.Chdir(c.Temp)
	defer os.Chdir(c.Dir)
	cmd := exec.Command("7z", "e", c.Path)
	debug("\nXeQ%v\n", cmd.Args)
	if *SHOWCMD {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if *DRYRUN {
		return
	}
	if err := cmd.Run(); err != nil {
		E(err)
	}
	Greenln("ok")
}

func images2pdf(c Configuration) {
	var listing []fs.DirEntry
	var err error
	if listing, err = os.ReadDir(c.Temp); err != nil {
		E(err)
	}
	var wg sync.WaitGroup
	for _, entry := range listing {
		wg.Add(1)
		go image2pdf(&wg, filepath.Join(c.Temp, entry.Name()))
	}
	wg.Wait()
}

func image2pdf(wg *sync.WaitGroup, img string) {
	defer wg.Done()
	var finfo fs.FileInfo
	var err error
	finfo, err = os.Stat(img)
	if err != nil {
		Yellow("cannot stat image: %q\n", img)
		return
	}
	if finfo.IsDir() {
		return
	}
	if strings.ToLower(filepath.Ext(img)) != ".jpg" {
		Yellow("%q: not a valid image, skipping\n", img)
		return
	}
	pdf := img + ".pdf"
	cmd := exec.Command("img2pdf", img, "-o", pdf)
	debug("\nXeQ%v\n", cmd.Args)
	if *SHOWCMD {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if *DRYRUN {
		return
	}
	if err = cmd.Run(); err != nil {
		Yellow("failed to convert image %q to pdf: %#v\n", img, err)
	} else {
		fmt.Printf("Image %q successfully converted to pdf\n", img)
	}
}

func createpdf(c Configuration) {
	os.Chdir(c.Temp)
	defer os.Chdir(c.Dir)
	var listing []fs.DirEntry
	var err error
	if listing, err = os.ReadDir("."); err != nil {
		E(err)
	}
	var args []string
	for _, entry := range listing {
		if filepath.Ext(entry.Name()) == ".pdf" {
			args = append(args, entry.Name())
		}
	}
	if len(args) == 0 {
		Yellow("no images found in %q for %q. aborting\n", c.Temp, c.Pdf)
		return
	}
	args = append(args, "cat", "output", c.Pdf)
	cmd := exec.Command("pdftk", args...)
	debug("XeQ%v\n", cmd.Args)
	if *SHOWCMD {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if *DRYRUN {
		return
	}
	if cmd.Run(); err != nil {
		E(err)
	}
	Green("PDF %q successfully created\n", c.Pdf)
}
