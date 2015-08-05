package fancytext

import (
	"fmt"
	"time"
)

func FormatSpinner(format string) func() {
	done := make(chan bool)
	go syncFormatSpinner(format, done)

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

func syncFormatSpinner(format string, done chan bool) {
	index := 0
	spinner := "-\\|/"
	for {
		index += 1
		char := spinner[index%len(spinner)]
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
