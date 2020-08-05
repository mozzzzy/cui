package input

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

type Input struct {
	elemChain elementChain.ElementChain
	canceled  bool
	finished  bool
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
	xCursor, _ := cursor.GetCursor()
	position := input.getAppendPosition(xCursor)
	oldAnswer := input.getAnswerElem().Str
	newAnswer :=
		string([]rune(oldAnswer)[:position]) +
			string(r) +
			string([]rune(oldAnswer)[position:])
	input.getAnswerElem().Str = newAnswer
}

func (input Input) getAnswerElem() *element.Element {
	return &input.elemChain.Elems[4]
}

func (input Input) getAnswerStart() (answerStart int) {
	for i := 0; i < 4; i++ {
		answerStart += len(input.elemChain.Elems[i].Str)
	}
	return
}

func (input Input) getAnswerEnd() (answerEnd int) {
	for i := 0; i < 5; i++ {
		answerEnd += len(input.elemChain.Elems[i].Str)
	}
	return
}

func (input Input) getAppendPosition(xCursor int) (position int) {
	answerStart := input.getAnswerStart()
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

func (input *Input) fixAnswer() {
	input.getAnswerElem().Colors = constants.AnswerColors
	input.elemChain.Elems[len(input.elemChain.Elems)-1].Str = constants.NewLine
}

func (input *Input) removeRune() {
	xCursor, _ := cursor.GetCursor()
	position := input.getAppendPosition(xCursor)
	oldAnswer := input.getAnswerElem().Str

	if len(oldAnswer) == 0 {
		return
	}

	newAnswer := string([]rune(oldAnswer)[:position-1])
	if position < len(oldAnswer)-1 {
		newAnswer += string([]rune(oldAnswer)[position:])
	}
	input.getAnswerElem().Str = newAnswer
}

/*
 * Public Methods
 */

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
	input.Print()
	xCursorShelter, yCursorShelter := cursor.GetCursor()
	inputHelper.SetRaw(true)
	for {
		input.elemChain.Erase()
		cursor.MoveCursorTo(input.GetStartX(), input.GetStartY())
		input.Print()
		if input.finished || input.canceled {
			break
		}
		cursor.MoveCursorTo(xCursorShelter, yCursorShelter)
		// Get keyboard input
		inputHelper.SetNoEcho(true)
		inputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		switch string(inputRunes) {
		case constants.Delete: // delete
			input.removeRune()
			answerStartX := input.getAnswerStart()
			if answerStartX < xCursorShelter {
				xCursorShelter--
			}
		case constants.Enter: // enter
			input.finished = true
			input.fixAnswer()
		case constants.CtrlC: // ctrl + c
			input.canceled = true
			input.fixAnswer()
		case constants.RightArrow: // right arrow
			answerEndX := input.getAnswerEnd()
			if answerEndX > xCursorShelter {
				xCursorShelter++
			}
		case constants.LeftArrow: // left arrow
			answerStartX := input.getAnswerStart()
			if answerStartX < xCursorShelter {
				xCursorShelter--
			}
		default:
			input.appendRunes(inputRunes)
			answerEndX := input.getAnswerEnd()
			if answerEndX > xCursorShelter {
				xCursorShelter++
			}
		}
	}
	inputHelper.SetRaw(false)
	return input.getAnswerElem().Str, input.canceled
}

func (input *Input) Print() {
	input.elemChain.Print()
}
