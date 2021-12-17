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
	choicesElemChains []elementChain.ElementChain
	choices           []string
	chosePositions    []int
	pointerPosition   int
	finished          bool
	canceled          bool
	printed           bool
	onePageLineLen    int
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

func New(choices []string) *Checkbox {
	/*
	 * checkbox is following format.
	 * +-------------+---------+----------+------+
	 * |ChosePrefix  | Padding | Choice 0 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 1 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 2 | \r\n |
	 * +-------------+---------+----------+------+
	 * |NoChosePrefix| Padding | Choice 3 | \r\n |
	 * +-------------+---------+----------+------+
	 */
	choicesElemChains := []elementChain.ElementChain{}

	for _, choice := range choices {
		choiceElemChain := elementChain.New([]element.Element{})
		// Prefix
		elemNoChosePrefix := element.New(NoChosePrefix, NoChosePrefixColors)
		// Padding
		elemPadding := element.New(constants.Padding, constants.PaddingColors)
		// Choice
		elemChoice := element.New(choice, NoChoseColors)
		// Next line
		elemNextLine := element.New(constants.NewLine, []string{})

		choiceElemChain.Elems = append(choiceElemChain.Elems, *elemNoChosePrefix)
		choiceElemChain.Elems = append(choiceElemChain.Elems, *elemPadding)
		choiceElemChain.Elems = append(choiceElemChain.Elems, *elemChoice)
		choiceElemChain.Elems = append(choiceElemChain.Elems, *elemNextLine)

		choicesElemChains = append(choicesElemChains, *choiceElemChain)
	}

	checkbox := Checkbox{
		choicesElemChains: choicesElemChains,
		choices:           choices,
	}

	initialPointerPosition := 0
	checkbox.movePointerTo(initialPointerPosition)

	termLineLen, err := cursor.GetTermLineLen()
	if err != nil {
		panic("Failed to get terminal line length : " + err.Error())
	}
	checkbox.onePageLineLen = termLineLen - 3

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
	if checkbox.pointerPosition > 0 {
		checkbox.movePointerTo(checkbox.pointerPosition - 1)
	}
}

func (checkbox *Checkbox) incrementPointer() {
	if checkbox.pointerPosition < len(checkbox.choices)-1 {
		checkbox.movePointerTo(checkbox.pointerPosition + 1)
	}
}

func (checkbox *Checkbox) movePointerTo(pointerPosition int) {
	// Update pointer position
	checkbox.pointerPosition = pointerPosition
	// Update elems
	checkbox.updateElems()
}

func (checkbox *Checkbox) unChoose(choiceIndex int) {
	// Update chosePositions
	checkbox.chosePositions = removeTargetFromSlice(checkbox.chosePositions, choiceIndex)
	// Update elems
	checkbox.updateElems()
}

func (checkbox *Checkbox) updateElems() {
	// Update prefix and colors
	for i := 0; i < len(checkbox.choicesElemChains); i++ {
		if contains(checkbox.chosePositions, i) { // If chose
			// Prefix
			checkbox.choicesElemChains[i].Elems[0].Str = ChosePrefix
			checkbox.choicesElemChains[i].Elems[0].Colors = ChosePrefixColors
			// Choice
			checkbox.choicesElemChains[i].Elems[2].Colors = ChoseColors
		} else { // if not chose
			// Prefix
			checkbox.choicesElemChains[i].Elems[0].Str = NoChosePrefix
			checkbox.choicesElemChains[i].Elems[0].Colors = NoChosePrefixColors
			// Choice
			checkbox.choicesElemChains[i].Elems[2].Colors = NoChoseColors
		}
	}
	// Set pointer color
	checkbox.choicesElemChains[checkbox.pointerPosition].Elems[0].Colors =
		constants.PointerColors
	checkbox.choicesElemChains[checkbox.pointerPosition].Elems[2].Colors =
		constants.PointerColors
}

func (checkbox *Checkbox) updatePage(goingDown bool) {
	var lastPointerPosition int
	if goingDown {
		lastPointerPosition = checkbox.pointerPosition - 1
	} else {
		lastPointerPosition = checkbox.pointerPosition + 1
	}

	lastPage := lastPointerPosition / checkbox.onePageLineLen
	currentPage := checkbox.pointerPosition / checkbox.onePageLineLen

	if lastPage == currentPage {
		return
	}

	lastPageTop := checkbox.onePageLineLen * lastPage
	lastPageBottom := lastPageTop + checkbox.onePageLineLen - 1
	for i := lastPageTop; i <= lastPageBottom; i++ {
		if i >= len(checkbox.choices) {
			break
		}
		checkbox.choicesElemChains[i].Erase()
	}
	cursor.MoveCursorTo(checkbox.GetStartX(), checkbox.GetStartY())
	checkbox.printed = false
}

/*
 * Public Methods
 */

