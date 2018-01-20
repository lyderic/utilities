package main

import (
	"flag"
	"fmt"
	"time"
)

func init() {
}

func main() {

	anneePtr := flag.Bool("a", false, "Affiche l'année")
	nohashPtr := flag.Bool("n", false, "Ne pas afficher '##'")
	nolinePtr := flag.Bool("e", false, "Ne pas insérer une ligne avant")
	flag.Parse()
	annee := *anneePtr
	nohash := *nohashPtr
	noline := *nolinePtr

	now := time.Now()

	if !noline {
		fmt.Println()
	}

	if !nohash {
		fmt.Print("##")
	}

	fmt.Printf("%s %d %s ", wday_fr[now.Weekday().String()],
		now.Day(), month_fr[now.Month().String()])

	if annee {
		fmt.Println(now.Year())
	} else {
		fmt.Println()
	}

}

var month_fr = map[string]string{
	"January":   "janvier",
	"February":  "février",
	"March":     "mars",
	"April":     "avril",
	"May":       "mai",
	"June:":     "juin",
	"July":      "juillet",
	"August":    "août",
	"September": "septembre",
	"October":   "octobre",
	"November":  "novembre",
	"December":  "décembre"}
var wday_fr = map[string]string{
	"Sunday":    "Dimanche",
	"Monday":    "Lundi",
	"Tuesday":   "Mardi",
	"Wednesday": "Mercredi",
	"Thursday":  "Jeudi",
	"Friday":    "Vendredi",
	"Saturday":  "Samedi"}
