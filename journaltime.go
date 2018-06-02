package main

import (
	"flag"
	"fmt"
	"time"
  "strings"
)

var dbg bool

type Options struct {
	ShowYear      bool
	EmptyLinesBefore    int
	EmptyLinesAfter    int
	PrependNumber int
	PrependSymbol string
	Language      string
}

var o Options

func main() {
	flag.BoolVar(&dbg, "debug", false, "debug mode")
	flag.BoolVar(&o.ShowYear, "y", false, "show year")
	flag.IntVar(&o.EmptyLinesBefore, "b", 0, "number of lines to insert before displaying date")
	flag.IntVar(&o.EmptyLinesAfter, "a", 0, "number of lines to insert after displaying date")
	flag.IntVar(&o.PrependNumber, "n", 0, "number of symbols to prepend to date")
	flag.StringVar(&o.PrependSymbol, "s", "#", "symbol to prepend to date")
	flag.StringVar(&o.Language, "l", "fr", "language")
	flag.Parse()
	debug("o: %#v\n", o)
	display(o)
}

func display(o Options) {
	now := time.Now()
  for i:=0;i<o.EmptyLinesBefore;i++ {
    fmt.Println()
  }
  if o.PrependNumber > 0 {
    fmt.Print(strings.Repeat(o.PrependSymbol, o.PrependNumber))
  }
	switch o.Language {
	case "fr":
		fmt.Printf("%s %d %s", wday_fr[now.Weekday().String()],
			now.Day(), month_fr[now.Month().String()])
	case "de":
		fmt.Printf("%s %d %s", wday_de[now.Weekday().String()],
			now.Day(), month_de[now.Month().String()])
	default:
		fmt.Printf("%s %d %s", now.Weekday(), now.Day(), now.Month())
	}
  if o.ShowYear {
    fmt.Printf(" %d\n", now.Year())
  } else {
    fmt.Println()
  }
  for i:=0;i<o.EmptyLinesAfter;i++ {
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
	"June":     "juin",
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
