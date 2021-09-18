package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of checkbox package.")

	q := cui.Question("What language do you like?")
	choices := []string{
		"c",
		"c++",
		"go",
		"java",
		"javascript",
		"php",
		"ruby",
	}
	answers, canceled, checkbox := cui.Checkbox(choices)
	checkbox.Erase()

	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
		return
	}

	answerStr := ""
	for index, a := range answers {
		answerStr += choices[a]
		if index < len(answers) - 1 {
			answerStr += ","
		}
	}
	q.SetAnswer(answerStr)
	time.Sleep(3 * time.Second)
	cui.Erase()
	return
}
