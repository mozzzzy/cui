package checkbox

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/cursor"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
	"github.com/mozzzzy/cui/v3/core/inputHelper"
)

/*
 * Types
 */

type Checkbox struct {
	firstLineElemChain elementChain.ElementChain
	choicesElemChain   elementChain.ElementChain
	choices            []string
	chosePositions     []int
	pointerPosition    int
	finished           bool
	canceled           bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	ChosePrefix       string   = "[x]"
	ChosePrefixColors []string = []string{color.CyanFg}
	ChoseColors       []string = []string{}

	NoChosePrefix       string = "[ ]"
	NoChosePrefixColors []string
	NoChoseColors       []string = []string{}

	AnswerNone      string = "Nothing was selected"
	AnswerSeparator string = ","
)

/*
 * Private Functions
 */

func removeTargetFromSlice(src []int, target int) []int {
	newSlice := []int{}
	for _, e := range src {
		if e == target {
			continue
		}
		newSlice = append(newSlice, e)
	}
	return newSlice
}

func contains(src []int, target int) bool {
	for _, i := range src {
		if i == target {
			return true
		}
	}
	return false
}

/*
 * Public Functions
 */

func New(question string, choices []string) *Checkbox {
	/*
	 * checkbox is following format.
	 * +--------+---------+------------+---------+--------+------+
	 * | Prefix | Padding | Question   | Padding | Answer | \r\n |
	 * +--------+----+----+----+-------+--+------+--------+------+
	 * |ChosePrefix  | Padding | Choice 0 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 1 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 2 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 3 | \r\n |
	 * +-------------+---------+----------+------+
	 */

	// Create first line's ElemChain
	// +--------+---------+------------+---------+--------+------+
	// | Prefix | Padding | Question   | Padding | Answer | \r\n |
	// +--------+----+----+----+-------+--+------+--------+------+
	//
	firstLineElems := []element.Element{
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
			Str:    "",
			Colors: constants.AnswerColors,
		},
		// Next line
		{
			Str:    constants.NewLine,
			Colors: []string{},
		},
	}
	firstLineElemChain := elementChain.New(firstLineElems)

	// Create choices ElemChain
	// +-------------+---------+----------+------+
	// |ChosePrefix  | Padding | Choice 0 | \r\n |
	// +-------------+---------+----------+------+
	// |NoChosePrefix| Padding | Choice 1 | \r\n |
	// +-------------+---------+----------+------+
	// |NoChosePrefix| Padding | Choice 2 | \r\n |
	// +-------------+---------+----------+------+
	// |NoChosePrefix| Padding | Choice 3 | \r\n |
	// +-------------+---------+----------+------+
	//
	choicesElemChain := elementChain.New([]element.Element{})
	initialPointerPosition := 0
	for _, choice := range choices {
		// Prefix
		elemNoChosePrefix := element.New(NoChosePrefix, NoChosePrefixColors)
		// Padding
		elemPadding := element.New(constants.Padding, constants.PaddingColors)
		// Choice
		elemChoice := element.New(choice, NoChoseColors)
		// Next line
		elemNextLine := element.New(constants.NewLine, []string{})

		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemNoChosePrefix)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemPadding)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemChoice)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemNextLine)
	}

	checkbox := Checkbox{
		firstLineElemChain: *firstLineElemChain,
		choicesElemChain:   *choicesElemChain,
		choices:            choices,
	}

	checkbox.movePointerTo(initialPointerPosition)
	return &checkbox
}

/*
 * Private Methods
 */

func (checkbox *Checkbox) choose(choiceIndex int) {
	// Update chosePositions
	checkbox.chosePositions = append(checkbox.chosePositions, choiceIndex)
	// Update elems
	checkbox.updateElems()
}

func (checkbox *Checkbox) decrementPointer() {
	checkbox.movePointerTo(checkbox.pointerPosition - 1)
}

func (checkbox *Checkbox) incrementPointer() {
	checkbox.movePointerTo(checkbox.pointerPosition + 1)
}

func (checkbox *Checkbox) movePointerTo(pointerPosition int) {
	// Update pointer position
	checkbox.pointerPosition = pointerPosition
	// Update elems
	checkbox.updateElems()
}

func (checkbox *Checkbox) print() {
	checkbox.firstLineElemChain.Print()
	if !checkbox.finished && !checkbox.canceled {
		checkbox.choicesElemChain.Print()
	} else {
		checkbox.choicesElemChain.Erase()
	}
}

func (checkbox *Checkbox) setAnswerElem() {
	answers := constants.OpenParenthesis
	for i, chosePosition := range checkbox.chosePositions {
		answers += checkbox.choices[chosePosition]
		if i != len(checkbox.chosePositions)-1 {
			answers += AnswerSeparator
		}
	}
	if len(checkbox.chosePositions) == 0 {
		answers += AnswerNone
	}
	answers += constants.CloseParenthesis

	checkbox.firstLineElemChain.Elems[4].Str = answers
}

