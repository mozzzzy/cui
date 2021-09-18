package main

/*
 * NOTE
 * Element type is for internal operations.
 * Each packages like Checkbox, List and so on have Elements inside of them.
 */

import (
	"time"

	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/element"
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
}
