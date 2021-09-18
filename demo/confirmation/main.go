package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of confirmation package.")

	q := "Do you like golang?"
	answer, canceled := cui.Confirmation(q)

	if canceled {
		cui.Warn("Canceled.")
		time.Sleep(1 * time.Second)
		cui.Erase()
		return
	}
	if answer {
		cui.Info("Accepted.")
	} else {
		cui.Info("Rejected.")
	}
	time.Sleep(3 * time.Second)
	cui.Erase()
}
