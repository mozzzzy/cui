package main

import (
	"github.com/mozzzzy/cui"
)

func main() {
	q := "What language do you like?"
	choices := []string{
		"c++",
		"go",
		"javascript",
		"php",
	}
	answers, canceled := cui.Checkbox(q, choices)

	if canceled {
		cui.Warn("Canceled.")
		return
	}
	cui.Info("Following choices are ckecked...")
	for _, a := range answers {
		cui.Info(a)
	}
}
