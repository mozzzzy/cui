package cui

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/checkableTable"
	"github.com/mozzzzy/cui/v3/checkbox"
	"github.com/mozzzzy/cui/v3/confirmation"
	"github.com/mozzzzy/cui/v3/core/cursor"
	"github.com/mozzzzy/cui/v3/debugMessage"
	"github.com/mozzzzy/cui/v3/errorMessage"
	"github.com/mozzzzy/cui/v3/infoMessage"
	"github.com/mozzzzy/cui/v3/input"
	"github.com/mozzzzy/cui/v3/list"
	"github.com/mozzzzy/cui/v3/message"
	"github.com/mozzzzy/cui/v3/noticeMessage"
	"github.com/mozzzzy/cui/v3/progressBar"
	"github.com/mozzzzy/cui/v3/secureInput"
	"github.com/mozzzzy/cui/v3/spinner"
	"github.com/mozzzzy/cui/v3/question"
	"github.com/mozzzzy/cui/v3/table"
	"github.com/mozzzzy/cui/v3/warnMessage"
)

/*
 * Types
 */

type Erasable interface {
	Erase()
}

/*
 * Constants and Package Scope Variables
 */

var erasables []Erasable = []Erasable{}

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func Message(msg string, colors []string) {
	e := message.New(msg, colors)
	e.Print()
	erasables = append(erasables, e)
}

func Debug(msg string) {
	e := debugMessage.New(msg)
	e.Print()
	erasables = append(erasables, e)
}

func Info(msg string) {
	e := infoMessage.New(msg)
	e.Print()
	erasables = append(erasables, e)
}

func Notice(msg string) {
	e := noticeMessage.New(msg)
	e.Print()
	erasables = append(erasables, e)
}

func Warn(msg string) {
	e := warnMessage.New(msg)
	e.Print()
	erasables = append(erasables, e)
}

func Error(msg string) {
	e := errorMessage.New(msg)
	e.Print()
	erasables = append(erasables, e)
}

func Table(data [][]string) {
	e := table.New(data)
	e.Print()
	erasables = append(erasables, e)
}

func Spinner(msg string) *spinner.Spinner {
	e := spinner.New(msg)
	e.Run()
	erasables = append(erasables, e)
	return e
}

func ProgressBar(msg string) *progressBar.ProgressBar {
	e := progressBar.New(msg)
	e.Print()
	erasables = append(erasables, e)
	return e
}

func List(choices []string) (int, bool, *list.List) {
	e := list.New(choices)
	answer, canceled, _ := e.Ask()
	erasables = append(erasables, e)
	return answer, canceled, e
}

func CheckableTable(choices [][]string) ([]int, bool, *checkableTable.CheckableTable) {
	e := checkableTable.New(choices)
	answers, canceled, _ := e.Ask()
	erasables = append(erasables, e)
	return answers, canceled, e
}

func Checkbox(choices []string) ([]int, bool, *checkbox.Checkbox) {
	e := checkbox.New(choices)
	answers, canceled, _ := e.Ask()
	erasables = append(erasables, e)
	return answers, canceled, e
}

func Confirmation(question string) (bool, bool) {
	e := confirmation.New(question)
	answer, canceled := e.Ask()
	erasables = append(erasables, e)
	return answer, canceled
}

func Input(question string) (string, bool) {
	e := input.New(question)
	answer, canceled := e.Ask()
	erasables = append(erasables, e)
	return answer, canceled
}

func SecureInput(question string) (string, bool) {
	e := secureInput.New(question)
	answer, canceled := e.Ask()
	erasables = append(erasables, e)
	return answer, canceled
}

func Question(q string) *question.Question {
	e := question.New(q)
	e.Print()
	erasables = append(erasables, e)
	return e
}

func Erase() {
	for _, e := range erasables {
		e.Erase()
	}
	cursor.MoveCursorToZeroZero()
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */
