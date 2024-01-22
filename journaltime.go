package main

import (
	"flag"
	"fmt"
	"time"
)

var dbg bool

type Options struct {
	Input            string
	ShowYear         bool
	ShowTime         bool
	EmptyLinesBefore int
	EmptyLinesAfter  int
	Prepend          string
	Language         string
}

var o Options

func main() {
	flag.BoolVar(&dbg, "debug", false, "debug mode")
	flag.StringVar(&o.Input, "i", "", "input time in the form '%d%m%H%M' (default: now)")
	flag.BoolVar(&o.ShowTime, "t", false, "show time")
	flag.BoolVar(&o.ShowYear, "y", false, "show year")
	flag.IntVar(&o.EmptyLinesBefore, "b", 0, "number of lines to insert before displaying date")
	flag.IntVar(&o.EmptyLinesAfter, "a", 0, "number of lines to insert after displaying date")
	flag.StringVar(&o.Prepend, "s", "", "string to prepend to date")
	flag.StringVar(&o.Language, "l", "fr", "language")
	flag.Parse()
	debug("o: %#v\n", o)
	display(o)
}

func display(o Options) {
	var datetime time.Time
	var err error
	now := time.Now()
	yearString := fmt.Sprintf("%d", now.Year())
	if o.Input == "" {
		datetime = now
	} else {
		if datetime, err = time.Parse("020115042006", o.Input+yearString); err != nil {
			fmt.Printf("Cannot parse input: %q\n", o.Input)
			return
		}
	}
	debug("raw datetime: %#v\n", datetime)
	for i := 0; i < o.EmptyLinesBefore; i++ {
		fmt.Println()
	}
	if o.Prepend != "" {
		fmt.Print(o.Prepend)
	}
	switch o.Language {
	case "fr":
		fmt.Printf("%s %d %s", wday_fr[datetime.Weekday().String()],
			datetime.Day(), month_fr[datetime.Month().String()])
	case "de":
		fmt.Printf("%s %d %s", wday_de[datetime.Weekday().String()],
			datetime.Day(), month_de[datetime.Month().String()])
	default:
		fmt.Printf("%s %d %s", datetime.Weekday(), datetime.Day(), datetime.Month())
	}
	if o.ShowYear {
		fmt.Printf(" %d", datetime.Year())
	}
	if o.ShowTime {
		fmt.Printf(" %02d:%02d", datetime.Hour(), datetime.Minute())
	}
	fmt.Println()
	for i := 0; i < o.EmptyLinesAfter; i++ {
		fmt.Println()
	}
}

func debug(format string, args ...interface{}) {
	format = "[DEBUG] " + format
	if dbg {
		fmt.Printf(format, args...)
	}
}

var month_fr = map[string]string{
	"January":   "janvier",
	"February":  "février",
	"March":     "mars",
	"April":     "avril",
	"May":       "mai",
	"June":      "juin",
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

var month_de = map[string]string{
	"January":   "Januar",
	"February":  "Februar",
	"March":     "März",
	"April":     "April",
	"May":       "Mai",
	"June:":     "Juni",
	"July":      "Juli",
	"August":    "August",
	"September": "September",
	"October":   "Oktober",
	"November":  "November",
	"December":  "Dezember"}
var wday_de = map[string]string{
	"Sunday":    "Sonntag",
	"Monday":    "Montag",
	"Tuesday":   "Dienstag",
	"Wednesday": "Mittwoch",
	"Thursday":  "Donnerstag",
	"Friday":    "Freitag",
	"Saturday":  "Samstag"}
