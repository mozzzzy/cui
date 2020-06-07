package main

/*
 * NOTE
 * Element type is for internal operations.
 * Each packages like Checkbox, List and so on have Elements inside of them.
 */

import (
	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/debugUtil"
	"github.com/mozzzzy/cui/v2/element"
	"time"
)

func main() {
	elem0 := element.New("Hello World!!", []string{color.GreenFg, color.Bold})
	elem0.Print()
	time.Sleep(1 * time.Second)
	elem0.Erase()
	time.Sleep(1 * time.Second)

	elem1 := element.New("Hello World!!", []string{color.CyanFg, color.Bold})
	elem1.Print()
	time.Sleep(1 * time.Second)
	elem1.Erase()
	time.Sleep(1 * time.Second)

	elem2 := element.New("Hello World!!", []string{color.BlueFg, color.Bold})
	elem2.Print()
	time.Sleep(1 * time.Second)
	elem2.Erase()
	time.Sleep(1 * time.Second)

	debugUtil.DebugCoordinate(elem0)
}
