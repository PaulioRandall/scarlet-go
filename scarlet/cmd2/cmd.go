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

	// HelpCmd is used to present help text.
	HelpCmd struct {
		Item string // Empty string indicates general help
	}

	// BuildCmd is used to scan, sanitise, parse, then compile a scroll.
	BuildCmd struct {
		Scroll string
	}

	// LogCmd is used to build but the output of each stage as a separate
	// file.
	LogCmd struct {
		BuildCmd
		Dir string // Empty string represents the current working directory
	}

	// RunCmd is used to when building then running in one go.
	RunCmd struct {
		BuildCmd
	}
)

func (c HelpCmd) cmd()  {}
func (c BuildCmd) cmd() {}
func (c LogCmd) cmd()   {}
func (c RunCmd) cmd()   {}

func (c LogCmd) log() {}
func (c RunCmd) run() {}

// Capture converts the program arguments into a form easy to work with.
func Capture(a *Args) (Command, error) {
	switch {
	case a.Empty(), a.Accept("help"): // help [<item>]
		c := HelpCmd{}
		e := captureHelpCmd(&c, a)
		return c, e

	case a.Accept("build"): // build <scroll>
		c := BuildCmd{}
		e := captureBuildCmd(&c, a)
		return c, e

	case a.Accept("log"): // log [-d <dir>] <scroll>
		c := LogCmd{}
		e := captureLogCmd(&c, a)
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
	if a.More() {
		c.Item = a.Shift()
	}
	return expectEndOfArgs(a)
}

func captureBuildCmd(c *BuildCmd, a *Args) error {
	if e := captureScroll(c, a); e != nil {
		return e
	}
	return expectEndOfArgs(a)
}

func captureLogCmd(c *LogCmd, a *Args) error {
	if a.Accept("-d") {
		if a.Empty() {
			return fmt.Errorf("Expected log directory")
		}
		c.Dir = a.Shift()
	}
	if e := captureScroll(&c.BuildCmd, a); e != nil {
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
	if a.Empty() {
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
