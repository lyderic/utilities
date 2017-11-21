package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	loginPtr := flag.Bool("login", false, "login stuff")
	logoutPtr := flag.Bool("logout", false, "logout stuff")
	flag.Parse()

	login := *loginPtr
	logout := *logoutPtr

	// login and logout are mutually exclusive
	if login && logout {
		log.Fatalln("Please make up your mind! Login or logout")
	}

	if login {
		doLogin()
	} else if logout {
		doLogout()
	} else {
		information()
	}

}

// Add this at the end of $HOME/.bashrc (provided it is read at login, you can
// also use .bash_login or .profile, etc.
// /path/to/bashinout --login
func doLogin() {
	fmt.Println("You are logging in to bash")
}

// At the end of $HOME/.bash_logout
// /path/to/bashinout --logout
func doLogout() {
	fmt.Print("You are logging out of bash")
	time.Sleep(time.Second)
	fmt.Print("\r                             \r")
}

// display information/usage/whatever, in case no login nor logout is called
func information() {
	fmt.Println("This is bashinout, a skeleton application only")
}
