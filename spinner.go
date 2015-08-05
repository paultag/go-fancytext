package fancytext

import (
	"fmt"
	"time"
)

const (
	BASIC_SPINNER = "-\\|/"
)

func FormatSpinner(format string) func() {
	done := make(chan bool)
	go syncFormatSpinner(format, BASIC_SPINNER, done)

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

func syncFormatSpinner(format string, chars string, done chan bool) {
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
