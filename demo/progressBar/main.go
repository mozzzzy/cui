package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of progressBar package.")

	str := "Waiting some operations..."

	pb0 := cui.ProgressBar(str + " (success example)")
	for i := 0; i <= 100; i++ {
		time.Sleep(50 * time.Millisecond)
		pb0.ReportProgress(i)
	}

	pb1 := cui.ProgressBar(str + " (failure example)")
	for i := 0; i <= 55; i++ {
		time.Sleep(50 * time.Millisecond)
		pb1.ReportProgress(i)
	}
	pb1.Failure()

	time.Sleep(3 * time.Second)
	cui.Erase()
}
