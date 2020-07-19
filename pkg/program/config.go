package program

import (
	"fmt"
	"path/filepath"
	"strings"
)

type config struct {
	script  string
	nofmt   bool
	log     bool
	logFile string
}

func (b config) logFilename(ext string) string {
	f := filepath.Base(b.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(b.logFile, f+ext)
}

func captureConfig(c *config, args Arguments) error {

	for args.more() {

		if !strings.HasPrefix(args.peek(), "-") {
			break
		}

		e := optionArg(c, args)
		if e != nil {
			return e
		}
	}

	if args.empty() {
		e := fmt.Errorf("Expected script filename")
		return NewGenErr(e)
	}

	c.script = args.take()

	if args.more() {
		e := fmt.Errorf("Unexpected argument %q", args.peek())
		return NewGenErr(e)
	}

	return nil
}

func optionArg(c *config, args Arguments) error {

	switch args.peek() {
	case "-nofmt":
		nofmtOption(c, args)

	case "-log":
		return logOption(c, args)

	default:
		e := fmt.Errorf("Unexpected option %q", args.peek())
		return NewGenErr(e)
	}

	return nil
}

func nofmtOption(c *config, args Arguments) {
	c.nofmt = true
	args.take()
}

func logOption(c *config, args Arguments) error {

	if args.count() < 2 {
		e := fmt.Errorf("Missing %q folder name", args.peek())
		return NewGenErr(e)
	}

	c.log = true
	args.take()
	c.logFile = args.take()

	return nil
}
