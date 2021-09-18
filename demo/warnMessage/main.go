package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of infoMessage package.")

	cui.Warn("Some Warn Message !!")
	cui.Warn("Some Warn Message !!")
	cui.Warn("Some Warn Message !!")
	cui.Warn("Some Warn Message !!")
	cui.Warn("Some Warn Message !!")

	time.Sleep(3 * time.Second)
	cui.Erase()
}
