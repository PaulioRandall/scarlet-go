package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
)

type config struct {
	nofmt  bool
	script string
	logDir string
}

func (c *config) logFilename(ext string) string {
	f := filepath.Base(c.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(c.logDir, f+ext)
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
