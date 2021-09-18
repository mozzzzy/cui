package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of debugMessage package.")

	cui.Debug("Some Debug Message !!")
	cui.Debug("Some Debug Message !!")
	cui.Debug("Some Debug Message !!")
	cui.Debug("Some Debug Message !!")
	cui.Debug("Some Debug Message !!")

	time.Sleep(3 * time.Second)
	cui.Erase()
}
