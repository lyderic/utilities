package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/lyderic/tools"
)

const (
	src = "cube:/stor"
	dst = "/stor"
)

func init() {
	err := tools.CheckBinaries("sshfs", "fusermount")
	if err != nil {
		panic(err)
	}
}

func main() {
	check_host("nuc", "tintin", "penguin")
	if !is_mounted() {
		mount()
	} else {
		reader := bufio.NewReader(os.Stdin)
		tools.PrintYellowf("%s is mounted. Unmount (y/N)? ", dst)
		ans, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if len(ans) == 0 {
			return
		}
		ans = strings.TrimSpace(strings.ToLower(ans))
		if strings.HasPrefix(ans, "y") {
			unmount()
		}
	}
}

func check_host(hosts ...string) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	for _, host := range hosts {
		if hostname == host {
			return
		}
	}
	tools.PrintRedf("this program works only on %v! you are on '%s'.\n", hosts, hostname)
	os.Exit(23)
}

func is_mounted() bool {
	cmd := exec.Command("mountpoint", "-q", dst)
	err := cmd.Run()
	return err == nil
}

func mount() {
	cmd := exec.Command("sshfs", src, dst)
	err := cmd.Run()
	if err == nil {
		tools.PrintGreenf("\r%s has been mounted successfully.\n", dst)
	} else {
		tools.PrintRedf("failed!\n")
	}
}

func unmount() {
	cmd := exec.Command("fusermount", "-u", dst)
	err := cmd.Run()
	if err == nil {
		tools.PrintGreenf("\r%s has been unmounted successfully.\n", dst)
	} else {
		tools.PrintRedf("failed!\n")
	}
}
