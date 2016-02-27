package storage

import (
	"regexp"
	"runtime"
)

func HandleString(text string) {
	// compile rexular expression
	ex, err := regexp.Compile("[^\\s]+")
	if err != nil {
		runtime.Goexit()
	}
	// get words from string
	words := ex.FindAllString(text, -1)
	if len(words) == 0 {
		runtime.Goexit()
	}

	// send to channel for sorting
	stringsChannel <- words
	runtime.Goexit()
}
