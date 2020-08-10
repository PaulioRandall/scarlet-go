package cmd

import (
	"fmt"
)

type config struct {
	nofmt  bool
	script string
	logDir string
}

func (c *config) captureConfig(args Arguments) error {
	c.captureOptions(args)
	return c.captureScriptFile(args)
}

func (c *config) captureOptions(args Arguments) error {
	for args.more() && args.isOption() {

		switch {
		case args.accept("-nofmt"):
			c.nofmt = true

		case args.accept("-log"):
			c.logDir = args.shiftDefault("")

		default:
			return fmt.Errorf("Unexpected option %q", args.peek())
		}
	}

	return nil
}

func (c *config) captureScriptFile(args Arguments) error {

	if args.empty() {
		return fmt.Errorf("Expected script filename")
	}

	c.script = args.shift()

	if args.more() {
		return fmt.Errorf("Unexpected argument %q", args.peek())
	}

	return nil
}
