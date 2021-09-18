package warnMessage

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

type WarnMessage struct {
	pMsg prefixedMessage.PrefixedMessage
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix       string   = " Warn "
	PrefixColors []string = []string{color.YellowBg, color.Bold}
	StrColors    []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *WarnMessage {
	pMsg := prefixedMessage.New(
		Prefix, PrefixColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var wMsg WarnMessage
	wMsg.pMsg = *pMsg
	return &wMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg WarnMessage) Erase() {
	msg.pMsg.Erase()
}

func (msg WarnMessage) GetMinX() int {
	return msg.pMsg.GetMinX()
}

func (msg WarnMessage) GetMinY() int {
	return msg.pMsg.GetMinY()
}

func (msg WarnMessage) GetMaxX() int {
	return msg.pMsg.GetMaxX()
}

func (msg WarnMessage) GetMaxY() int {
	return msg.pMsg.GetMaxY()
}

func (msg WarnMessage) GetStartX() int {
	return msg.pMsg.GetStartX()
}

func (msg WarnMessage) GetStartY() int {
	return msg.pMsg.GetStartY()
}

func (msg WarnMessage) GetEndX() int {
	return msg.pMsg.GetEndX()
}

func (msg WarnMessage) GetEndY() int {
	return msg.pMsg.GetEndY()
}

func (msg WarnMessage) Print() {
	msg.pMsg.Print()
}
