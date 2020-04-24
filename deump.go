package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	keys []string
)

func init() {
	keys = []string{
		"AWS",
		"B2",
		"CLOUDFLARE",
		"GO",
		"GPG",
		"LPASS",
		"LPPK",
		"OBSCURA",
		"PASSWORD_STORE",
		"RCLONE",
		"RESTIC",
		"SSA",
		"SSH",
		"SSK",
		"SUBB",
		"TMOUT",
		"VAULT",
		"p",
	}
}

func main() {
	deump()
}

func deump() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, envvar := range os.Environ() {
		for _, key := range keys {
			if strings.HasPrefix(envvar, key) {
				bits := strings.Split(envvar, "=")
				fmt.Fprintf(w, "%s\t%s\n", bits[0], bits[1])
			}
		}
	}
	w.Flush()
}
