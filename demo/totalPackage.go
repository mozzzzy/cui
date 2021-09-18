package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
	"github.com/mozzzzy/cui/v3/color"
)

func main() {
	choices := []string{
		"message",
		"debug",
		"info",
		"notice",
		"warn",
		"error",
		"table",
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
		cui.Erase()
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
			cui.Table([][]string{
				{"column0", "column1", "column2", "column3"},
				{"data00", "data01", "data02", "data03"},
				{"", "data11", "data12", "data13"},
				{"data20", "", "data22", "data23"},
				{"data30", "data31", "", "data33"},
				{"data40", "data41", "data42", ""},
			})
		case 7:
			answer, canceled := cui.Input("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case 8:
			answer, canceled := cui.SecureInput("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + answer + "\"")
		case 9:
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
		case 10:
			osNames := []string{"Windows", "Linux", "MacOS"}
			answer, canceled := cui.List("Which operation system do you like?", osNames)
			if canceled {
				cui.Warn("Canceled")
				return
			}
			cui.Info("Answer: \"" + osNames[answer] + "\"")
		case 11:
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
		case 12:
			spnr0 := cui.Spinner("Waiting some operations... (success case)")
			time.Sleep(3 * time.Second)
			spnr0.Complete()

			spnr1 := cui.Spinner("Waiting some operations... (failure case)")
			time.Sleep(3 * time.Second)
			spnr1.Failure()
		}
	}

	cui.Info("Sleep 3 seconds and erase all output.")
	spnrSleep := cui.Spinner("Sleeping...")
	time.Sleep(3 * time.Second)
	spnrSleep.Complete()
	cui.Erase()
}
