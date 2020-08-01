package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	script      string
	lineEndings string
	nofmt       bool
	log         bool
	logFile     string
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
		return fmt.Errorf("Expected script filename")
	}

	c.script = args.shift()

	if args.more() {
		return fmt.Errorf("Unexpected argument %q", args.peek())
	}

	return identifyLineEndings(c)
}

func optionArg(c *config, args Arguments) error {

	switch args.peek() {
	case "-nofmt":
		nofmtOption(c, args)

	case "-log":
		return logOption(c, args)

	default:
		return fmt.Errorf("Unexpected option %q", args.peek())
	}

	return nil
}

func nofmtOption(c *config, args Arguments) {
	c.nofmt = true
	args.shift()
}

func logOption(c *config, args Arguments) error {

	if args.count() < 2 {
		return fmt.Errorf("Missing %q folder name", args.peek())
	}

	c.log = true
	args.shift()
	c.logFile = args.shift()

	return nil
}

func identifyLineEndings(c *config) error {

	f, e := os.Open(c.script)
	if e != nil {
		return e
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	ok := s.Scan()

	if s.Err() != nil {
		return s.Err()
	}

	if !ok {
		return nil
	}

	t := string(s.Text())
	if strings.HasSuffix(t, "\r") {
		c.lineEndings = "\r\n"
	} else {
		c.lineEndings = "\n"
	}

	return nil
}
