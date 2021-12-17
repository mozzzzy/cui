package progressBar

/*
 * Module Dependencies
 */

import (
	"strconv"

	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/cursor"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
)

/*
 * Types
 */

type ProgressBar struct {
	elemChain elementChain.ElementChain
	finished  bool
	succeeded bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	StrColors               []string
	ProgressLen             int      = 75
	Progress                string   = " "
	ProgressColors          []string = []string{color.CyanBg}
	ProgressFailureColors   []string = []string{color.RedBg}
	ProgressSpace           string   = "."
	ProgressSpaceColors     []string = []string{}
	Percent                 string   = "%"
	PercentageFailureColors []string = []string{color.RedFg}
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *ProgressBar {
	var progressSpaceStr string
	for i := 0; i < ProgressLen; i++ {
		progressSpaceStr += ProgressSpace
	}

	/*
	 * progress bar is following format.
	 * +--------+---------+-----+------+
	 * | Prefix | Padding | Str | \r\n |
	 * +--------++--------+-+---+------+----+---------+------------+------+
	 * | Padding | Progress | ProgressSpace | Padding | Percentage | \r\n |
	 * +---------+----------+---------------+---------+------------+------+
	 */

	elemChain := elementChain.New([]element.Element{
		// Prefix
		{
			Str:    "",
			Colors: []string{},
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// Str
		{
			Str:    str,
			Colors: StrColors,
		},
		// New line
		{
			Str:    constants.NewLine,
			Colors: []string{},
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// Progress
		{
			Str:    "",
			Colors: ProgressColors,
		},
		// Progress space
		{
			Str:    progressSpaceStr,
			Colors: ProgressSpaceColors,
		},
		// Padding
		{
			Str:    constants.Padding,
			Colors: constants.PaddingColors,
		},
		// Percentage
		{
			Str:    "",
			Colors: StrColors,
		},
		// Next line
		{
			Str:    "",
			Colors: []string{},
		},
	})

	bar := ProgressBar{
		elemChain: *elemChain,
	}
	bar.setPercentageElem(0)
	return &bar
}

/*
 * Private Methods
 */

func (pb ProgressBar) getPrefixElemPtr() *element.Element {
	return &(pb.elemChain.Elems[0])
}

func (pb ProgressBar) getProgressElemPtr() *element.Element {
	return &(pb.elemChain.Elems[5])
}

func (pb ProgressBar) getProgressSpaceElemPtr() *element.Element {
	return &(pb.elemChain.Elems[6])
}

func (pb ProgressBar) getPercentageElemPtr() *element.Element {
	return &(pb.elemChain.Elems[8])
}

func (pb ProgressBar) getNextLineElemPtr() *element.Element {
	return &(pb.elemChain.Elems[9])
}

func (pb *ProgressBar) setPercentageElem(percentage int) {
	precentageStr := strconv.Itoa(percentage) + Percent
	pb.getPercentageElemPtr().Str = precentageStr
}

func (pb *ProgressBar) setProgressElem(percentage int) {
	var progressStr, progressSpaceStr string
	progressStrLen := percentage * ProgressLen / 100
	for i := 0; i < progressStrLen; i++ {
		progressStr += Progress
	}
	progressSpaceStrLen := ProgressLen - progressStrLen
	for i := 0; i < progressSpaceStrLen; i++ {
		progressSpaceStr += ProgressSpace
	}
	// Progress elem
	pb.getProgressElemPtr().Str = progressStr
	pb.getProgressSpaceElemPtr().Str = progressSpaceStr
}

func (pb *ProgressBar) setResultElem() {
	if !pb.finished {
		return
	}
	if pb.succeeded {
		pb.getPrefixElemPtr().Str = constants.Complete
		pb.getPrefixElemPtr().Colors = constants.CompleteColors
	} else {
		pb.getPrefixElemPtr().Str = constants.Failure
		pb.getPrefixElemPtr().Colors = constants.FailureColors
	}
}

/*
 * Public Methods
 */

func (pb ProgressBar) Erase() {
	pb.elemChain.Erase()
}

func (pb ProgressBar) GetMinX() int {
	return pb.elemChain.GetMinX()
}

func (pb ProgressBar) GetMinY() int {
	return pb.elemChain.GetMinY()
}

func (pb ProgressBar) GetMaxX() int {
	return pb.elemChain.GetMaxX()
}

func (pb ProgressBar) GetMaxY() int {
	return pb.elemChain.GetMaxY()
}

func (pb ProgressBar) GetStartX() int {
	return pb.elemChain.GetStartX()
}

func (pb ProgressBar) GetStartY() int {
	return pb.elemChain.GetStartY()
}

func (pb ProgressBar) GetEndX() int {
	return pb.elemChain.GetEndX()
}

func (pb ProgressBar) GetEndY() int {
	return pb.elemChain.GetEndY()
}

func (pb *ProgressBar) Print() {
	pb.elemChain.Print()
}

func (pb *ProgressBar) Failure() {
	if pb.finished {
		return
	}
	pb.getProgressElemPtr().Colors = ProgressFailureColors
	pb.getPercentageElemPtr().Colors = PercentageFailureColors
	pb.getNextLineElemPtr().Str = constants.NewLine
	pb.succeeded = false
	pb.finished = true
	pb.setResultElem()
	pb.elemChain.Erase()
	cursor.MoveCursorTo(pb.elemChain.GetStartX(), pb.elemChain.GetStartY())
	pb.Print()
}

func (pb *ProgressBar) ReportProgress(percentage int) {
	if pb.finished {
		return
	}
	if percentage >= 100 {
		percentage = 100
		pb.getPercentageElemPtr().Colors = constants.AnswerColors
		pb.getNextLineElemPtr().Str = constants.NewLine
		pb.finished = true
		pb.succeeded = true
		pb.setResultElem()
	}
	pb.setPercentageElem(percentage)
	pb.setProgressElem(percentage)
	pb.elemChain.Erase()
	cursor.MoveCursorTo(pb.elemChain.GetStartX(), pb.elemChain.GetStartY())
	pb.Print()
}
