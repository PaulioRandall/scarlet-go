package cmd

import (
	"fmt"
)

type config struct {
	nofmt  bool
	script string
}

func captureConfig(c *config, args Arguments) error {
	captureOptions(c, args)
	return captureScriptFile(c, args)
}

func captureOptions(c *config, args Arguments) error {
	for args.more() && args.isOption() {

		switch {
		case args.accept("-nofmt"):
			c.nofmt = true

		default:
			return fmt.Errorf("Unexpected option %q", args.peek())
		}
	}

	return nil
}

func captureScriptFile(c *config, args Arguments) error {

	if args.empty() {
		return fmt.Errorf("Expected script filename")
	}

	c.script = args.shift()

	if args.more() {
		return fmt.Errorf("Unexpected argument %q", args.peek())
	}

	return nil
}
