package input

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

type Input struct {
	elemChain    elementChain.ElementChain
	inputPointer int
	canceled     bool
	finished     bool
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(question string) *Input {
	/*
	 * input is following format.
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

	input := Input{
		elemChain: *elementChain.New(elems),
	}
	return &input
}

/*
 * Private Methods
 */

func (input *Input) appendRunes(r []rune) {
	oldAnswer := input.getAnswerElemPtr().Str

	oldAnswerFront := string([]rune(oldAnswer)[:input.inputPointer])
	oldAnswerRear := string([]rune(oldAnswer)[input.inputPointer:])

	newAnswer := oldAnswerFront + string(r) + oldAnswerRear
	input.getAnswerElemPtr().Str = newAnswer
	input.inputPointer += len(r)
}

func (input Input) getPrefixElemPtr() *element.Element {
	return &input.elemChain.Elems[0]
}

func (input Input) getPaddingElemPtr() *element.Element {
	return &input.elemChain.Elems[1]
}

func (input Input) getQuestionElemPtr() *element.Element {
	return &input.elemChain.Elems[2]
}

func (input Input) getQuestionSuffixElemPtr() *element.Element {
	return &input.elemChain.Elems[3]
}

func (input Input) getAnswerElemPtr() *element.Element {
	return &input.elemChain.Elems[4]
}

func (input Input) getNextLineElemPtr() *element.Element {
	return &input.elemChain.Elems[5]
}

func (input Input) getAnswerStartX() (answerStartX int) {
	answerStartX += input.GetStartX()
	answerStartX += len(input.getPrefixElemPtr().Str)
	answerStartX += len(input.getPaddingElemPtr().Str)
	answerStartX += len(input.getQuestionElemPtr().Str)
	answerStartX += len(input.getQuestionSuffixElemPtr().Str)
	return
}

func (input Input) getAnswerEndX() (answerEndX int) {
	answerEndX += input.getAnswerStartX()
	answerEndX += len(input.getAnswerElemPtr().Str)
	return
}

func (input *Input) finalizeAnswer() {
	input.getAnswerElemPtr().Colors = constants.AnswerColors
	input.getNextLineElemPtr().Str = constants.NewLine
}

func (input *Input) print() {
	input.elemChain.Print()
}

func (input *Input) removeRune() {
	oldAnswer := input.getAnswerElemPtr().Str

	if len(oldAnswer) == 0 {
		return
	}

	newAnswer := string([]rune(oldAnswer)[:input.inputPointer-1])
	if input.inputPointer < len(oldAnswer)-1 {
		newAnswer += string([]rune(oldAnswer)[input.inputPointer:])
	}
	input.getAnswerElemPtr().Str = newAnswer
	input.inputPointer--
}

/*
 * Public Methods
 */

func (input Input) Erase() {
	input.elemChain.Erase()
}

func (input Input) GetMinX() int {
	return input.elemChain.GetMinX()
}

func (input Input) GetMinY() int {
	return input.elemChain.GetMinY()
}

func (input Input) GetMaxX() int {
	return input.elemChain.GetMaxX()
}

func (input Input) GetMaxY() int {
	return input.elemChain.GetMaxY()
}

func (input Input) GetStartX() int {
	return input.elemChain.GetStartX()
}

func (input Input) GetStartY() int {
	return input.elemChain.GetStartY()
}

func (input Input) GetEndX() int {
	return input.elemChain.GetEndX()
}

func (input Input) GetEndY() int {
	return input.elemChain.GetEndY()
}

func (input *Input) Ask() (string, bool) {
	input.print()
	for {
		cursor.MoveCursorTo(input.GetStartX(), input.GetStartY())
		input.print()
		if input.finished || input.canceled {
			break
		}
		cursor.MoveCursorTo(input.getAnswerStartX() + input.inputPointer, input.GetStartY())

		// Get keyboard input
		inputHelper.SetRaw(true)
		inputHelper.SetNoEcho(true)
		inputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		inputHelper.SetRaw(false)

		input.Erase()

		switch string(inputRunes) {
		case constants.Delete: // delete
			input.removeRune()
		case constants.Enter: // enter
			input.finished = true
			input.finalizeAnswer()
		case constants.CtrlC: // ctrl + c
			input.canceled = true
			input.finalizeAnswer()
		case constants.RightArrow: // right arrow
			if len(input.getAnswerElemPtr().Str) > input.inputPointer {
				input.inputPointer++
			}
		case constants.LeftArrow: // left arrow
			if 0 < input.inputPointer {
				input.inputPointer--
			}
		default:
			input.appendRunes(inputRunes)
		}
	}
	return input.getAnswerElemPtr().Str, input.canceled
}
