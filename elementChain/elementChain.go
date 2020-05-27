package elementChain

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/cursor"
	"github.com/mozzzzy/cui/element"
)

/*
 * Types
 */

type ElementChain struct {
	Elems []element.Element
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

func New(elems []element.Element) *ElementChain {
	var elemChain ElementChain
	elemChain.Elems = elems
	return &elemChain
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (elemChain ElementChain) GetMinX() int {
	var minX int
	for _, elem := range elemChain.Elems {
		if elem.GetMinX() < minX {
			minX = elem.GetMinX()
		}
	}
	return minX
}

func (elemChain ElementChain) GetMinY() int {
	var minY int
	for _, elem := range elemChain.Elems {
		if minY > elem.GetMinY() {
			minY = elem.GetMinY()
		}
	}
	return minY
}

func (elemChain ElementChain) GetMaxX() int {
	var maxX int
	for _, elem := range elemChain.Elems {
		if maxX < elem.GetMaxX() {
			maxX = elem.GetMaxX()
		}
	}
	return maxX
}

func (elemChain ElementChain) GetMaxY() int {
	var maxY int
	for _, elem := range elemChain.Elems {
		if maxY < elem.GetMaxY() {
			maxY = elem.GetMaxY()
		}
	}
	return maxY
}

func (elemChain ElementChain) GetStartX() int {
	firstElem := elemChain.Elems[0]
	return firstElem.GetStartX()
}

func (elemChain ElementChain) GetStartY() int {
	firstElem := elemChain.Elems[0]
	return firstElem.GetStartY()
}

func (elemChain ElementChain) GetEndX() int {
	lastElem := elemChain.Elems[len(elemChain.Elems)-1]
	return lastElem.GetEndX()
}

func (elemChain ElementChain) GetEndY() int {
	lastElem := elemChain.Elems[len(elemChain.Elems)-1]
	return lastElem.GetEndY()
}

func (elemChain *ElementChain) Print() {
	for i := 0; i < len(elemChain.Elems); i++ {
		elemChain.Elems[i].Print()
	}
}

func (elemChain ElementChain) Erase() {
	for i := 0; i < len(elemChain.Elems); i++ {
		elemChain.Elems[i].Erase()
	}
	cursor.MoveCursorTo(elemChain.GetStartX(), elemChain.GetStartY())
}
