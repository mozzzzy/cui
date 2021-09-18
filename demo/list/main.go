package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	q := cui.Question("What is your favorite language?")
	choices := []string{
		"c++",
		"go",
		"javascript",
		"php",
	}
	answer, canceled, list := cui.List(choices)
	list.Erase()

	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
		return
	}

	q.SetAnswer(choices[answer])
	time.Sleep(3 * time.Second)
	cui.Erase()
}
