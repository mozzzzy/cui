package main

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func main() {
	answer, canceled := cui.SecureInput("Please type something and Enter")
	if canceled {
		cui.Warn("Canceled.")
		return
	}
	cui.Info("Answer is \"" + answer + "\".")
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

