package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
	"github.com/mozzzzy/cui/v3/color"
)

func main() {
	q := cui.Question("Which demonstration do you run?")

	choices := []string{
		"message",
		"debug",
		"info",
		"notice",
		"warn",
		"error",
		"question",
		"table",
		"checkableTable",
		"input",
		"secureInput",
		"confirmation",
		"list",
		"progressBar",
		"spinner",
	}
	answers, canceled, checkbox := cui.Checkbox(choices)
	checkbox.Erase()

	if canceled {
		cui.Warn("Canceled")
		cui.Erase()
		return
	}

	answerStr := ""
	for index, a := range answers {
		answerStr += choices[a]
		if index < len(answers)-1 {
			answerStr += ","
		}
	}
	q.SetAnswer(answerStr)

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
			qDemo0 := cui.Question("What is your favorite language?")
			time.Sleep(1 * time.Second)
			qDemo0.SetAnswer("go")
		case 7:
			cui.Table([][]string{
				{"column0", "column1", "column2", "column3"},
				{"data00", "data01", "data02", "data03"},
				{"", "data11", "data12", "data13"},
				{"data20", "", "data22", "data23"},
				{"data30", "data31", "", "data33"},
				{"data40", "data41", "data42", ""},
			})
		case 8:
			qDemo1 := cui.Question("Check some rows...")
			rows := [][]string{
				{"column0", "column1", "column2", "column3"},
				{"data00", "data01", "data02", "data03"},
				{"data10", "data11", "data12", "data13"},
				{"data20", "data21", "data22", "data23"},
				{"data30", "data31", "data32", "data33"},
			}
			answers, canceled, chkTbl := cui.CheckableTable(rows)
			chkTbl.Erase()
			if canceled {
				cui.Warn("Canceled")
			} else {
				answerStr := ""
				for index, a := range answers {
					answerStr += rows[a][0]
					if index < len(answers)-1 {
						answerStr += ","
					}
				}
				qDemo1.SetAnswer(answerStr)
			}
		case 9:
			answer, canceled := cui.Input("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
			} else {
				cui.Info("Answer: \"" + answer + "\"")
			}
		case 10:
			answer, canceled := cui.SecureInput("Please type something and press Enter")
			if canceled {
				cui.Warn("Canceled")
			} else {
				cui.Info("Answer: \"" + answer + "\"")
			}
		case 11:
			answer, canceled := cui.Confirmation("Please type Y or n and press Enter")
			if canceled {
				cui.Warn("Canceled")
			} else {
				if answer {
					cui.Info("Accepted.")
				} else {
					cui.Warn("Rejected.")
				}
			}
		case 12:
			qDemo2 := cui.Question("Which operation system do you like?")

			osNames := []string{"Windows", "Linux", "MacOS"}
			answer, canceled, list := cui.List(osNames)
			list.Erase()
			if canceled {
				cui.Warn("Canceled")
			} else {
				qDemo2.SetAnswer(osNames[answer])
			}
		case 13:
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
		case 14:
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
