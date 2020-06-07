package cui

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v2/checkbox"
	"github.com/mozzzzy/cui/v2/confirmation"
	"github.com/mozzzzy/cui/v2/debugMessage"
	"github.com/mozzzzy/cui/v2/errorMessage"
	"github.com/mozzzzy/cui/v2/infoMessage"
	"github.com/mozzzzy/cui/v2/input"
	"github.com/mozzzzy/cui/v2/list"
	"github.com/mozzzzy/cui/v2/message"
	"github.com/mozzzzy/cui/v2/noticeMessage"
	"github.com/mozzzzy/cui/v2/prefixedMessage"
	"github.com/mozzzzy/cui/v2/progressBar"
	"github.com/mozzzzy/cui/v2/secureInput"
	"github.com/mozzzzy/cui/v2/spinner"
	"github.com/mozzzzy/cui/v2/table"
	"github.com/mozzzzy/cui/v2/warnMessage"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func Message(msg string, colors []string) {
	message.New(msg, colors).Print()
}

func PrefixedMessage(
	prefix string, prefixColors []string,
	padding string, paddingColors []string,
	msg string, colors []string) {
	prefixedMessage.New(prefix, prefixColors, padding, paddingColors, msg, colors).Print()
}

func Debug(msg string) {
	debugMessage.New(msg).Print()
}

func Info(msg string) {
	infoMessage.New(msg).Print()
}

func Notice(msg string) {
	noticeMessage.New(msg).Print()
}

func Warn(msg string) {
	warnMessage.New(msg).Print()
}

func Error(msg string) {
	errorMessage.New(msg).Print()
}

func Table(data [][]string) {
	table.New(data).Print()
}

func Spinner(msg string) *spinner.Spinner {
	spnr := spinner.New(msg)
	spnr.Run()
	return spnr
}

func ProgressBar(msg string) *progressBar.ProgressBar {
	pb := progressBar.New(msg)
	pb.Print()
	return pb
}

func List(question string, choices []string) (int, bool) {
	return list.New(question, choices).Ask()
}

func Checkbox(question string, choices []string) ([]int, bool) {
	return checkbox.New(question, choices).Ask()
}

func Confirmation(question string) (bool, bool) {
	return confirmation.New(question).Ask()
}

func Input(question string) (string, bool) {
	return input.New(question).Ask()
}

func SecureInput(question string) (string, bool) {
	return secureInput.New(question).Ask()
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */
