package debugMessage

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/prefixedMessage"
)

/*
 * Types
 */

type DebugMessage struct {
	pMsg prefixedMessage.PrefixedMessage
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix       string   = " Debug "
	PrefixColors []string = []string{color.BlueBg, color.Bold}
	StrColors    []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *DebugMessage {
	pMsg := prefixedMessage.New(
		Prefix, PrefixColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var nMsg DebugMessage
	nMsg.pMsg = *pMsg
	return &nMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg DebugMessage) Erase() {
	msg.pMsg.Erase()
}

func (msg DebugMessage) GetMinX() int {
	return msg.pMsg.GetMinX()
}

func (msg DebugMessage) GetMinY() int {
	return msg.pMsg.GetMinY()
}

func (msg DebugMessage) GetMaxX() int {
	return msg.pMsg.GetMaxX()
}

func (msg DebugMessage) GetMaxY() int {
	return msg.pMsg.GetMaxY()
}

func (msg DebugMessage) GetStartX() int {
	return msg.pMsg.GetStartX()
}

func (msg DebugMessage) GetStartY() int {
	return msg.pMsg.GetStartY()
}

func (msg DebugMessage) GetEndX() int {
	return msg.pMsg.GetEndX()
}

func (msg DebugMessage) GetEndY() int {
	return msg.pMsg.GetEndY()
}

func (msg *DebugMessage) Print() {
	msg.pMsg.Print()
}
