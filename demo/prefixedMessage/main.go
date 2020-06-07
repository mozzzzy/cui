package main

import (
	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2"
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
