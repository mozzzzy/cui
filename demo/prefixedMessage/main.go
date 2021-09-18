package main

/*
 * NOTE
 * PrefixedMessage type is for internal operations.
 * Each packages like InfoMessage, ErrorMessage and so on have PrefixedMessage inside of them.
 */

import (
	"time"

	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/prefixedMessage"
)

func main() {
	prefix := " Prefix "
	prefixColors := []string{color.GreenBg, color.Bold}
	padding := " "
	paddingColors := []string{}
	str := "Hello World!!"
	colors := []string{}
	pmsg1 := prefixedMessage.New(prefix, prefixColors, padding, paddingColors, str, colors)
	pmsg1.Print()

	time.Sleep(1 * time.Second)
	pmsg1.Erase()
	time.Sleep(1 * time.Second)
	pmsg1.SetPrefix(prefix, []string{color.CyanBg, color.Bold})
	pmsg1.Print()
	time.Sleep(1 * time.Second)
	pmsg1.Erase()
	time.Sleep(1 * time.Second)
}
