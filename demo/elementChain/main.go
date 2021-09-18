package main

/*
 * NOTE
 * ElementChain type is for internal operations.
 * Each packages like Checkbox, List and so on have ElementChains inside of them.
 */

import (
	"time"

	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/cursor"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
)

func main() {
	elems0 := []element.Element{
		{
			Str:    "Hello ",
			Colors: []string{color.Bold, color.BlueFg},
		},
		{
			Str:    "World",
			Colors: []string{color.Bold, color.CyanFg},
		},
		{
			Str:    "!!",
			Colors: []string{color.Bold, color.GreenFg},
		},
		{
			Str:    "\n",
			Colors: []string{},
		},
	}
	elemChain0 := elementChain.New(elems0)

	elems1 := []element.Element{
		{
			Str:    "Hello ",
			Colors: []string{color.Bold, color.BlueBg},
		},
		{
			Str:    "World",
			Colors: []string{color.Bold, color.CyanBg},
		},
		{
			Str:    "!!",
			Colors: []string{color.Bold, color.GreenBg},
		},
		{
			Str:    "\n",
			Colors: []string{},
		},
	}
	elemChain1 := elementChain.New(elems1)

	elems2 := []element.Element{
		{
			Str:    "Hello ",
			Colors: []string{color.Bold, color.BlueFg},
		},
		{
			Str:    "World",
			Colors: []string{color.Bold, color.CyanBg},
		},
		{
			Str:    "!!",
			Colors: []string{color.Bold, color.GreenFg},
		},
		{
			Str:    "\n",
			Colors: []string{},
		},
	}
	elemChain2 := elementChain.New(elems2)

	elems3 := []element.Element{
		{
			Str:    "Hello ",
			Colors: []string{color.Bold, color.BlueBg},
		},
		{
			Str:    "World",
			Colors: []string{color.Bold, color.CyanFg},
		},
		{
			Str:    "!!",
			Colors: []string{color.Bold, color.GreenBg},
		},
		{
			Str:    "\n",
			Colors: []string{},
		},
	}
	elemChain3 := elementChain.New(elems3)


	elemChain0.Print()
	time.Sleep(1 * time.Second)
	elemChain1.Print()
	time.Sleep(1 * time.Second)
	elemChain2.Print()
	time.Sleep(1 * time.Second)
	elemChain3.Print()
	time.Sleep(1 * time.Second)

	elemChain0.Erase()
	time.Sleep(1 * time.Second)
	elemChain1.Erase()
	time.Sleep(1 * time.Second)
	elemChain2.Erase()
	time.Sleep(1 * time.Second)
	elemChain3.Erase()
	time.Sleep(1 * time.Second)

	cursor.MoveCursorToZeroZero()
}