func (checkbox *Checkbox) unChoose(choiceIndex int) {
	// Update chosePositions
	checkbox.chosePositions = removeTargetFromSlice(checkbox.chosePositions, choiceIndex)
	// Update elems
	checkbox.updateElems()
}

func (checkbox *Checkbox) updateElems() {
	if checkbox.finished {
		checkbox.setAnswerElem()
		return
	}
	// Set chose / unset prefix and colors
	for i := 0; i < len(checkbox.choicesElemChain.Elems); i += 4 {
		// Convert element index to choice index
		choiceIndex := i / 4
		if contains(checkbox.chosePositions, choiceIndex) { // If chose
			// Prefix
			checkbox.choicesElemChain.Elems[i].Str = ChosePrefix
			checkbox.choicesElemChain.Elems[i].Colors = ChosePrefixColors
			// Choice
			checkbox.choicesElemChain.Elems[i+2].Colors = ChoseColors
		} else { // is not chose
			// Prefix
			checkbox.choicesElemChain.Elems[i].Str = NoChosePrefix
			checkbox.choicesElemChain.Elems[i].Colors = NoChosePrefixColors
			// Choice
			checkbox.choicesElemChain.Elems[i+2].Colors = NoChoseColors
		}
	}
	// Set pointer color
	pointerElemStart := checkbox.pointerPosition * 4
	checkbox.choicesElemChain.Elems[pointerElemStart].Colors =
		constants.PointerColors
	checkbox.choicesElemChain.Elems[pointerElemStart+2].Colors =
		constants.PointerColors
}

/*
 * Public Methods
 */

func (checkbox Checkbox) Erase() {
	checkbox.firstLineElemChain.Erase()
	if checkbox.finished || checkbox.canceled {
		return
	}
	checkbox.choicesElemChain.Erase()
}

func (checkbox Checkbox) GetMinX() int {
	firstMinX := checkbox.firstLineElemChain.GetMinX()
	choicesMinX := checkbox.choicesElemChain.GetMinX()

	if firstMinX < choicesMinX {
		return firstMinX
	} else {
		return choicesMinX
	}
}

func (checkbox Checkbox) GetMinY() int {
	firstMinY := checkbox.firstLineElemChain.GetMinY()
	choicesMinY := checkbox.choicesElemChain.GetMinY()

	if firstMinY < choicesMinY {
		return firstMinY
	} else {
		return choicesMinY
	}
}

func (checkbox Checkbox) GetMaxX() int {
	firstMaxX := checkbox.firstLineElemChain.GetMaxX()
	choicesMaxX := checkbox.choicesElemChain.GetMaxX()

	if firstMaxX < choicesMaxX {
		return firstMaxX
	} else {
		return choicesMaxX
	}
}

func (checkbox Checkbox) GetMaxY() int {
	firstMaxY := checkbox.firstLineElemChain.GetMaxY()
	choicesMaxY := checkbox.choicesElemChain.GetMaxY()

	if firstMaxY < choicesMaxY {
		return firstMaxY
	} else {
		return choicesMaxY
	}
}

func (checkbox Checkbox) GetStartX() int {
	return checkbox.firstLineElemChain.GetStartX()
}

func (checkbox Checkbox) GetStartY() int {
	return checkbox.firstLineElemChain.GetStartY()
}

func (checkbox Checkbox) GetEndX() int {
	return checkbox.choicesElemChain.GetEndX()
}

func (checkbox Checkbox) GetEndY() int {
	return checkbox.choicesElemChain.GetEndY()
}

func (checkbox *Checkbox) Ask() ([]int, bool) {
	checkbox.print()
	inputHelper.SetRaw(true)
	for {
		cursor.MoveCursorTo(checkbox.GetStartX(), checkbox.GetStartY())
		checkbox.print()
		if checkbox.finished || checkbox.canceled {
			break
		}
		// Get keyboard input
		inputHelper.SetNoEcho(true)
		inputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)

		switch string(inputRunes) {
		case constants.UpArrow: // up arrow key
			fallthrough
		case "k": // up
			if checkbox.pointerPosition > 0 {
				checkbox.decrementPointer()
			}
		case constants.DownArrow: // down arrow key
			fallthrough
		case "j": // down
			if checkbox.pointerPosition < len(checkbox.choices)-1 {
				checkbox.incrementPointer()
			}
		case constants.Enter: // enter
			checkbox.finished = true
			checkbox.updateElems()
		case " ": // space
			// update chose <-> no chose
			if contains(checkbox.chosePositions, checkbox.pointerPosition) {
				checkbox.unChoose(checkbox.pointerPosition)
			} else {
				checkbox.choose(checkbox.pointerPosition)
			}
		case constants.CtrlC: // ctrl + c
			checkbox.canceled = true
		}
	}
	inputHelper.SetRaw(false)

	var answers []string
	if checkbox.finished {
		for _, chosePosition := range checkbox.chosePositions {
			answers = append(answers, checkbox.choices[chosePosition])
		}
	}
	return checkbox.chosePositions, checkbox.canceled
}
