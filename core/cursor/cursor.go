package cursor

/*
 * Module Dependencies
 */

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mozzzzy/cui/v3/color"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

var (
	cursorX, cursorY int
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func GetCursor() (int, int) {
	return cursorX, cursorY
}

func MoveCursorTo(x, y int) {
	if x < 0 || y < 0 {
		panic("Invalid coordinates (x, y) = (" + strconv.Itoa(x) + ", " + strconv.Itoa(y) + ")")
	}
	/*
	 *    ^      A
	 *  <-|->  D   C
	 *    v      B
	 *
	 *  \e[n;mH : Go to (n,m)
	 */
	if cursorX-x > 0 {
		fmt.Printf("\033[%dD", cursorX-x)
	} else if cursorX-x < 0 {
		fmt.Printf("\033[%dC", x-cursorX)
	}
	if cursorY-y > 0 {
		fmt.Printf("\033[%dA", cursorY-y)
	} else if cursorY-y < 0 {
		fmt.Printf("\033[%dB", y-cursorY)
	}
	cursorX = x
	cursorY = y
}

func Print(str string, colors []string) {
	// Ignore "\r"
	str = strings.ReplaceAll(str, "\r", "")

	lines := strings.Split(str, "\n")
	lineNum := len(lines)
	for i, line := range lines {
		for _, c := range colors {
			fmt.Print(c)
		}
		fmt.Print(line)
		fmt.Print(color.Default)
		cursorX += len(line)
		if i < lineNum-1 {
			fmt.Print("\r\n")
			cursorX = 0
			cursorY++
		}
	}
}

func MoveCursorToZeroZero() {
	MoveCursorTo(0, 0)
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */
