package prefixedMessage

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
)

/*
 * Types
 */

type PrefixedMessage struct {
	elemChain elementChain.ElementChain
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(
	prefix string, prefixColors []string,
	padding string, paddingColors []string,
	str string, colors []string) *PrefixedMessage {
	str += constants.NewLine
	prefixElem := element.New(prefix, prefixColors)
	paddingElem := element.New(padding, paddingColors)
	msgElem := element.New(str, colors)
	elemChain := elementChain.New([]element.Element{*prefixElem, *paddingElem, *msgElem})
	var pMsg PrefixedMessage
	pMsg.elemChain = *elemChain
	return &pMsg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg PrefixedMessage) Erase() {
	msg.elemChain.Erase()
}

func (msg PrefixedMessage) GetMinX() int {
	return msg.elemChain.GetMinX()
}

func (msg PrefixedMessage) GetMinY() int {
	return msg.elemChain.GetMinY()
}

func (msg PrefixedMessage) GetMaxX() int {
	return msg.elemChain.GetMaxX()
}

func (msg PrefixedMessage) GetMaxY() int {
	return msg.elemChain.GetMaxY()
}

func (msg PrefixedMessage) GetStartX() int {
	return msg.elemChain.GetStartX()
}

func (msg PrefixedMessage) GetStartY() int {
	return msg.elemChain.GetStartY()
}

func (msg PrefixedMessage) GetEndX() int {
	return msg.elemChain.GetEndX()
}

func (msg PrefixedMessage) GetEndY() int {
	return msg.elemChain.GetEndY()
}

func (msg *PrefixedMessage) Print() {
	msg.elemChain.Print()
}

func (msg PrefixedMessage) SetPrefix(str string, colors []string) {
	msg.elemChain.Elems[0].Str = str
	msg.elemChain.Elems[0].Colors = colors
}
