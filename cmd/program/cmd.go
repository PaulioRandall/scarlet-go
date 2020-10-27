package program

import (
	"fmt"
)

type (
	// Command is a structured representation of the program arguments. Its
	// specific type and fields determine the actions to take.
	Command interface {
		cmd()
	}

	// HelpCmd is used to present help text.
	HelpCmd struct {
	}

	// BuildCmd is used to scan, sanitise, parse, then compile a scroll.
	BuildCmd struct {
		Scroll string
	}

	// RunCmd is used to when building then running in one go.
	RunCmd struct {
		BuildCmd BuildCmd
	}
)

func (c HelpCmd) cmd()  {}
func (c BuildCmd) cmd() {}
func (c RunCmd) cmd()   {}

// Capture converts the program arguments into a form easy to work with.
func Capture(a *Args) (Command, error) {
	switch {
	case !a.More(), a.Accept("help"): // help [<item>]
		c := HelpCmd{}
		e := captureHelpCmd(&c, a)
		return c, e

	case a.Accept("build"): // build <scroll>
		c := BuildCmd{}
		e := captureBuildCmd(&c, a)
		return c, e

	case a.Accept("run"): // run <scroll>
		c := RunCmd{}
		e := captureBuildCmd(&c.BuildCmd, a)
		return c, e

	default:
		return HelpCmd{}, nil
	}
}

func captureHelpCmd(c *HelpCmd, a *Args) error {
	return expectEndOfArgs(a)
}

func captureBuildCmd(c *BuildCmd, a *Args) error {
	if e := captureScroll(c, a); e != nil {
		return e
	}
	return expectEndOfArgs(a)
}

func captureRunCmd(c *RunCmd, a *Args) error {
	if e := captureScroll(&c.BuildCmd, a); e != nil {
		return e
	}
	return expectEndOfArgs(a)
}

func captureScroll(c *BuildCmd, a *Args) error {
	if !a.More() {
		return fmt.Errorf("Expected scroll filename")
	}
	c.Scroll = a.Shift()
	return nil
}

func expectEndOfArgs(a *Args) error {
	if a.More() {
		return fmt.Errorf("Unexpected argument %q", a.Peek())
	}
	return nil
}
