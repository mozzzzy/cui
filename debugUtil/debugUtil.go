package debugUtil

/*
 * Module Dependencies
 */

import (
	"strconv"
	"time"

	"github.com/mozzzzy/cui/v3/core/cursor"
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
	//Print()
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

func Debug(p printable) {
	cursorX, cursorY := cursor.GetCursor()
	cursor.MoveCursorTo(0, 20)
	cursor.Print(
		"-------------------------------------\n"+
			"elem.startX : "+strconv.Itoa(p.GetStartX())+"\n"+
			"elem.startY : "+strconv.Itoa(p.GetStartY())+"\n"+
			"cursorX : "+strconv.Itoa(cursorX)+"\n"+
			"cursorY : "+strconv.Itoa(cursorY)+"\n",
		[]string{})
	cursor.MoveCursorTo(p.GetStartX(), p.GetStartY())
	time.Sleep(1 * time.Second)
	cursor.MoveCursorTo(0, 20)
	cursor.Print(
		"                                                                 \n"+
		"                                                                 \n"+
		"                                                                 \n"+
		"                                                                 \n"+
		"                                                                 ",
		[]string{})
	cursor.MoveCursorTo(cursorX, cursorY)
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

/*
func DebugCoordinate(p printable) {
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
*/
