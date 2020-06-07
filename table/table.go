package table

/*
 * Module Dependencies
 */

import (
	"strings"

	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/constants"
	"github.com/mozzzzy/cui/v2/element"
	"github.com/mozzzzy/cui/v2/elementChain"
)

/*
 * Types
 */

type Table struct {
	elemChain elementChain.ElementChain
}

/*
 * Constants and Package Scope Variables
 */

var (
	ColumnStrColors          []string = []string{color.GreenFg, color.Bold}
	StrColors                []string
	LineColors               []string
	VerticalLine             string = "│"
	HorizontalLine           string = "─"
	UpperLeftCorner          string = "╭"
	UpperRightCorner         string = "╮"
	UpperLineBetweenColumns  string = "┬"
	BottomLeftCorner         string = "╰"
	BottomRightCorner        string = "╯"
	BottomLineBetweenColumns string = "┴"
	LeftLineBetweenLines     string = "├"
	RightLineBetweenLines    string = "┤"
	Intersection             string = "┼"
)

/*
 * Private Functions
 */

func getColumnNum(lines [][]string) int {
	columnNumMax := 0
	for _, line := range lines {
		if len(line) > columnNumMax {
			columnNumMax = len(line)
		}
	}
	return columnNumMax
}

func getEachColmunLenMax(lines [][]string, columnNumMax int) []int {
	eachColumnLenMax := make([]int, columnNumMax)
	for _, line := range lines {
		for i, column := range line {
			// +2 is length of left and right paddings
			columnLen := len(column) + 2
			if columnLen > eachColumnLenMax[i] {
				eachColumnLenMax[i] = columnLen
			}
		}
	}
	return eachColumnLenMax
}

func buildLine(left, middle, right string, eachColumnLenMax []int) []element.Element {
	elems := []element.Element{
		// append padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// append left
		{
			Str:    left,
			Colors: LineColors,
		},
	}

	for i, columnLen := range eachColumnLenMax {
		// append "─"
		for ci := 0; ci < columnLen; ci++ {
			elems = append(
				elems,
				element.Element{
					Str:    HorizontalLine,
					Colors: LineColors,
				},
			)
		}
		if i == len(eachColumnLenMax)-1 {
			// append right
			elems = append(
				elems,
				element.Element{
					Str:    right,
					Colors: LineColors,
				},
			)
			break
		}
		// append middle
		elems = append(
			elems,
			element.Element{
				Str:    middle,
				Colors: LineColors,
			},
		)
	}
	elems = append(
		elems,
		element.Element{
			Str:    constants.NewLine,
			Colors: []string{},
		},
	)
	return elems
}

func buildTopLine(eachColumnLenMax []int) []element.Element {
	return buildLine(UpperLeftCorner, UpperLineBetweenColumns, UpperRightCorner, eachColumnLenMax)
}

func buildMiddleLine(eachColumnLenMax []int) []element.Element {
	return buildLine(LeftLineBetweenLines, Intersection, RightLineBetweenLines, eachColumnLenMax)
}

func buildBottomLine(eachColumnLenMax []int) []element.Element {
	return buildLine(BottomLeftCorner, BottomLineBetweenColumns, BottomRightCorner, eachColumnLenMax)
}

func buildColumnDataLine(line []string, eachColumnLenMax []int) []element.Element {
	return buildDataLine(line, ColumnStrColors, eachColumnLenMax)
}

func buildDataLine(line []string, strColors []string, eachColumnLenMax []int) []element.Element {
	elems := []element.Element{
		// padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// "│"
		{
			Str:    VerticalLine,
			Colors: LineColors,
		},
	}

	for i, columnLen := range eachColumnLenMax {
		data := line[i]
		// +2 is length of left and right paddings
		paddingLen := columnLen - (len(data) + 2)
		padding := strings.Repeat(constants.Padding, paddingLen)
		elems = append(
			elems,
			[]element.Element{
				// left padding
				{
					Str:    constants.Padding,
					Colors: constants.PaddingColors,
				},
				// data
				{
					Str:    data,
					Colors: strColors,
				},
				// right padding
				{
					Str:    constants.Padding,
					Colors: constants.PaddingColors,
				},
				// extra padding
				{
					Str:    padding,
					Colors: constants.PaddingColors,
				},
				// "│"
				{
					Str:    VerticalLine,
					Colors: LineColors,
				},
			}...,
		)
	}
	elems = append(
		elems,
		element.Element{
			Str:    constants.NewLine,
			Colors: []string{},
		},
	)
	return elems
}

/*
 * Public Functions
 */

func New(lines [][]string) *Table {
	columnNum := getColumnNum(lines)
	eachColumnLen := getEachColmunLenMax(lines, columnNum)

	/*
			 * table is following format.
			 * +---------+-----+-----+     +-----+-----+     +-----+
			 * | Padding | (╭) | (─) | ... | (┬) | (─) | ... | (╮) |
			 * +---------+-----+-----+     +-----+-----+     +-----+
			 * | Padding | (│) |   (data)  | (│) |   (data)  | (│) |
			 * +---------+-----+-----+     +-----+-----+     +-----+
			 * | Padding | (├) | (─) | ... | (┼) | (─) | ... | (┤) |
			 * +---------+-----+-----+     +-----+-----+     +-----+
		 	 * | Padding | (│) |   (data)  | (│) |   (data)  | (│) |
			 * +---------+-----+-----+     +-----+-----+     +-----+
			 * | Padding | (╰) | (─) | ... | (┴) | (─) | ... | (╯) |
			 * +---------+-----+-----+     +-----+-----+     +-----+
	*/
	elemChain := elementChain.New([]element.Element{})
	// Append top line elems
	elemChain.Elems = append(elemChain.Elems, buildTopLine(eachColumnLen)...)

	// Append elems of data line and line between lines
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			elemChain.Elems = append(elemChain.Elems, buildColumnDataLine(lines[i], eachColumnLen)...)
		} else {
			elemChain.Elems = append(elemChain.Elems, buildDataLine(lines[i], StrColors, eachColumnLen)...)
		}
		if i == len(lines)-1 {
			elemChain.Elems = append(elemChain.Elems, buildBottomLine(eachColumnLen)...)
			break
		}
		elemChain.Elems = append(elemChain.Elems, buildMiddleLine(eachColumnLen)...)
	}

	table := Table{
		elemChain: *elemChain,
	}
	return &table
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (tbl *Table) Print() {
	tbl.elemChain.Print()
}
