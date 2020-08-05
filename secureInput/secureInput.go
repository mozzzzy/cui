package secureInput

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v2/constants"
	"github.com/mozzzzy/cui/v2/cursor"
	"github.com/mozzzzy/cui/v2/element"
	"github.com/mozzzzy/cui/v2/elementChain"
	"github.com/mozzzzy/cui/v2/inputHelper"
)

/*
 * Types
 */

type SecureInput struct {
	elemChain elementChain.ElementChain
	answer    string
	canceled  bool
	finished  bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	AnswerFake rune = '*'
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(question string) *SecureInput {
	/*
	 * secureInput is following format.
	 * +--------+---------+----------+----------------+--------+------+
	 * | Prefix | Padding | Question | QuestionSuffix | Answer | \r\n |
	 * +--------+---------+----------+----------------+--------+------+
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
		// Question suffix
		{
			Str:    constants.QuestionSuffix,
			Colors: constants.QuestionColors,
		},
		// Answer
		{
			Str:    "",
			Colors: constants.QuestionColors,
		},
		// Next line
		{
			Str:    "",
			Colors: []string{},
		},
	}

	secureInput := SecureInput{
		elemChain: *elementChain.New(elems),
	}
	return &secureInput
}

/*
 * Private Methods
 */

func (secureInput *SecureInput) appendRunes(r []rune) {
	xCursor, _ := cursor.GetCursor()
	position := secureInput.getAppendPosition(xCursor)
	oldAnswer := secureInput.answer
	newAnswer :=
		string([]rune(oldAnswer)[:position]) +
			string(r) +
			string([]rune(oldAnswer)[position:])
	oldAnswerFake := secureInput.getAnswerElem().Str
	newAnswerFake :=
		string([]rune(oldAnswerFake)[:position]) +
			string(AnswerFake) +
			string([]rune(oldAnswerFake)[position:])
	secureInput.answer = newAnswer
	secureInput.getAnswerElem().Str = newAnswerFake
}

func (secureInput SecureInput) getAnswerElem() *element.Element {
	return &secureInput.elemChain.Elems[4]
}

func (secureInput SecureInput) getAnswerStart() (answerStart int) {
	for i := 0; i < 4; i++ {
		answerStart += len(secureInput.elemChain.Elems[i].Str)
	}
	return
}

func (secureInput SecureInput) getAnswerEnd() (answerEnd int) {
	for i := 0; i < 5; i++ {
		answerEnd += len(secureInput.elemChain.Elems[i].Str)
	}
	return
}

func (secureInput SecureInput) getAppendPosition(xCursor int) (position int) {
	answerStart := secureInput.getAnswerStart()
	/*
	 * startX                                            cursorX
	 * |                                                 |
	 * v                                                 v
	 * +--------+---------+----------+----------------+--------+------+
	 * | Prefix | Padding | Question | QuestionSuffix | Answer | \r\n |
	 * +--------+---------+----------+----------------+--------+------+
	 *                                                ^
	 *                                                |
	 *                                                answerStart
	 */
	position = xCursor - answerStart
	return
}

func (secureInput *SecureInput) fixAnswer() {
	secureInput.getAnswerElem().Colors = constants.AnswerColors
	secureInput.elemChain.Elems[len(secureInput.elemChain.Elems)-1].Str = constants.NewLine
}

func (secureInput *SecureInput) removeRune() {
	xCursor, _ := cursor.GetCursor()
	position := secureInput.getAppendPosition(xCursor)
	oldAnswerFake := secureInput.getAnswerElem().Str
	oldAnswer := secureInput.answer
	if len(oldAnswerFake) == 0 {
		return
	}

	newAnswerFake := string([]rune(oldAnswerFake)[:position-1])
	newAnswer := string([]rune(oldAnswer)[:position-1])
	if position < len(oldAnswer)-1 {
		newAnswerFake += string([]rune(oldAnswerFake)[position:])
		newAnswer += string([]rune(oldAnswer)[position:])
	}
	secureInput.getAnswerElem().Str = newAnswerFake
	secureInput.answer = newAnswer
}

/*
 * Public Methods
 */

func (secureInput SecureInput) GetMinX() int {
	return secureInput.elemChain.GetMinX()
}

func (secureInput SecureInput) GetMinY() int {
	return secureInput.elemChain.GetMinY()
}

func (secureInput SecureInput) GetMaxX() int {
	return secureInput.elemChain.GetMaxX()
}

func (secureInput SecureInput) GetMaxY() int {
	return secureInput.elemChain.GetMaxY()
}

func (secureInput SecureInput) GetStartX() int {
	return secureInput.elemChain.GetStartX()
}

func (secureInput SecureInput) GetStartY() int {
	return secureInput.elemChain.GetStartY()
}

func (secureInput SecureInput) GetEndX() int {
	return secureInput.elemChain.GetEndX()
}

func (secureInput SecureInput) GetEndY() int {
	return secureInput.elemChain.GetEndY()
}

func (secureInput *SecureInput) Ask() (string, bool) {
	secureInput.Print()
	xCursorShelter, yCursorShelter := cursor.GetCursor()

	inputHelper.SetRaw(true)
	for {
		secureInput.elemChain.Erase()
		cursor.MoveCursorTo(secureInput.GetStartX(), secureInput.GetStartY())
		secureInput.Print()
		if secureInput.finished || secureInput.canceled {
			break
		}
		cursor.MoveCursorTo(xCursorShelter, yCursorShelter)
		// Get keyboard secureInput
		inputHelper.SetNoEcho(true)
		secureInputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		switch string(secureInputRunes) {
		case constants.Delete: // delete
			secureInput.removeRune()
			answerStartX := secureInput.getAnswerStart()
			if answerStartX < xCursorShelter {
				xCursorShelter--
			}
		case constants.Enter: // enter
			secureInput.finished = true
			secureInput.fixAnswer()
		case constants.CtrlC: // ctrl + c
			secureInput.canceled = true
			secureInput.fixAnswer()
		case constants.RightArrow: // right arrow
			answerEndX := secureInput.getAnswerEnd()
			if answerEndX > xCursorShelter {
				xCursorShelter++
			}
		case constants.LeftArrow: // left arrow
			answerStartX := secureInput.getAnswerStart()
			if answerStartX < xCursorShelter {
				xCursorShelter--
			}
		default:
			secureInput.appendRunes(secureInputRunes)
			answerEndX := secureInput.getAnswerEnd()
			if answerEndX > xCursorShelter {
				xCursorShelter++
			}
		}
	}
	inputHelper.SetRaw(false)
	return secureInput.answer, secureInput.canceled
}

func (secureInput *SecureInput) Print() {
	secureInput.elemChain.Print()
}
