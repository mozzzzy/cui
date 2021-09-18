package list

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

type List struct {
	firstLineElemChain elementChain.ElementChain
	choicesElemChain   elementChain.ElementChain
	choices            []string
	pointerPosition    int
	finished           bool
	canceled           bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	Pointer            string   = "❯"
	PointerSpace       string   = " "
	PointerSpaceColors []string = []string{}

	ChoiceColors   []string = []string{color.CyanFg, color.Bold}
	NoChoiceColors []string = []string{}
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(question string, choices []string) *List {
	/*
	 * list is following format.
	 * +--------+---------+----------+---------+--------+------+
	 * | Prefix | Padding | Question | Padding | Answer | \r\n |
	 * +--------+----+----+----+-----+----+-----++-------+------+
	 * |Pointer      | Padding | Choice 0 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 1 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 2 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 3 | \r\n |
	 * +-------------+---------+----------+------+
	 */
	// Create first line's ElemChain
	// +--------+---------+----------+---------+--------+------+
	// | Prefix | Padding | Question | Padding | Answer | \r\n |
	// +--------+---------+----------+---------+--------+------+
	//
	elemsFirstLine := []element.Element{
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
	firstLineElemChain := elementChain.New(elemsFirstLine)

	// Create choices ElemChain
	// +-------------+---------+----------+------+
	// |Pointer      | Padding | Choice 0 | \r\n |
	// +-------------+---------+----------+------+
	// |PointerSpace | Padding | Choice 1 | \r\n |
	// +-------------+---------+----------+------+
	// |PointerSpace | Padding | Choice 2 | \r\n |
	// +-------------+---------+----------+------+
	// |PointerSpace | Padding | Choice 3 | \r\n |
	// +-------------+---------+----------+------+
	//
	choicesElemChain := elementChain.New([]element.Element{})
	initialPointerPosition := 0
	for i, choice := range choices {
		// Pointer space
		elemPointerSpace := element.New(PointerSpace, PointerSpaceColors)
		// Padding
		elemPadding := element.New(constants.Padding, constants.PaddingColors)
		// Choice
		var elemChoice *element.Element
		if i == initialPointerPosition {
			elemChoice = element.New(choice, ChoiceColors)
		} else {
			elemChoice = element.New(choice, NoChoiceColors)
		}
		// Next line
		elemNextLine := element.New(constants.NewLine, []string{})

		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemPointerSpace)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemPadding)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemChoice)
		choicesElemChain.Elems = append(choicesElemChain.Elems, *elemNextLine)
	}

	list := List{
		firstLineElemChain: *firstLineElemChain,
		choicesElemChain:   *choicesElemChain,
		choices:            choices,
	}

	list.setPointer(initialPointerPosition)
	return &list
}

/*
 * Private Methods
 */

func (list *List) print() {
	list.firstLineElemChain.Print()
	if !list.finished && !list.canceled {
		list.choicesElemChain.Print()
	} else {
		list.choicesElemChain.Erase()
		cursor.MoveCursorTo(list.firstLineElemChain.GetEndX(), list.firstLineElemChain.GetEndY())
	}
}

func (list *List) setAnswerElem() {
	list.firstLineElemChain.Elems[4].Str =
		constants.OpenParenthesis +
			list.choices[list.pointerPosition] +
			constants.CloseParenthesis
}

func (list *List) setPointer(pointerPosition int) {
	if list.finished {
		return
	}
	list.pointerPosition = pointerPosition
	for i := 0; i < len(list.choicesElemChain.Elems); i += 4 {
		// convert element index to choice index
		choiceIndex := i / 4
		if choiceIndex == list.pointerPosition {
			list.choicesElemChain.Elems[i].Str = Pointer
			list.choicesElemChain.Elems[i].Colors = constants.PointerColors
			list.choicesElemChain.Elems[i+2].Colors = ChoiceColors
		} else {
			list.choicesElemChain.Elems[i].Str = PointerSpace
			list.choicesElemChain.Elems[i].Colors = PointerSpaceColors
			list.choicesElemChain.Elems[i+2].Colors = NoChoiceColors
		}
	}
}

func (list *List) decrementPointer() {
	list.setPointer(list.pointerPosition - 1)
}

func (list *List) incrementPointer() {
	list.setPointer(list.pointerPosition + 1)
}

/*
 * Public Methods
 */

func (list List) Erase() {
	list.firstLineElemChain.Erase()
	if list.finished || list.canceled {
		return
	}
	list.choicesElemChain.Erase()
}

func (list List) GetMinX() int {
	firstMinX := list.firstLineElemChain.GetMinX()
	choicesMinX := list.choicesElemChain.GetMinX()

	if firstMinX < choicesMinX {
		return firstMinX
	} else {
		return choicesMinX
	}
}

func (list List) GetMinY() int {
	firstMinY := list.firstLineElemChain.GetMinY()
	choicesMinY := list.choicesElemChain.GetMinY()

	if firstMinY < choicesMinY {
		return firstMinY
	} else {
		return choicesMinY
	}
}

func (list List) GetMaxX() int {
	firstMaxX := list.firstLineElemChain.GetMaxX()
	choicesMaxX := list.choicesElemChain.GetMaxX()

	if firstMaxX < choicesMaxX {
		return firstMaxX
	} else {
		return choicesMaxX
	}
}

func (list List) GetMaxY() int {
	firstMaxY := list.firstLineElemChain.GetMaxY()
	choicesMaxY := list.choicesElemChain.GetMaxY()

	if firstMaxY < choicesMaxY {
		return firstMaxY
	} else {
		return choicesMaxY
	}
}

func (list List) GetStartX() int {
	return list.firstLineElemChain.GetStartX()
}

func (list List) GetStartY() int {
	return list.firstLineElemChain.GetStartY()
}

func (list List) GetEndX() int {
	return list.choicesElemChain.GetEndX()
}

func (list List) GetEndY() int {
	return list.choicesElemChain.GetEndY()
}

func (list *List) Ask() (int, bool) {
	list.print()
	for {
		cursor.MoveCursorTo(list.GetStartX(), list.GetStartY())
		list.print()
		if list.finished || list.canceled {
			break
		}
		// Get keyboard input
		inputHelper.SetRaw(true)
		inputHelper.SetNoEcho(true)
		inputRunes := inputHelper.Getch()
		inputHelper.SetNoEcho(false)
		inputHelper.SetRaw(false)

		switch string(inputRunes) {
		case constants.UpArrow: // up arrow
			fallthrough
		case "k": // up
			if list.pointerPosition > 0 {
				list.decrementPointer()
			}
		case constants.DownArrow: // down arrow
			fallthrough
		case "j": // down
			if list.pointerPosition < len(list.choices)-1 {
				list.incrementPointer()
			}
		case constants.Enter: // enter
			list.setAnswerElem()
			list.finished = true
		case constants.CtrlC: // ctrl + c
			list.canceled = true
		}
	}
	if list.finished {
		return list.pointerPosition, list.canceled
	}
	return 0, list.canceled
}
