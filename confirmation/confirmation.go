package confirmation

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/constants"
	"github.com/mozzzzy/cui/cursor"
	"github.com/mozzzzy/cui/element"
	"github.com/mozzzzy/cui/elementChain"
	"github.com/mozzzzy/cui/inputHelper"
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

var (
	QuestionSuffix  string = ": "
	AnswerYes       string = "Y"
	AnswerNo        string = "n"
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

func (confirm *Confirmation) setAnswerElem() {
	if confirm.finished {
		// Set Answer
		if confirm.answer {
			confirm.elemChain.Elems[4].Str = AnswerYes
		} else {
			confirm.elemChain.Elems[4].Str = AnswerNo
		}
		confirm.elemChain.Elems[4].Colors = constants.AnswerColors

		// Set Next Line
		confirm.elemChain.Elems[5].Str = constants.NewLine

		// Unset temporary answer
		confirm.unsetAnswerElem()
		return
	}

	if confirm.canceled {
		// Set Next Line
		confirm.elemChain.Elems[5].Str = constants.NewLine
		// Unset temporary answer
		confirm.unsetAnswerElem()
		return
	}

	if confirm.answer {
		confirm.elemChain.Elems[6].Str = AnswerYes
	} else {
		confirm.elemChain.Elems[6].Str = AnswerNo
	}
}

func (confirm *Confirmation) unsetAnswerElem() {
	confirm.elemChain.Elems[6].Str = ""
}

/*
 * Public Methods
 */

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
	confirm.Print()
	inputHelper.SetRaw(true)
	for {
		confirm.elemChain.Erase()
		cursor.MoveCursorTo(confirm.GetStartX(), confirm.GetStartY())
		confirm.Print()
		if confirm.finished || confirm.canceled {
			break
		}
		inputRunes := inputHelper.Getch()
		switch string(inputRunes) {
		case AnswerYes:
			confirm.answer = true
			confirm.setAnswerElem()
		case AnswerNo:
			confirm.answer = false
			confirm.setAnswerElem()
		case constants.Delete: // delete
			confirm.unsetAnswerElem()
		case constants.Enter: // enter
			confirm.finished = true
			confirm.setAnswerElem()
		case constants.CtrlC: // ctrl + c
			confirm.canceled = true
			confirm.setAnswerElem()
		}
	}
	inputHelper.SetRaw(false)
	return confirm.answer, confirm.canceled
}

func (confirm *Confirmation) Print() {
	confirm.elemChain.Print()
}
