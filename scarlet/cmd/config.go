package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	script      string
	lineEndings string
	nofmt       bool
	log         bool
	logFile     string
}

func (b Config) logFilename(ext string) string {
	f := filepath.Base(b.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(b.logFile, f+ext)
}

func CaptureConfig(c *Config, args Arguments) error {

	for args.More() {

		if !strings.HasPrefix(args.Peek(), "-") {
			break
		}

		e := optionArg(c, args)
		if e != nil {
			return e
		}
	}

	if args.Empty() {
		return fmt.Errorf("Expected script filename")
	}

	c.script = args.Shift()

	if args.More() {
		return fmt.Errorf("Unexpected argument %q", args.Peek())
	}

	return identifyLineEndings(c)
}

func optionArg(c *Config, args Arguments) error {

	switch args.Peek() {
	case "-nofmt":
		nofmtOption(c, args)

	case "-log":
		return logOption(c, args)

	default:
		return fmt.Errorf("Unexpected option %q", args.Peek())
	}

	return nil
}

func nofmtOption(c *Config, args Arguments) {
	c.nofmt = true
	args.Shift()
}

func logOption(c *Config, args Arguments) error {

	if args.Count() < 2 {
		return fmt.Errorf("Missing %q folder name", args.Peek())
	}

	c.log = true
	args.Shift()
	c.logFile = args.Shift()

	return nil
}

func identifyLineEndings(c *Config) error {

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
