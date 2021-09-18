package main

import (
	"time"

	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of message package.")

	cui.Message("Some Message !!", []string{color.Bold, color.GreenFg})
	cui.Message("Some Message !!", []string{color.Bold, color.GreenBg})
	cui.Message("Some Message !!", []string{color.Bold, color.CyanFg})
	cui.Message("Some Message !!", []string{color.Bold, color.CyanBg})

	time.Sleep(3 * time.Second)
	cui.Erase()
}
