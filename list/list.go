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
	choicesElemChain elementChain.ElementChain
	choices          []string
	pointerPosition  int
	finished         bool
	canceled         bool
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

func New(choices []string) *List {
	/*
	 * list is following format.
	 * +-------------+---------+----------+------+
	 * |Pointer      | Padding | Choice 0 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 1 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 2 | \r\n |
	 * +-------------+---------+----------+------+
	 * |PointerSpace | Padding | Choice 3 | \r\n |
	 * +-------------+---------+----------+------+
	 */
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
		choicesElemChain: *choicesElemChain,
		choices:          choices,
	}

	list.setPointer(initialPointerPosition)
	return &list
}

/*
 * Private Methods
 */

func (list *List) print() {
	list.choicesElemChain.Print()
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
	list.choicesElemChain.Erase()
}

func (list List) GetMinX() int {
	return list.choicesElemChain.GetMinX()
}

func (list List) GetMinY() int {
	return list.choicesElemChain.GetMinY()
}

func (list List) GetMaxX() int {
	return list.choicesElemChain.GetMaxX()
}

func (list List) GetMaxY() int {
	return list.choicesElemChain.GetMaxY()
}

func (list List) GetStartX() int {
	return list.choicesElemChain.GetStartX()
}

func (list List) GetStartY() int {
	return list.choicesElemChain.GetStartY()
}

func (list List) GetEndX() int {
	return list.choicesElemChain.GetEndX()
}

func (list List) GetEndY() int {
	return list.choicesElemChain.GetEndY()
}

func (list *List) Ask() (int, bool, *List) {
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
			list.finished = true
		case constants.CtrlC: // ctrl + c
			list.canceled = true
		}
	}
	if list.finished {
		return list.pointerPosition, list.canceled, list
	}
	return 0, list.canceled, list
}
