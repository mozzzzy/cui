package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of secureInput package.")

	answer, canceled := cui.SecureInput("Please type something and press Enter.")
	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
	} else {
		cui.Info("Answer is \"" + answer + "\".")
		time.Sleep(3 * time.Second)
	}
	cui.Erase()
	return
}
