package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of question package.")

	q := cui.Question("What language do you like?")
	time.Sleep(1 * time.Second)
	q.SetAnswer("go")
	time.Sleep(3 * time.Second)
	cui.Erase()
	return
}
