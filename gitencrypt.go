/* In order for this to work, a ~/.gitencrypt directory needs to be created
and files put in it. Look here for more explanation:
https://gist.github.com/shadowhand/873637 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	gitencryptDir = filepath.Join(os.Getenv("HOME"), ".gitencrypt")
	verbose       bool
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var err error
	if _, err = os.Stat(gitencryptDir); os.IsNotExist(err) {
		log.Fatal("cannot use gitencrypt without this directory: %s\n", gitencryptDir)
	}
	flag.BoolVar(&verbose, "v", false, "be verbose")
	flag.Parse()
	ln := len(flag.Args()) // no command provided
	if ln == 0 {
		usage()
		return
	}
	switch flag.Args()[0] {
	case "init":
		dir := "."
		if ln == 2 {
			dir = flag.Args()[1]
		}
		if err = gitinit(dir); err != nil {
			log.Fatal(err)
		}
	case "clone":
		if err = clone(); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("%s: invalid command", os.Args[1])
		return
	}
}

/* use instead of 'git init' to initiate an encrypted git repository */
func gitinit(dir string) (err error) {
	gitdir := filepath.Join(dir, ".git")
	if _, err = os.Stat(gitdir); os.IsNotExist(err) {
		if err = git("init", dir); err != nil {
			return
		}
		return setBlobs(dir)
	}
	fmt.Println(gitdir, "already created")
	return
}

func git(args ...string) (err error) {
	cmd := exec.Command("git", args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

func clone() (err error) {
	if len(flag.Args()) < 2 {
		return fmt.Errorf("missing url to clone")
	}
	args := flag.Args()
	url := args[1]
	repository := getRepository(args...)
	if err = git("clone", "-n", url, repository); err != nil {
		return
	}
	if err = setBlobs(repository); err != nil {
		return
	}
	return git("-C", repository, "reset", "--hard", "HEAD")
}

func setBlobs(dir string) (err error) {
	/* repo/.git/config */
	gitConfig := filepath.Join(dir, ".git", "config")
	if err = setBlob(blobGitConfig, gitConfig); err != nil {
		return
	}
	/* repo/.gitattributes */
	gitattributes := filepath.Join(dir, ".git", "info", "attributes")
	if err = setBlob(blobGitattributes, gitattributes); err != nil {
		return
	}
	return
}

/* append blob to file, if file does not exist, create it */
func setBlob(blob, file string) (err error) {
	/* if the file doesn't exist, create it and write the blob */
	if _, err = os.Stat(file); os.IsNotExist(err) {
		return ioutil.WriteFile(file, []byte(blob), 0644)
	}
	/* if the file exists, does it already contain the blob? */
	var bb []byte
	if bb, err = ioutil.ReadFile(file); err != nil {
		return
	}
	content := string(bb)
	if strings.Contains(content, blob) {
		/* nothing to do, blob is already in file */
		fmt.Println("blob found in", file)
		return
	}
	content = content + blob
	return ioutil.WriteFile(file, []byte(content), 0644)
}

func getRepository(args ...string) string {
	/* the URL can take several forms, for example:
		urls := []string{"https://path/to/repo.git",
			"http://path/to/another/repo",
			"file:///path/to/local/repo",
			"simple:repo",
			"simple:repo.git",
	    "ssh://complex:repo",
	    "ssh://complex:repo.git",
	    "ssh://with/a/path/repo"}*/
	n := len(args)
	if n > 2 {
		return args[2]
	}
	url := args[1]
	ln := len(url)
	if strings.HasSuffix(url, ".git") {
		url = url[:ln-4]
	}
	last := path.Base(url)
	if !strings.Contains(last, ":") {
		return last
	}
	return strings.Split(last, ":")[1]
}

func usage() {
	fmt.Printf("Usage: %s [init|clone]\n", path.Base(os.Args[0]))
	fmt.Println(" init [<dir>]      initiate encrypted git repository")
	fmt.Println(" clone URL [<dir>] clone remote encrypted git repository")
}

var blobGitConfig = `[filter "openssl"]
	smudge = ~/.gitencrypt/smudge_filter_openssl
	clean = ~/.gitencrypt/clean_filter_openssl
[diff "openssl"]
	textconv = ~/.gitencrypt/diff_filter_openssl
`

var blobGitattributes = `* filter=openssl diff=openssl
[merge]
	renormalize=true
`