func (checkbox Checkbox) Erase() {
	if !checkbox.printed {
		return
	}
	currentPage := checkbox.pointerPosition / checkbox.onePageLineLen
	currentPageTop := checkbox.onePageLineLen * currentPage
	currentPageBottom := currentPageTop + checkbox.onePageLineLen - 1

	for i := currentPageTop; i <= currentPageBottom; i++ {
		if i >= len(checkbox.choices) {
			break
		}
		checkbox.choicesElemChains[i].Erase()
	}
	cursor.MoveCursorTo(
		checkbox.choicesElemChains[0].GetStartX(),
		checkbox.choicesElemChains[0].GetStartY())
}

func (checkbox Checkbox) GetMinX() int {
	var minX int
	for _, elemChain := range checkbox.choicesElemChains {
		currentMinX := elemChain.GetMinX()
		if minX > currentMinX {
			minX = currentMinX
		}
	}
	return minX
}

func (checkbox Checkbox) GetMinY() int {
	return checkbox.choicesElemChains[0].GetMinY()
}

func (checkbox Checkbox) GetMaxX() int {
	var maxX int
	for _, elemChain := range checkbox.choicesElemChains {
		currentMaxX := elemChain.GetMaxX()
		if maxX < currentMaxX {
			maxX = currentMaxX
		}
	}
	return maxX
}

func (checkbox Checkbox) GetMaxY() int {
	return checkbox.choicesElemChains[len(checkbox.choices)].GetMaxY()
}

func (checkbox Checkbox) GetStartX() int {
	return checkbox.choicesElemChains[0].GetStartX()
}

func (checkbox Checkbox) GetStartY() int {
	return checkbox.choicesElemChains[0].GetStartY()
}

func (checkbox Checkbox) GetEndX() int {
	return checkbox.choicesElemChains[len(checkbox.choices)].GetEndX()
}

func (checkbox Checkbox) GetEndY() int {
	return checkbox.choicesElemChains[len(checkbox.choices)].GetEndY()
}

func (checkbox *Checkbox) Ask() ([]int, bool, *Checkbox) {
	checkbox.Print()
	inputHelper.SetRaw(true)
	inputHelper.SetNoEcho(true)
	for {
		cursor.MoveCursorTo(checkbox.GetStartX(), checkbox.GetStartY())
		checkbox.Print()
		cursor.MoveCursorTo(
			checkbox.choicesElemChains[checkbox.pointerPosition].GetStartX(),
			checkbox.choicesElemChains[checkbox.pointerPosition].GetStartY())
		if checkbox.finished || checkbox.canceled {
			break
		}
		// Get keyboard input
		inputRunes := inputHelper.Getch()

		switch string(inputRunes) {
		case constants.UpArrow: // up arrow key
			fallthrough
		case "k": // up
			if checkbox.pointerPosition > 0 {
				checkbox.decrementPointer()
				checkbox.updatePage(false)
			}
		case constants.DownArrow: // down arrow key
			fallthrough
		case "j": // down
			if checkbox.pointerPosition < len(checkbox.choices)-1 {
				checkbox.incrementPointer()
				checkbox.updatePage(true)
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
	inputHelper.SetNoEcho(false)
	inputHelper.SetRaw(false)

	var answers []string
	if checkbox.finished {
		for _, chosePosition := range checkbox.chosePositions {
			answers = append(answers, checkbox.choices[chosePosition])
		}
	}
	return checkbox.chosePositions, checkbox.canceled, checkbox
}

func (checkbox *Checkbox) Print() {
	currentPage := checkbox.pointerPosition / checkbox.onePageLineLen
	currentPageTop := checkbox.onePageLineLen * currentPage
	currentPageBottom := currentPageTop + checkbox.onePageLineLen - 1

	if checkbox.printed {
		if checkbox.pointerPosition > currentPageTop {
			cursor.MoveCursorTo(
				checkbox.choicesElemChains[checkbox.pointerPosition-1].GetStartX(),
				checkbox.choicesElemChains[checkbox.pointerPosition-1].GetStartY())
			checkbox.choicesElemChains[checkbox.pointerPosition-1].Print()
		}
		if checkbox.pointerPosition < len(checkbox.choices)-1 && checkbox.pointerPosition < currentPageBottom {
			cursor.MoveCursorTo(
				checkbox.choicesElemChains[checkbox.pointerPosition+1].GetStartX(),
				checkbox.choicesElemChains[checkbox.pointerPosition+1].GetStartY())
			checkbox.choicesElemChains[checkbox.pointerPosition+1].Print()
		}
		cursor.MoveCursorTo(
			checkbox.choicesElemChains[checkbox.pointerPosition].GetStartX(),
			checkbox.choicesElemChains[checkbox.pointerPosition].GetStartY())
		checkbox.choicesElemChains[checkbox.pointerPosition].Print()
	} else {
		for i := currentPageTop; i <= currentPageBottom; i++ {
			if i >= len(checkbox.choices) {
				break
			}
			checkbox.choicesElemChains[i].Print()
		}
		checkbox.printed = true
	}
}
