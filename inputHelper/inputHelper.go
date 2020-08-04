package inputHelper

/*
 * Module Dependencies
 */

import (
	"bufio"
	"os"
	"os/exec"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Private Functions
 */

/*
 * Public Functions
 */

func SetRaw(start bool) error {
	opt := "raw"
	if !start {
		opt = "-raw"
	}

	cmd := exec.Command("stty", opt)
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func SetNoEcho(start bool) error {
	opt := "-echo"
	if !start {
		opt = "echo"
	}

	cmd := exec.Command("stty", opt)
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func Getch() []rune {
	var runes []rune
	input := bufio.NewReader(os.Stdin)
	for {
		r, _, err := input.ReadRune()
		if err != nil {
			SetRaw(false)
			panic(err)
		}
		runes = append(runes, r)
		if input.Buffered() == 0 {
			break
		}
	}
	return runes
}

/*
 * Private Methods
 */

/*
 * Public Methods
 */
