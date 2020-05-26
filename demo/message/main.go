package main

import (
	"github.com/mozzzzy/cui/color"
	"github.com/mozzzzy/cui"
)

func main() {
	cui.Message("Hello World!!", []string{color.Bold, color.BlueFg})
}
