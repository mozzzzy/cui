package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of noticeMessage package.")

	cui.Notice("Some Notice Message !!")
	cui.Notice("Some Notice Message !!")
	cui.Notice("Some Notice Message !!")
	cui.Notice("Some Notice Message !!")
	cui.Notice("Some Notice Message !!")

	time.Sleep(3 * time.Second)
	cui.Erase()
}
