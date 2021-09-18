package checkableTable

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/checkbox"
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
)

/*
 * Types
 */

type CheckableTable struct {
	tableHeaderElemChain elementChain.ElementChain
	checkbox             checkbox.Checkbox
}

/*
 * Constants and Package Scope Variables
 */

var (
	Line                     string   = "-"
	LineColors               []string = []string{}
)

/*
 * Private Functions
 */

func getEachColumnMaxLens(rows [][]string) []int {
	rowNum := len(rows)
	columnNum := len(rows[0])
	var columnMaxLens []int
	for columnIndex := 0; columnIndex < columnNum; columnIndex++ {
		columnMaxLens = append(columnMaxLens, 0)
		for rowIndex := 0; rowIndex < rowNum; rowIndex++ {
			if columnMaxLens[columnIndex] < len(rows[rowIndex][columnIndex]) {
				columnMaxLens[columnIndex] = len(rows[rowIndex][columnIndex])
			}
		}
	}
	return columnMaxLens
}

func getMaxLen(strs []string) int {
	var maxLen int
	for _, str := range strs {
		if maxLen < len(str) {
			maxLen = len(str)
		}
	}
	return maxLen
}

/*
 * Public Functions
 */

func New(rows [][]string) *CheckableTable {
	// Get each column's max length
	eachColumnMaxLens := getEachColumnMaxLens(rows)
	// Append paddings to each column
	var rowsWithPadding [][]string
	for rowIndex, row := range rows {
		rowsWithPadding = append(rowsWithPadding, []string{})
		for columnIndex, column := range row {
			// Add padding until the column length same with the max one
			for {
				if len(column) >= eachColumnMaxLens[columnIndex] {
					break
				}
				column += constants.Padding
			}
			rowsWithPadding[rowIndex] = append(rowsWithPadding[rowIndex], column)
		}
	}

	/*
	 * checkableTable is following format.
	 * +---------+---------+---------------+------+--+-----+---------------+----------+
	 * | Padding | Padding | ColumnTitle 0 | Padding | ... | ColumnTitle N | NextLine |
	 * +---------+---------+---------------+------+--+-----+---------------+----------+
	 * | ------------------------------ Line ----------------------------- | NextLine |
	 * +---------------+---------+----------+---------+---------+----------+----------+
	 * |  ChosePrefix  | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	 * +---------------+---------+----------+---------+---------+----------+----------+
	 * | NoChosePrefix | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	 * +---------------+---------+----------+---------+---------+----------+----------+
	 * |                                    ...                                       |
	 * +---------------+---------+----------+---------+---------+----------+----------+
	 * | NoChosePrefix | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	 * +---------------+---------+----------+---------+---------+----------+----------+
	 */

	// Create table header ElemChain
	// +---------+---------+---------------+---------+-----+---------------+----------+
	// | Padding | Padding | ColumnTitle 0 | Padding | ... | ColumnTitle N | NextLine |
	// +---------+---------+---------------+---------+-----+---------------+----------+
	// | ------------------------------ Line ----------------------------- | NextLine |
	// +-------------------------------------------------------------------+----------+
	//
	chosePrefixLen := len(checkbox.ChosePrefix)
	chosePrefixPadding := ""
	for i := 0; i < chosePrefixLen; i += len(constants.Padding) {
		chosePrefixPadding += constants.Padding
	}

	tableHeaderElems := []element.Element{
		// ChosePrefixPadding
		{
			Str:    chosePrefixPadding,
			Colors: constants.PaddingColors,
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
	}

	// Column Titles and Paddings
	columnNum := len(rows[0])
	for columnIndex, column := range rowsWithPadding[0] {
		// Column Title
		tableHeaderElems = append(
			tableHeaderElems,
			element.Element{
				Str:    column,
				Colors: checkbox.NoChoseColors,
			})
		if columnIndex < columnNum-1 {
			// Padding
			tableHeaderElems = append(
				tableHeaderElems,
				element.Element{
					Str:    constants.Padding,
					Colors: constants.PaddingColors,
				})
		}
	}
	// NextLine
	tableHeaderElems = append(
		tableHeaderElems,
		element.Element{
			Str:    constants.NewLine,
			Colors: []string{},
		})
	// Line
	lineLen := len(chosePrefixPadding) + len(constants.Padding)
	for _, columnMaxLen := range eachColumnMaxLens {
		lineLen += columnMaxLen
		lineLen += len(constants.Padding)
	}
	var line string
	for i := 0; i < lineLen; i++ {
		line += Line
	}
	tableHeaderElems = append(
		tableHeaderElems,
		element.Element{
			Str:    line,
			Colors: LineColors,
		})
	// NextLine
	tableHeaderElems = append(
		tableHeaderElems,
		element.Element{
			Str:    constants.NewLine,
			Colors: []string{},
		})

	tableHeaderElemChain := elementChain.New(tableHeaderElems)

	// Create choices ElemChain
	// +---------------+---------+----------+---------+---------+----------+----------+
	// |  ChosePrefix  | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	// +---------------+---------+----------+---------+---------+----------+----------+
	// | NoChosePrefix | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	// +---------------+---------+----------+---------+---------+----------+----------+
	// |                                    ...                                       |
	// +---------------+---------+----------+---------+---------+----------+----------+
	// | NoChosePrefix | Padding | Column 0 | Padding |   ...   | Column N | NextLine |
	// +---------------+---------+----------+---------+---------+----------+----------+
	//
	var choices []string
	for rowIndex, row := range rowsWithPadding {
		// First row is used to column title
		if rowIndex == 0 {
			continue
		}
		var choice string

		for columnIndex, column := range row {
			// Column
			choice += column
			if columnIndex < len(row)-1 {
				// Padding
				choice += constants.Padding
			}
		}
		choices = append(choices, choice)
	}
	checkbox := checkbox.New(choices)

	checkableTable := CheckableTable{
		tableHeaderElemChain: *tableHeaderElemChain,
		checkbox:     *checkbox,
	}
	return &checkableTable
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (checkableTable *CheckableTable) Ask() ([]int, bool, *CheckableTable) {
	checkableTable.tableHeaderElemChain.Print()
	answers, canceled, _ := checkableTable.checkbox.Ask()
	// The first row is used to column title.
	for i := 0; i < len(answers); i++ {
		answers[i]++
	}
	return answers, canceled, checkableTable
}

func (checkableTable *CheckableTable) Erase() {
	checkableTable.tableHeaderElemChain.Erase()
	checkableTable.checkbox.Erase()
}
