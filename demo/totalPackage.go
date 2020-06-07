package main

import (
	"time"

	"github.com/mozzzzy/cui/v2"
	"github.com/mozzzzy/cui/v2/color"
)

func main() {
	choices := []string{
		"message",
		"debug",
		"info",
		"notice",
		"warn",
		"error",
		"input",
		"secureInput",
		"confirmation",
		"list",
		"progressBar",
		"spinner",
	}
	answers, canceled := cui.Checkbox("Which demonstration do you run?", choices)

	if canceled {
		cui.Warn("Canceled")
		return
	}

	for _, a := range answers {
		switch a {
		case 0:
			cui.Message("This is test message!!", []string{color.BlueFg, color.Bold})
		case 1:
			cui.Debug("This is debug message!!")
		case 2:
			cui.Info("This is info message!!")
		case 3:
			cui.Notice("This is notice message!!")
		case 4:
			cui.Warn("This is warning message!!")
		case 5:
			cui.Error("This is error message!!")
		case 6:
			answer, canceled := cui.Input("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case 7:
			answer, canceled := cui.SecureInput("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case 8:
			answer, canceled := cui.Confirmation("Please type Y or n and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			if answer {
				cui.Info("Accepted.")
			} else {
				cui.Warn("Rejected.")
			}
		case 9:
			osNames := []string{"Windows", "Linux", "MacOS"}
			answer, canceled := cui.List("Which operation system do you like?", osNames)
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + osNames[answer] + "\"")
		case 10:
			pb0 := cui.ProgressBar("Waiting some operations... (success case)")
			for i := 1; i <= 100; i++ {
				time.Sleep(100 * time.Millisecond)
				pb0.ReportProgress(i)
			}
			pb1 := cui.ProgressBar("Waiting some operations... (failure case)")
			for i := 1; i <= 55; i++ {
				time.Sleep(100 * time.Millisecond)
				pb1.ReportProgress(i)
			}
			pb1.Failure()
		case 11:
			spnr0 := cui.Spinner("Waiting some operations... (success case)")
			time.Sleep(3 * time.Second)
			spnr0.Complete()

			spnr1 := cui.Spinner("Waiting some operations... (failure case)")
			time.Sleep(3 * time.Second)
			spnr1.Failure()
		}
	}
}
