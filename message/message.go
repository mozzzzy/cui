package message

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/core/constants"
	"github.com/mozzzzy/cui/v3/core/element"
)

/*
 * Types
 */

type Message struct {
	msgElem element.Element
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

func New(str string, colors []string) *Message {
	str += constants.NewLine
	msgElem := element.New(str, colors)
	var msg Message
	msg.msgElem = *msgElem
	return &msg
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */

func (msg Message) Erase() {
	msg.msgElem.Erase()
}

func (msg Message) GetMinX() int {
	return msg.msgElem.GetMinX()
}

func (msg Message) GetMinY() int {
	return msg.msgElem.GetMinY()
}

func (msg Message) GetMaxX() int {
	return msg.msgElem.GetMaxX()
}

func (msg Message) GetMaxY() int {
	return msg.msgElem.GetMaxY()
}

func (msg Message) GetStartX() int {
	return msg.msgElem.GetStartX()
}

func (msg Message) GetStartY() int {
	return msg.msgElem.GetStartY()
}

func (msg Message) GetEndX() int {
	return msg.msgElem.GetEndX()
}

func (msg Message) GetEndY() int {
	return msg.msgElem.GetEndY()
}

func (msg *Message) Print() {
	msg.msgElem.Print()
}
