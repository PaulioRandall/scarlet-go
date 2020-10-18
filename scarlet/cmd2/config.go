package cmd2

import (
	"fmt"
)

type (
	// Command is a structured representation of the program arguments. Its
	// specific type and fields determine the actions to take.
	Command interface {
		cmd()
	}

	// HelpCmd represents a help command.
	HelpCmd struct {
		Item string // Empty string indicates general help
	}

	// BuildCmd represents a build command.
	BuildCmd struct {
		Run    bool
		Script string
		Log    bool
		LogDir string
	}
)

func (c HelpCmd) cmd()  {}
func (c BuildCmd) cmd() {}

// Capture converts the program arguments into a form easy to work with.
func Capture(a *Args) (Command, error) {
	switch {
	case a.Empty(), a.Accept("help"):
		return captureHelpCmd(a)
	case a.Accept("build"), a.Accept("log"):
		return captureBuildCmd(a, false)
	case a.Accept("run"):
		return captureBuildCmd(a, true)
	default:
		return captureHelpCmd(a)
	}
}

func captureHelpCmd(a *Args) (Command, error) {
	c := HelpCmd{}
	if a.More() {
		c.Item = a.Shift()
	}
	if a.More() {
		return nil, fmt.Errorf("Unexpected argument %q", a.Peek())
	}
	return c, nil
}

func captureBuildCmd(a *Args, run bool) (Command, error) {
	c := BuildCmd{Run: run}
	if e := captureOptions(&c, a); e != nil {
		return nil, e
	}
	if e := captureScroll(&c, a); e != nil {
		return nil, e
	}
	if a.More() {
		return nil, fmt.Errorf("Unexpected argument %q", a.Peek())
	}
	return c, nil
}

func captureOptions(c *BuildCmd, a *Args) error {
	for a.More() && a.isOption() {
		switch {
		case a.Accept("-log"):
			c.Log = true
			c.LogDir = a.ShiftDefault("")

		default:
			return fmt.Errorf("Unknown option %q", a.Peek())
		}
	}
	return nil
}

func captureScroll(c *BuildCmd, a *Args) error {
	if a.Empty() {
		return fmt.Errorf("Expected scroll filename")
	}
	c.Script = a.Shift()
	return nil
}
