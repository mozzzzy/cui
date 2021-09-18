package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {

	cui.Info("This is demo code of checkbox package.")

	q := "What language do you like?"
	choices := []string{
		"c",
		"c++",
		"go",
		"java",
		"javascript",
		"php",
		"ruby",
	}
	answers, canceled := cui.Checkbox(q, choices)

	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
		return
	}

	cui.Info("Following choices are checked...")
	for _, a := range answers {
		cui.Info(choices[a])
	}
	time.Sleep(3 * time.Second)
	cui.Erase()
	return
}
