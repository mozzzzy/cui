package noticeMessage

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/color"
	"github.com/mozzzzy/cui/constants"
	"github.com/mozzzzy/cui/prefixedMessage"
)

/*
 * Types
 */

type NoticeMessage struct {
	pMsg prefixedMessage.PrefixedMessage
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix       string   = " Notice "
	PrefixColors []string = []string{color.MagentaBg, color.Bold}
	StrColors    []string
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(str string) *NoticeMessage {
	pMsg := prefixedMessage.New(
		Prefix, PrefixColors, constants.Padding, constants.PaddingColors, str, StrColors)
	var nMsg NoticeMessage
	nMsg.pMsg = *pMsg
	return &nMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg NoticeMessage) GetMinX() int {
	return msg.pMsg.GetMinX()
}

func (msg NoticeMessage) GetMinY() int {
	return msg.pMsg.GetMinY()
}

func (msg NoticeMessage) GetMaxX() int {
	return msg.pMsg.GetMaxX()
}

func (msg NoticeMessage) GetMaxY() int {
	return msg.pMsg.GetMaxY()
}

func (msg NoticeMessage) GetStartX() int {
	return msg.pMsg.GetStartX()
}

func (msg NoticeMessage) GetStartY() int {
	return msg.pMsg.GetStartY()
}

func (msg NoticeMessage) GetEndX() int {
	return msg.pMsg.GetEndX()
}

func (msg NoticeMessage) GetEndY() int {
	return msg.pMsg.GetEndY()
}

func (msg *NoticeMessage) Print() {
	msg.pMsg.Print()
}
