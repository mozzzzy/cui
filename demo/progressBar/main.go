package main

import (
	"time"
	"github.com/mozzzzy/cui/v2"
)

func main() {
	str := "Waiting some operations..."

	pb0 := cui.ProgressBar(str)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		pb0.ReportProgress(i)
	}

	pb1 := cui.ProgressBar(str)
	for i := 0; i <= 55; i++ {
		time.Sleep(100 * time.Millisecond)
		pb1.ReportProgress(i)
	}
	pb1.Failure()

	time.Sleep(1 * time.Second)
}
