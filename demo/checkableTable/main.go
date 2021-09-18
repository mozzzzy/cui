package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of checkable package.")

	q := cui.Question("Check what you want to buy")

	choices := [][]string{
		{"Product", "Price"},
		{"iMac", "$1299"},
		{"iPad", "$329"},
		{"iPhone", "$41.62"},
	}
	answers, canceled, chkTbl := cui.CheckableTable(choices)
	chkTbl.Erase()

	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
		return
	}

	answerStr := ""
	for index, a := range answers {
		answerStr += choices[a][0]
		if index < len(answers)-1 {
			answerStr += ","
		}
	}
	q.SetAnswer(answerStr)

	time.Sleep(3 * time.Second)
	cui.Erase()
	return
}
