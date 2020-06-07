package main

import (
	"github.com/mozzzzy/cui/v2"
)

func main() {
	cui.Table([][]string{
		{
			"column0",
			"column1",
			"column2",
		},
		{
			"data00",
			"data01",
			"data02",
		},
		{
			"",
			"data11",
			"data12",
		},
		{
			"data20",
			"",
			"data22",
		},
		{
			"data30",
			"data31",
			"",
		},
		{
			"",
			"",
			"data42",
		},
		{
			"",
			"data51",
			"",
		},
		{
			"data60",
			"",
			"",
		},
		{
			"",
			"",
			"",
		},
		{
			"data long string",
			"",
			"",
		},

	})
}
