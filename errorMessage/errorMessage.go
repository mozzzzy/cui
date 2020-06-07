package errorMessage

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

type ErrorMessage struct {
	pMsg prefixedMessage.PrefixedMessage
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix       string   = " Error "
	PrefixColors []string = []string{color.RedBg, color.Bold}
	StrColors    []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *ErrorMessage {
	pMsg := prefixedMessage.New(
		Prefix, PrefixColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var eMsg ErrorMessage
	eMsg.pMsg = *pMsg
	return &eMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg ErrorMessage) GetMinX() int {
	return msg.pMsg.GetMinX()
}

func (msg ErrorMessage) GetMinY() int {
	return msg.pMsg.GetMinY()
}

func (msg ErrorMessage) GetMaxX() int {
	return msg.pMsg.GetMaxX()
}

func (msg ErrorMessage) GetMaxY() int {
	return msg.pMsg.GetMaxY()
}

func (msg ErrorMessage) GetStartX() int {
	return msg.pMsg.GetStartX()
}

func (msg ErrorMessage) GetStartY() int {
	return msg.pMsg.GetStartY()
}

func (msg ErrorMessage) GetEndX() int {
	return msg.pMsg.GetEndX()
}

func (msg ErrorMessage) GetEndY() int {
	return msg.pMsg.GetEndY()
}

func (msg *ErrorMessage) Print() {
	msg.pMsg.Print()
}
