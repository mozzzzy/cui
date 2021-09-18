package question

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v3/color"
	"github.com/mozzzzy/cui/v3/core/element"
	"github.com/mozzzzy/cui/v3/core/elementChain"
)

/*
 * Types
 */

type Question struct {
	elemChain elementChain.ElementChain
}

/*
 * Constants and Package Scope Variables
 */

var (
	Prefix                  string   = "?"
	PrefixColors            []string = []string{color.GreenFg, color.Bold}
	Padding                 string   = " "
	PaddingColors           []string = []string{}
	QuestionColors          []string = []string{}
	AnswerPrefix            string   = "("
	AnswerSuffix            string   = ")"
	AnswerColors            []string = []string{color.CyanFg, color.Bold}
	AnswerPlaceHolder       string   = ""
	AnswerPlaceHolderColors []string = []string{}
	NewLine                 string   = "\r\n"
	NewLineColors           []string = []string{}
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func New(question string) *Question {
	elems := []element.Element{
		// Prefix
		{
			Str:    Prefix,
			Colors: PrefixColors,
		},
		// Padding
		{
			Str:    Padding,
			Colors: PaddingColors,
		},
		// Question
		{
			Str:    question,
			Colors: QuestionColors,
		},
		// Padding
		{
			Str:    Padding,
			Colors: PaddingColors,
		},
		// Answer
		{
			Str:    AnswerPlaceHolder,
			Colors: AnswerPlaceHolderColors,
		},
		// New Line
		{
			Str:    NewLine,
			Colors: NewLineColors,
		},
	}
	elemChain := elementChain.New(elems)

	q := Question{
		elemChain: *elemChain,
	}
	return &q
}

/*
 * Private Methods
 */

func (question Question) getAnswerElemPtr() *element.Element {
	return &question.elemChain.Elems[4]
}

/*
 * Public Methods
 */

func (question Question) Erase() {
	question.elemChain.Erase()
}

func (question *Question) Print() {
	question.elemChain.Print()
}

func (question *Question) SetAnswer(answer string) {
	question.Erase()
	answerElemPtr := question.getAnswerElemPtr()
	answerElemPtr.Str = AnswerPrefix + answer + AnswerSuffix
	answerElemPtr.Colors = AnswerColors
	question.Print()
}
