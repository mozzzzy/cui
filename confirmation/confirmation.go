package confirmation

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/cursor"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
	"github.com/mozzzzy/cui/v3/core/inputHelper"
)

/*
 * Types
 */

type Confirmation struct {
	elemChain elementChain.ElementChain
	answer    bool
	canceled  bool
	finished  bool
}

/*
 * Constants and Package Scope Variables
 */

const (
	QuestionSuffix  string = ": "
	AnswerYes       string = "y"
	AnswerNo        string = "N"
	AnswerSeparator string = "/"
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(question string) *Confirmation {
	/*
	 * confirmation is following format.
	 * +--------+---------+----------+---------+--------+--------------+-----------------+
	 * | Prefix | Padding | Question | Padding | Answer | \r\n or ": " | Temporary input |
	 * +--------+---------+----------+---------+--------+--------------+-----------------+
	 */

	elems := []element.Element{
		// Prefix
		{
			Str:    constants.QuestionPrefix,
			Colors: constants.QuestionPrefixColors,
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// Question
		{
			Str:    question,
			Colors: constants.QuestionColors,
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// Answer
		{
			Str: constants.OpenParenthesis +
				AnswerYes +
				AnswerSeparator +
				AnswerNo +
				constants.CloseParenthesis,
			Colors: constants.QuestionColors,
		},
		// Question suffix
		{
			Str:    QuestionSuffix,
			Colors: constants.QuestionColors,
		},
		// Temporary input
		{
			Str:    "",
			Colors: constants.AnswerColors,
		},
	}

	confirm := Confirmation{
		elemChain: *elementChain.New(elems),
	}
	return &confirm
}

/*
 * Private Methods
 */

func (confirm Confirmation) getAnswerElemPtr() *element.Element {
	return &(confirm.elemChain.Elems[4])
}

func (confirm Confirmation) getSuffixElemPtr() *element.Element {
	return &(confirm.elemChain.Elems[5])
}

func (confirm Confirmation) getTmpInputElemPtr() *element.Element {
	return &(confirm.elemChain.Elems[6])
}

func (confirm *Confirmation) print() {
	confirm.elemChain.Print()
}

func (confirm *Confirmation) setAnswerElem() {
	if confirm.finished {
		// Set Answer
		if confirm.answer {
			confirm.getAnswerElemPtr().Str = AnswerYes
		} else {
			confirm.getAnswerElemPtr().Str = AnswerNo
		}
		confirm.getAnswerElemPtr().Colors = constants.AnswerColors

		// Set Next Line
		confirm.getSuffixElemPtr().Str = constants.NewLine

		// Unset temporary answer
		confirm.unsetTmpInputElem()
		return
	}

	if confirm.canceled {
		// Set Next Line
		confirm.getSuffixElemPtr().Str = constants.NewLine
		// Unset temporary answer
		confirm.unsetTmpInputElem()
		return
	}

	if confirm.answer {
		confirm.getTmpInputElemPtr().Str = AnswerYes
	} else {
		confirm.getTmpInputElemPtr().Str = AnswerNo
	}
}

func (confirm *Confirmation) unsetTmpInputElem() {
	confirm.getTmpInputElemPtr().Str = ""
}

/*
 * Public Methods
 */

func (confirm Confirmation) Erase() {
	confirm.elemChain.Erase()
}

func (confirm Confirmation) GetMinX() int {
	return confirm.elemChain.GetMinX()
}

func (confirm Confirmation) GetMinY() int {
	return confirm.elemChain.GetMinY()
}

func (confirm Confirmation) GetMaxX() int {
	return confirm.elemChain.GetMaxX()
}

func (confirm Confirmation) GetMaxY() int {
	return confirm.elemChain.GetMaxY()
}

func (confirm Confirmation) GetStartX() int {
	return confirm.elemChain.GetStartX()
}

func (confirm Confirmation) GetStartY() int {
	return confirm.elemChain.GetStartY()
}

func (confirm Confirmation) GetEndX() int {
	return confirm.elemChain.GetEndX()
}

func (confirm Confirmation) GetEndY() int {
	return confirm.elemChain.GetEndY()
}

func (confirm *Confirmation) Ask() (bool, bool) {
	confirm.print()
	for {
		cursor.MoveCursorTo(confirm.GetStartX(), confirm.GetStartY())
		confirm.print()
		if confirm.finished || confirm.canceled {
			break
		}

		inputHelper.SetRaw(true)
		inputHelper.SetNoEcho(true)
		inputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		inputHelper.SetRaw(false)

		confirm.Erase()

		switch string(inputRunes) {
		case AnswerYes:
			confirm.answer = true
			confirm.setAnswerElem()
		case AnswerNo:
			confirm.answer = false
			confirm.setAnswerElem()
		case constants.Delete: // delete
			confirm.unsetTmpInputElem()
		case constants.Enter: // enter
			confirm.finished = true
			confirm.setAnswerElem()
		case constants.CtrlC: // ctrl + c
			confirm.canceled = true
			confirm.setAnswerElem()
		}
	}
	return confirm.answer, confirm.canceled
}
