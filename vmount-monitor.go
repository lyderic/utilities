package main

import (
	"fmt"
  "time"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	case "--plus":
		plus()
	case "--minus":
		minus()
	default:
		log.Fatalln("Invalid!")
	}
}

func plus() {
	value := computeNew("+")
	writeFile(value)
  report()
}

func minus() {
	value := computeNew("-")
	// value cannot be negative
	if value < 0 {
		value = 0
	}
	writeFile(value)
  report()
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
	fmt.Print("Bash counter: ", getFileContent())
  time.Sleep(time.Second)
  fmt.Print("\r                                      \r")
}
