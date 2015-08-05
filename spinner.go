package fancytext

import (
	"fmt"
	"time"
)

var (
	BASIC_SPINNER   = []rune("-\\|/")
	LOADING_SPINNER = []rune("█▉▊▋▌▍▎▏ ▏▎▍▌▋▊▉█")
	SPINNING_WHEEL  = []rune("◴◷◶◵")
)

func FormatSpinner(format string) func() {
	done := make(chan bool)
	go syncFormatSpinner(format, SPINNING_WHEEL, done)

	return func() {
		done <- true
	}
}

func Spinner() func() {
	return FormatSpinner("[%s] - working!")
}

func TopLeftFormatSpinner(format string) func() {
	return FormatSpinner("[0;0H[K" + format)
}

func syncFormatSpinner(format string, chars []rune, done chan bool) {
	index := 0
	for {
		index += 1
		char := chars[index%len(chars)]
		select {
		case _ = <-done:
			fmt.Printf("\r[K")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(format, string(char))
			time.Sleep(time.Second / 8)
		}
	}
}
