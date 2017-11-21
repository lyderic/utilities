package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const file = "/dev/shm/bash-counter.run"

func init() {
	log.SetFlags(log.Lshortfile)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("File", file, "not found. Initialising...")
		f, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString("0")
	}
}

func main() {

	if len(os.Args) == 1 {
		report()
		return
	}

	switch os.Args[1] {
	case "--login":
		login()
	case "--logout":
		logout()
	default:
		log.Fatalln("Invalid!")
	}
}

func login() {
	value := computeNew("+")
	writeFile(value)
	report()
}

func logout() {
	value := computeNew("-")
	// value cannot be negative
	if value < 0 {
		value = 0
	}
	if value == 0 {
    fmt.Print("\033[31m") // red
    fmt.Println("Bash counter has reached 0! vmount dismount initiated...")
    fmt.Print("\033[0m") // reset color
		vmountDismount()
	}
	writeFile(value)
	report()
	time.Sleep(time.Second) // to have a chance to read it
}

func writeFile(value int) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%d", value))
	if err != nil {
		log.Fatal(err)
	}
}

func computeNew(operation string) int {
	current := getFileContent()
	switch operation {
	case "+":
		return current + 1
	case "-":
		return current - 1
	default:
		log.Fatalln(operation, ": invalid operation!")
	}
	return current
}

func getFileContent() int {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	counter, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatal(err)
	}
	return counter
}

func report() {
	fmt.Print("\033[?25l") // hide cursor
	fmt.Print("\033[32m")  // green
	//fmt.Print("\033[7m")   // reverse
	message := fmt.Sprintf("Bash counter: %d", getFileContent())
	fmt.Println(message)
	fmt.Print(vmountStatus())
	fmt.Print("\033[0m") // reset color
	fmt.Print("\r")
	for i := 0; i < len(message); i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
	fmt.Print("\033[?25h") // show cursor
}

func vmountStatus() string {
	output, err := exec.Command("vmount").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func vmountDismount() {
	cmd := exec.Command("vmount", "--dismount")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
