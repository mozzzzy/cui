package main

import (
	"time"

	"github.com/mozzzzy/cui/v3"
)

func main() {
	cui.Info("This is demo code of table package.")

	cui.Table([][]string{
		{
			"column0", "column1", "column2",
		},
		{
			"data00", "data01", "data02",
		},
		{
			"", "data11", "data12",
		},
		{
			"data20", "", "data22",
		},
		{
			"data30", "data31", "",
		},
		{
			"", "", "data42",
		},
		{
			"", "data51", "",
		},
		{
			"data60", "", "",
		},
		{
			"", "", "",
		},
		{
			"data long string", "", "",
		},
	})
	time.Sleep(3 * time.Second)
	cui.Erase()
}
