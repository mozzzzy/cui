package spinner

/*
 * Module Dependencies
 */

import (
	"sync"
	"time"

	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/constants"
	"github.com/mozzzzy/cui/v2/cursor"
	"github.com/mozzzzy/cui/v2/prefixedMessage"
)

/*
 * Types
 */

type Spinner struct {
	pMsg           prefixedMessage.PrefixedMessage
	state          int
	finished       bool
	succeeded      bool
	serializeMutex sync.Mutex
}

/*
 * Constants and Package Scope Variables
 */

var (
	Spins      []string = []string{"⠙", "⠸", "⠴", "⠦", "⠇", "⠋"}
	SpinColors []string = []string{color.CyanFg, color.Bold}

	StrColors []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *Spinner {
	var state int
	pMsg := prefixedMessage.New(
		Spins[state], SpinColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var spinner Spinner
	spinner.pMsg = *pMsg
	spinner.state = state
	return &spinner
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (spnr Spinner) GetMinX() int {
	return spnr.pMsg.GetMinX()
}

func (spnr Spinner) GetMinY() int {
	return spnr.pMsg.GetMinY()
}

func (spnr Spinner) GetMaxX() int {
	return spnr.pMsg.GetMaxX()
}

func (spnr Spinner) GetMaxY() int {
	return spnr.pMsg.GetMaxY()
}

func (spnr Spinner) GetStartX() int {
	return spnr.pMsg.GetStartX()
}

func (spnr Spinner) GetStartY() int {
	return spnr.pMsg.GetStartY()
}

func (spnr Spinner) GetEndX() int {
	return spnr.pMsg.GetEndX()
}

func (spnr Spinner) GetEndY() int {
	return spnr.pMsg.GetEndY()
}

func (spnr *Spinner) Print() {
	spnr.pMsg.Print()
}

func (spnr *Spinner) Complete() {
	if spnr.finished {
		return
	}
	spnr.succeeded = true
	spnr.finished = true
	spnr.serializeMutex.Lock()
	spnr.serializeMutex.Unlock()
}

func (spnr *Spinner) Failure() {
	if spnr.finished {
		return
	}
	spnr.succeeded = false
	spnr.finished = true
	spnr.serializeMutex.Lock()
	spnr.serializeMutex.Unlock()
}

func (spnr *Spinner) Run() {
	spnr.serializeMutex.Lock()
	spnr.finished = false
	go func() {
		spnr.Print()
		for !spnr.finished {
			time.Sleep(100 * time.Millisecond)
			if spnr.state == len(Spins)-1 {
				spnr.state = 0
			} else {
				spnr.state++
			}
			cursor.MoveCursorTo(spnr.GetStartX(), spnr.GetStartY())
			spnr.pMsg.SetPrefix(Spins[spnr.state], SpinColors)
			spnr.Print()
		}
		if spnr.succeeded {
			spnr.pMsg.SetPrefix(constants.Complete, constants.CompleteColors)
		} else {
			spnr.pMsg.SetPrefix(constants.Failure, constants.FailureColors)
		}
		cursor.MoveCursorTo(spnr.GetStartX(), spnr.GetStartY())
		spnr.Print()
		spnr.serializeMutex.Unlock()
	}()
}
