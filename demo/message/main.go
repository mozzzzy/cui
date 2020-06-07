package main

import (
	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2"
)

func main() {
	cui.Message("Hello World!!", []string{color.Bold, color.BlueFg})
}
