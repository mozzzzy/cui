package element

/*
 * Module Dependencies
 */

import (
	"strings"

	"github.com/mozzzzy/cui/v2/cursor"
)

/*
 * Types
 */

type Element struct {
	Str    string
	Colors []string
	// NOTE
	// Only startX and startY is specified in out of Element.
	// Other coordinates are calculated by element it self.
	minX, maxX   int
	minY, maxY   int
	startX, endX int
	startY, endY int
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

func New(str string, colors []string) *Element {
	var elem Element
	elem.Str = str
	elem.Colors = colors
	return &elem
}

/*
 * Private Methods
 */

func (elem *Element) calculateCoordinates(xCursor, yCursor int) {
	// startX
	elem.startX = xCursor
	// startY
	elem.startY = yCursor

	elemLines := strings.Split(elem.Str, "\n")
	// minX
	if len(elemLines) > 1 {
		elem.minX = 0
	} else {
		elem.minX = elem.startX
	}
	// maxX
	for index, line := range elemLines {
		var oneLineLen int
		if index == 0 {
			oneLineLen = elem.startX + len(line)
		} else {
			oneLineLen = len(line)
		}
		if oneLineLen > elem.maxX {
			elem.maxX = oneLineLen
		}
	}
	// minY
	elem.minY = elem.startY
	// maxY
	elem.maxY = elem.startY + len(elemLines) - 1

	// endX
	if len(elemLines) == 1 {
		elem.endX = elem.startX + len(elemLines[len(elemLines)-1])
	} else {
		elem.endX = len(elemLines[len(elemLines)-1])
	}
	// endY
	elem.endY = elem.maxY
}

/*
 * Public Methods
 */

func (elem Element) GetMinX() int {
	return elem.minX
}

func (elem Element) GetMinY() int {
	return elem.minY
}

func (elem Element) GetMaxX() int {
	return elem.maxX
}

func (elem Element) GetMaxY() int {
	return elem.maxY
}

func (elem Element) GetStartX() int {
	return elem.startX
}

func (elem Element) GetStartY() int {
	return elem.startY
}

func (elem Element) GetEndX() int {
	return elem.endX
}

func (elem Element) GetEndY() int {
	return elem.endY
}

func (elem *Element) Print() {
	xCursor, yCursor := cursor.GetCursor()

	// NOTE
	// The coordinates of element is updated only when this Print() function is called.
	// They are calculated by given xCursor, yCursor and elem.Str.
	elem.calculateCoordinates(xCursor, yCursor)
	cursor.Print(elem.Str, elem.Colors)
}

func (elem Element) Erase() {
	cursor.MoveCursorTo(elem.GetMinX(), elem.GetMinY())
	for y := elem.GetMinY(); y <= elem.GetMaxY(); y++ {
		for x := elem.GetMinX(); x <= elem.GetMaxX(); x++ {
			cursor.Print(" ", []string{})
		}
		if y != elem.GetMaxY() {
			cursor.Print("\n", []string{})
		}
	}
	cursor.MoveCursorTo(elem.GetStartX(), elem.GetStartY())
}
