package main

import (
	"github.com/mozzzzy/cui/color"
	"github.com/mozzzzy/cui"
)

func main() {
	prefix := " Prefix "
	prefixColors := []string{color.GreenBg, color.Bold}
	padding := " "
	paddingColors := []string{}
	str := "Hello World!!"
	colors := []string{}

	cui.PrefixedMessage(prefix, prefixColors, padding, paddingColors, str, colors)
}
