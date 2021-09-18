package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of errorMessage package.")

	cui.Error("Some Error Message !!")
	cui.Error("Some Error Message !!")
	cui.Error("Some Error Message !!")
	cui.Error("Some Error Message !!")
	cui.Error("Some Error Message !!")

	time.Sleep(3 * time.Second)
	cui.Erase()
}
