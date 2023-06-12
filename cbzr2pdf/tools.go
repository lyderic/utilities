package main

import (
	. "github.com/lyderic/tools"
)

func debug(format string, args ...interface{}) {
	if *DEBUG {
		Cyan(format, args...)
	}
}
