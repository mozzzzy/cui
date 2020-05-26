package cui

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/checkbox"
	"github.com/mozzzzy/cui/confirmation"
	"github.com/mozzzzy/cui/debugMessage"
	"github.com/mozzzzy/cui/errorMessage"
	"github.com/mozzzzy/cui/infoMessage"
	"github.com/mozzzzy/cui/input"
	"github.com/mozzzzy/cui/list"
	"github.com/mozzzzy/cui/message"
	"github.com/mozzzzy/cui/noticeMessage"
	"github.com/mozzzzy/cui/prefixedMessage"
	"github.com/mozzzzy/cui/progressBar"
	"github.com/mozzzzy/cui/secureInput"
	"github.com/mozzzzy/cui/spinner"
	"github.com/mozzzzy/cui/warnMessage"
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

func List(question string, choices []string) (string, bool) {
	return list.New(question, choices).Ask()
}

func Checkbox(question string, choices []string) ([]string, bool) {
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
