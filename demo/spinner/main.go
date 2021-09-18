package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of spinner package.")

	str := "Waiting some operations..."
	spinner0 := cui.Spinner(str + " (success example)")
	time.Sleep(3 * time.Second)
	spinner0.Complete()

	spinner1 := cui.Spinner(str + " (failure example)")
	time.Sleep(3 * time.Second)
	spinner1.Failure()

	time.Sleep(1 * time.Second)
	cui.Erase()
}
