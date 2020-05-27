package main

import (
	"github.com/mozzzzy/cui"
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
