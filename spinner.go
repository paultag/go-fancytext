package fancytext

import (
	"fmt"
	"time"
)

var (
	basicSpinner   = []rune("-\\|/")
	loadingSpinner = []rune("█▉▊▋▌▍▎▏ ▁▂▃▄▅▆▇")
	circleSpinner  = []rune("◴◷◶◵")
)

func BooleanFormatSpinner(format string) func(bool) {
	done := make(chan bool)
	go syncFormatSpinner(format, loadingSpinner, (time.Second / 8), done)

	return func(good bool) {
		if done == nil {
			return
		}
		done <- true
		done = nil
		char := "✓"
		if !good {
			char = "✗"
		}
		fmt.Printf(format+"\n", char)
	}
}

func FormatSpinner(format string) func() {
	done := make(chan bool)
	go syncFormatSpinner(format, loadingSpinner, (time.Second / 8), done)

	return func() {
		if done == nil {
			return
		}
		done <- true
		done = nil
	}
}

func Spinner() func() {
	return FormatSpinner("[%s] - working!")
}

func TopLeftFormatSpinner(format string) func() {
	return FormatSpinner("[0;0H[K" + format)
}

func syncFormatSpinner(format string, chars []rune, speed time.Duration, done chan bool) {
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
			time.Sleep(speed)
		}
	}
}
