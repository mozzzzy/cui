package secureInput

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

type SecureInput struct {
	elemChain elementChain.ElementChain
	inputPointer int
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
	oldAnswer := secureInput.answer
	oldAnswerFront := string([]rune(oldAnswer)[:secureInput.inputPointer])
	oldAnswerRear := string([]rune(oldAnswer))[secureInput.inputPointer:]
	newAnswer := oldAnswerFront + string(r) + oldAnswerRear
	secureInput.answer = newAnswer

	oldAnswerFake := secureInput.getAnswerElemPtr().Str
	oldAnswerFakeFront := string([]rune(oldAnswerFake)[:secureInput.inputPointer])
	oldAnswerFakeRear := string([]rune(oldAnswerFake)[secureInput.inputPointer:])
	newAnswerFake := oldAnswerFakeFront + string(AnswerFake) + oldAnswerFakeRear
	secureInput.getAnswerElemPtr().Str = newAnswerFake

	secureInput.inputPointer += len(r)
}

func (secureInput SecureInput) getPrefixElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[0]
}

func (secureInput SecureInput) getPaddingElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[1]
}

func (secureInput SecureInput) getQuestionElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[2]
}

func (secureInput SecureInput) getQuestionSuffixElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[3]
}

func (secureInput SecureInput) getAnswerElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[4]
}

func (secureInput SecureInput) getNextLineElemPtr() *element.Element {
	return &secureInput.elemChain.Elems[5]
}

func (secureInput SecureInput) getAnswerStartX() (answerStartX int) {
	answerStartX += secureInput.GetStartX()
	answerStartX += len(secureInput.getPrefixElemPtr().Str)
	answerStartX += len(secureInput.getPaddingElemPtr().Str)
	answerStartX += len(secureInput.getQuestionElemPtr().Str)
	answerStartX += len(secureInput.getQuestionSuffixElemPtr().Str)
	return
}

func (secureInput SecureInput) getAnswerEndX() (answerEndX int) {
	answerEndX += secureInput.getAnswerStartX()
	answerEndX += len(secureInput.getAnswerElemPtr().Str)
	return
}

func (secureInput *SecureInput) finalizeAnswer() {
	secureInput.getAnswerElemPtr().Colors = constants.AnswerColors
	secureInput.getNextLineElemPtr().Str = constants.NewLine
}

func (secureInput *SecureInput) print() {
	secureInput.elemChain.Print()
}

func (secureInput *SecureInput) removeRune() {
	oldAnswerFake := secureInput.getAnswerElemPtr().Str
	oldAnswer := secureInput.answer

	if len(oldAnswerFake) == 0 {
		return
	}

	newAnswerFake := string([]rune(oldAnswerFake)[:secureInput.inputPointer-1])
	newAnswer := string([]rune(oldAnswer)[:secureInput.inputPointer-1])
	if secureInput.inputPointer < len(oldAnswer)-1 {
		newAnswerFake += string([]rune(oldAnswerFake)[secureInput.inputPointer:])
		newAnswer += string([]rune(oldAnswer)[secureInput.inputPointer:])
	}
	secureInput.getAnswerElemPtr().Str = newAnswerFake
	secureInput.answer = newAnswer
	secureInput.inputPointer--
}

/*
 * Public Methods
 */

func (secureInput SecureInput) Erase() {
	secureInput.elemChain.Erase()
}

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
	secureInput.print()
	for {
		cursor.MoveCursorTo(secureInput.GetStartX(), secureInput.GetStartY())
		secureInput.print()
		if secureInput.finished || secureInput.canceled {
			break
		}
		cursor.MoveCursorTo(
			secureInput.getAnswerStartX() + secureInput.inputPointer,
			secureInput.GetStartY())

		// Get keyboard secureInput
		inputHelper.SetRaw(true)
		inputHelper.SetNoEcho(true)
		secureInputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		inputHelper.SetRaw(false)

		secureInput.Erase()

		switch string(secureInputRunes) {
		case constants.Delete: // delete
			secureInput.removeRune()
		case constants.Enter: // enter
			secureInput.finished = true
			secureInput.finalizeAnswer()
		case constants.CtrlC: // ctrl + c
			secureInput.canceled = true
			secureInput.finalizeAnswer()
		case constants.RightArrow: // right arrow
			if len(secureInput.getAnswerElemPtr().Str) > secureInput.inputPointer {
				secureInput.inputPointer++
			}
		case constants.LeftArrow: // left arrow
			if 0 < secureInput.inputPointer {
				secureInput.inputPointer--
			}
		default:
			secureInput.appendRunes(secureInputRunes)
		}
	}
	return secureInput.answer, secureInput.canceled
}
