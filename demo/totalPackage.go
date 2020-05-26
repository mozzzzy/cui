package main

import (
	"time"

	"github.com/mozzzzy/cui"
	"github.com/mozzzzy/cui/color"
)

func main() {
	answers, canceled := cui.Checkbox("Which demonstration do you run?", []string{
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
	})

	if canceled {
		cui.Warn("Canceled")
		return
	}

	for _, a := range answers {
		switch a {
		case "message":
			cui.Message("This is test message!!", []string{color.BlueFg, color.Bold})
		case "debug":
			cui.Debug("This is debug message!!")
		case "info":
			cui.Info("This is info message!!")
		case "notice":
			cui.Notice("This is notice message!!")
		case "warn":
			cui.Warn("This is warning message!!")
		case "error":
			cui.Error("This is error message!!")
		case "input":
			answer, canceled := cui.Input("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case "secureInput":
			answer, canceled := cui.SecureInput("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case "confirmation":
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
		case "list":
			answer, canceled := cui.List("Which operation system do you like?", []string{"Windows", "Linux", "MacOS"})
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case "progressBar":
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
		case "spinner":
			spnr0 := cui.Spinner("Waiting some operations... (success case)")
			time.Sleep(3 * time.Second)
			spnr0.Complete()

			spnr1 := cui.Spinner("Waiting some operations... (failure case)")
			time.Sleep(3 * time.Second)
			spnr1.Failure()
		}
	}
}
