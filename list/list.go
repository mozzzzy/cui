package list

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/constants"
	"github.com/mozzzzy/cui/v2/cursor"
	"github.com/mozzzzy/cui/v2/element"
	"github.com/mozzzzy/cui/v2/elementChain"
	"github.com/mozzzzy/cui/v2/inputHelper"
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
	 * +--------+---------+------------+---------+--------+------+
	 * | Prefix | Padding | Question   | Padding | Answer | \r\n |
	 * +--------+----+----+----+-------+-+------++--------+------+
	 * |Pointer      | Padding | Choice0 | \r\n |
	 * +-------------+---------+---------+------+
	 * |PointerSpace | Padding | Choice1 | \r\n |
	 * +-------------+---------+---------+------+
	 * |PointerSpace | Padding | Choice2 | \r\n |
	 * +-------------+---------+---------+------+
	 * |PointerSpace | Padding | Choice3 | \r\n |
	 * +-------------+---------+---------+------+
	 */
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

	list.SetPointer(initialPointerPosition)
	return &list
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

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
	inputHelper.SetRaw(true)
	list.Print()
	for {
		cursor.MoveCursorTo(list.GetStartX(), list.GetStartY())
		list.Print()
		if list.finished || list.canceled {
			break
		}
		// Get keyboard input
		inputRunes := inputHelper.Getch()
		switch string(inputRunes) {
		case constants.UpArrow: // up arrow
			fallthrough
		case "k": // up
			if list.pointerPosition > 0 {
				list.DecrementPointer()
			}
		case constants.DownArrow: // down arrow
			fallthrough
		case "j": // down
			if list.pointerPosition < len(list.choices)-1 {
				list.incrementPointer()
			}
		case constants.Enter: // enter
			list.SetAnswerElem()
			list.finished = true
		case constants.CtrlC: // ctrl + c
			list.canceled = true
		}
	}
	inputHelper.SetRaw(false)
	if list.finished {
		return list.pointerPosition, list.canceled
	}
	return 0, list.canceled
}

func (list *List) Print() {
	list.firstLineElemChain.Print()
	if !list.finished {
		list.choicesElemChain.Print()
	} else {
		list.choicesElemChain.Erase()
		cursor.MoveCursorTo(list.firstLineElemChain.GetEndX(), list.firstLineElemChain.GetEndY())
	}
}

func (list *List) SetAnswerElem() {
	list.firstLineElemChain.Elems[4].Str =
		constants.OpenParenthesis +
			list.choices[list.pointerPosition] +
			constants.CloseParenthesis
}

func (list *List) SetPointer(pointerPosition int) {
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

func (list *List) DecrementPointer() {
	list.SetPointer(list.pointerPosition - 1)
}

func (list *List) incrementPointer() {
	list.SetPointer(list.pointerPosition + 1)
}
