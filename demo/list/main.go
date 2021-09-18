package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	q := "What is your favorite language?"
	choices := []string{
		"c++",
		"go",
		"javascript",
		"php",
	}
	answer, canceled := cui.List(q, choices)
	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
	} else {
		cui.Info("Answer is \"" + choices[answer] + "\".")
		time.Sleep(3 * time.Second)
	}
	cui.Erase()
}
