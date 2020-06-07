package main

import (
	"github.com/mozzzzy/cui"
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
		return
	}

	cui.Info("Answer is " + choices[answer] + ".")
}
