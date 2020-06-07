package main

/*
 * NOTE
 * ElementChain type is for internal operations.
 * Each packages like Checkbox, List and so on have ElementChains inside of them.
 */

import (
	"time"

	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/debugUtil"
	"github.com/mozzzzy/cui/v2/element"
	"github.com/mozzzzy/cui/v2/elementChain"
)

func main() {
	elem0 := element.New("Hello", []string{color.Bold, color.BlueFg})
	elem1 := element.New(" ", []string{})
	elem2 := element.New("World", []string{color.Bold, color.CyanFg})
	elem3 := element.New("!!", []string{color.Bold, color.GreenFg})

	elemChain0 := elementChain.New(
		[]element.Element{*elem0, *elem1, *elem2, *elem3})
	elemChain0.Print()

	time.Sleep(1 * time.Second)
	elemChain0.Erase()
	time.Sleep(1 * time.Second)

	debugUtil.DebugCoordinate(elemChain0)
}
