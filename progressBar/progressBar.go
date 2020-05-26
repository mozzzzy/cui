package progressBar

/*
 * Module Dependencies
 */

import (
	"strconv"

	"github.com/mozzzzy/cui/color"
	"github.com/mozzzzy/cui/constants"
	"github.com/mozzzzy/cui/cursor"
	"github.com/mozzzzy/cui/element"
	"github.com/mozzzzy/cui/elementChain"
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
	StrColors           []string
	ProgressLen         int      = 75
	Progress            string   = " "
	ProgressColors      []string = []string{color.CyanBg}
	ProgressSpace       string   = " "
	ProgressSpaceColors []string = []string{color.WhiteBg}
	Percent             string   = "%"
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *ProgressBar {
	progressSpaceStr := ""
	for i := 0; i < ProgressLen; i++ {
		progressSpaceStr += ProgressSpace
	}

	/*
	 * progress bar is following format.
	 * +--------+---------+-----+------+
	 * | Prefix | Padding | Str | \r\n |
	 * +--------++----+---+-++--+------+----+---------+------------+------+
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
		// New line
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

func (pb *ProgressBar) setPercentageElem(percentage int) {
	precentageStr := strconv.Itoa(percentage) + Percent
	pb.elemChain.Elems[8].Str = precentageStr
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
	pb.elemChain.Elems[5].Str = progressStr
	pb.elemChain.Elems[6].Str = progressSpaceStr
}

func (pb *ProgressBar) setResultElem() {
	if !pb.finished {
		return
	}
	if pb.succeeded {
		pb.elemChain.Elems[0].Str = constants.Complete
		pb.elemChain.Elems[0].Colors = constants.CompleteColors
	} else {
		pb.elemChain.Elems[0].Str = constants.Failure
		pb.elemChain.Elems[0].Colors = constants.FailureColors
	}
}

/*
 * Public Methods
 */

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
	pb.elemChain.Elems[8].Colors = constants.FailureColors
	pb.elemChain.Elems[9].Str = constants.NewLine
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
		pb.elemChain.Elems[8].Colors = constants.AnswerColors
		pb.elemChain.Elems[9].Str = constants.NewLine
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
