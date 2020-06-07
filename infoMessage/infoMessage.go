package infoMessage

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v2/color"
	"github.com/mozzzzy/cui/v2/constants"
	"github.com/mozzzzy/cui/v2/prefixedMessage"
)

/*
 * Types
 */

type InfoMessage struct {
	pMsg prefixedMessage.PrefixedMessage
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix       string   = " Info "
	PrefixColors []string = []string{color.GreenBg, color.Bold}
	StrColors    []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *InfoMessage {
	pMsg := prefixedMessage.New(
		Prefix, PrefixColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var iMsg InfoMessage
	iMsg.pMsg = *pMsg
	return &iMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg InfoMessage) GetMinX() int {
	return msg.pMsg.GetMinX()
}

func (msg InfoMessage) GetMinY() int {
	return msg.pMsg.GetMinY()
}

func (msg InfoMessage) GetMaxX() int {
	return msg.pMsg.GetMaxX()
}

func (msg InfoMessage) GetMaxY() int {
	return msg.pMsg.GetMaxY()
}

func (msg InfoMessage) GetStartX() int {
	return msg.pMsg.GetStartX()
}

func (msg InfoMessage) GetStartY() int {
	return msg.pMsg.GetStartY()
}

func (msg InfoMessage) GetEndX() int {
	return msg.pMsg.GetEndX()
}

func (msg InfoMessage) GetEndY() int {
	return msg.pMsg.GetEndY()
}

func (msg *InfoMessage) Print() {
	msg.pMsg.Print()
}
