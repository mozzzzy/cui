package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of infoMessage package.")

	cui.Info("Some Info Message !!")
	cui.Info("Some Info Message !!")
	cui.Info("Some Info Message !!")
	cui.Info("Some Info Message !!")
	cui.Info("Some Info Message !!")

	time.Sleep(3 * time.Second)
	cui.Erase()
}
