package constants

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/cui/v2/color"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

var (
	QuestionPrefix       string   = "?"
	QuestionPrefixColors []string = []string{color.GreenFg, color.Bold}
	QuestionSuffix       string   = ": "
	QuestionColors       []string = []string{}

	Padding       string   = " "
	PaddingColors []string = []string{}

	PointerColors      []string = []string{color.CyanFg, color.Bold}
	PointerSpaceColors []string = []string{}

	AnswerColors     []string = []string{color.CyanFg, color.Bold}
	OpenParenthesis  string   = "("
	CloseParenthesis string   = ")"

	NewLine string = "\r\n"

	UpArrow    string = "\x1b\x5b\x41"
	DownArrow  string = "\x1b\x5b\x42"
	RightArrow string = "\x1b\x5b\x43"
	LeftArrow  string = "\x1b\x5b\x44"
	Enter      string = "\r"
	CtrlC      string = "\x03"
	Delete     string = "\x7f"

	InputColors []string = []string{}

	Complete       string   = "✔"
	CompleteColors []string = []string{color.GreenFg, color.Bold}
	Failure        string   = "✖"
	FailureColors  []string = []string{color.RedFg, color.Bold}
)

/*
 * Private Functions
 */

/*
 * Public Functions
 */

/*
 * Private Methods
 */

/*
 * Public Methods
 */
