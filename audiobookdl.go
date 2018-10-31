package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	unset = ""
)

type Config struct {
	Dest    string
	Url     string
	Title   string
	Author  string
	Mp3     string
	Mp3Path string
}

var c Config

func main() {

	flag.StringVar(&c.Dest, "d", ".", "Destination `directory`")
	flag.StringVar(&c.Url, "u", unset, "YouTube `URL`")
	flag.StringVar(&c.Title, "t", unset, "Audiobook's `title`")
	flag.StringVar(&c.Author, "a", unset, "Audiobook's `author`")
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.PrintDefaults()
		return
	}
	if _, err := os.Stat(c.Dest); os.IsNotExist(err) {
		fmt.Println("Destination directory not found:", c.Dest)
		return
	}
	if c.Url == unset || c.Title == unset || c.Author == unset {
		fmt.Println("Missing url, title or author!")
		return
	}
	c.Mp3 = fmt.Sprintf("%s - %s.mp3", c.Title, c.Author)
	c.Mp3Path = filepath.Join(c.Dest, c.Mp3)

	//fmt.Println(c)

	download()
	tag()
}

func download() {
	fmt.Println("download")
	output := fmt.Sprintf("%s/%s - %s.%%(ext)s",
		c.Dest, c.Title, c.Author)
	cmd := exec.Command("youtube-dl",
		"--extract-audio",
		"--audio-format", "mp3",
		"--audio-quality", "64k",
		"--output", output,
		c.Url)
	//fmt.Println(cmd)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	cmd.Run()
}

func tag() {
	fmt.Println("tag")
	cmd := exec.Command("id3v2",
		"--artist", c.Author,
		"--album", c.Title,
		"--song", c.Title,
		c.Mp3Path)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	cmd.Run()
}
