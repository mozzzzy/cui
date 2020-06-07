package main

import (
	"github.com/mozzzzy/cui/v2"
)

func main() {
	q := "Do you like golang?"

	answer, canceled := cui.Confirmation(q)

	if canceled {
		cui.Warn("Canceled.")
		return
	}

	if answer {
		cui.Info("Accepted.")
	} else {
		cui.Warn("Rejected.")
	}
}
