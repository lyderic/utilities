package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	quiet = false
)

func main() {
	flag.BoolVar(&quiet, "q", false, "be quiet")
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
		return
	}
	host := flag.Args()[0]
	code := 0
	if err := ping(host); err == nil {
		display(host + " is pingable")
	} else {
		display(host + " is not available")
		code = 1
	}
	os.Exit(code)
}

func ping(host string) error {
	cmd := exec.Command("ping", "-c", "1", "-w", "1", host)
	return cmd.Run()
}

func display(message string) {
	if !quiet {
		fmt.Println(message)
	}
}

func usage() {
	fmt.Println(`Usage: catena [-q] <host>
test if a host is pingable`)
	flag.PrintDefaults()
}
