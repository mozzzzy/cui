package main

import (
	"time"
	"github.com/mozzzzy/cui/v2"
)

func main() {
	str := "Waiting some operations..."
	spinner0 := cui.Spinner(str)
	time.Sleep(3 * time.Second)
	spinner0.Complete()

	spinner1 := cui.Spinner(str)
	time.Sleep(3 * time.Second)
	spinner1.Failure()
}
