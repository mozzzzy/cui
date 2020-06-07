package debugUtil

/*
 * Module Dependencies
 */

import (
	"fmt"
	"github.com/mozzzzy/cui/v2/cursor"
)

/*
 * Types
 */

type printable interface {
	GetMinX() int
	GetMinY() int
	GetMaxX() int
	GetMaxY() int
	GetStartX() int
	GetStartY() int
	GetEndX() int
	GetEndY() int
	Print()
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

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func DebugCoordinate(p printable) {
	fmt.Print("\n")
	fmt.Printf("startX: %d\n", p.GetStartX())
	fmt.Printf("startY: %d\n", p.GetStartY())
	fmt.Printf("endX: %d\n", p.GetEndX())
	fmt.Printf("endY: %d\n", p.GetEndY())
	fmt.Printf("minX: %d\n", p.GetMinX())
	fmt.Printf("minY: %d\n", p.GetMinY())
	fmt.Printf("maxX: %d\n", p.GetMaxX())
	fmt.Printf("maxY: %d\n", p.GetMaxY())

	cursorX, cursorY := cursor.GetCursor()
	fmt.Printf("cursorX: %d\n", cursorX)
	fmt.Printf("cursorY: %d\n", cursorY)
}
