package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/lyderic/tools"
)

type Configuration struct {
	GoVersionToInstall string
	GoInstalledVersion string
	Archive            string
	LocalArchivePath   string
	Url                string
	OS                 string
	Arch               string
}

var config Configuration

func init() {
	config.OS = runtime.GOOS
	config.Arch = runtime.GOARCH
}

func main() {
	if config.OS != "linux" {
		fmt.Println("This program works only on Linux!")
		return
	}
	config.GoInstalledVersion = getGoInstalledVersion()
	if len(os.Args) == 1 {
		usage()
		return
	}
	config.GoVersionToInstall = "go" + os.Args[1]
	if config.GoVersionToInstall == config.GoInstalledVersion {
		fmt.Println("Version", config.GoInstalledVersion, "already installed")
		return
	}
	config.Archive = fmt.Sprintf("%s.%s-%s.tar.gz",
		config.GoVersionToInstall,
		config.OS,
		tools.Ternary(config.Arch == "arm", "armv6l", config.Arch))
	config.LocalArchivePath = "/tmp/" + config.Archive
	config.Url = fmt.Sprintf("https://dl.google.com/go/%s", config.Archive)
	download()
	testArchive()
	cleango()
	unarchive()
	chownAll()
	fmt.Println("Go version is now:", getGoInstalledVersion())
}

func cleango() {
	cmd := exec.Command("sudo", "rm", "-rf", "/usr/local/go")
	xeqv(cmd)
}

func download() {
	err := os.Chdir("/tmp")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("curl", "-L", "-O", config.Url)
	err = xeqv(cmd)
	if err != nil {
		panic(err)
	}
}

func testArchive() {
	cmd := exec.Command("tar", "tf", config.LocalArchivePath)
	err := xeq(cmd, false)
	if err != nil {
		fmt.Println("Invalid archive:", config.LocalArchivePath)
		os.Remove(config.LocalArchivePath)
		fmt.Println("ABORTING")
		os.Exit(23)
	}
}

func unarchive() {
	cmd := exec.Command("sudo", "tar", "xzf", "/tmp/"+config.Archive, "-C", "/usr/local")
	err := xeqv(cmd)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	os.Remove(config.LocalArchivePath)
	if err != nil {
		os.Exit(23)
	}
}

func chownAll() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("sudo", "chown", "-R", user.Username+":", "/usr/local/go")
	xeqv(cmd)
}

func getGoInstalledVersion() string {
	cmd := exec.Command("/usr/local/go/bin/go", "version")
	raw, err := cmd.Output()
	if err != nil {
		return "none"
	}
	output := strings.TrimSpace(string(raw))
	if len(output) == 0 {
		return "none"
	}
	return strings.Split(output, " ")[2]
}

func usage() {
	fmt.Println("Currently installed go version:", config.GoInstalledVersion)
	fmt.Println("Usage: goset <go version>")
	fmt.Println("example: goset 1.11.5")
}

func xeqv(cmd *exec.Cmd) error {
	return xeq(cmd, true)
}

func xeq(cmd *exec.Cmd, verbose bool) error {
	fmt.Print("XEQ")
	fmt.Println(cmd.Args)
	if verbose {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	return cmd.Run()
}
