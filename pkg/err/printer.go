package err

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// PrintErr prints an error.
func PrintErr(e Err, file string) {

	var (
		LINES_BEFORE = 5
		LINES_AFTER  = 4
		script       = loadScript(file)
		// `¯\_(ツ)_/¯`
		msg    = e.Error()
		cause  = e.Cause()
		line   = e.LineIndex()
		col    = e.ColIndex()
		length = e.Length()
	)

	if line < 0 || col < 0 {
		writeln("[ERROR] %s", msg)

	} else {
		// +1 converts from index to count
		writeLines(script, line-LINES_BEFORE, line)
		writeln("---")
		writeLines(script, line, line+1)
		writeln("--- [%d:%d->%d] %s", line+1, col, col+length, msg)
		writeLines(script, line+1, line+1+LINES_AFTER)
	}

	if cause != nil {
		writeln("[CAUSE] %s", cause.Error())
	}
}

// loadScript loads a file.
func loadScript(file string) []string {

	bytes, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	code := string(bytes)
	strings.ReplaceAll(code, "\r", "")
	return strings.Split(code, "\n")
}

// write writes the specified string to the output.
func write(s string, args ...interface{}) {
	s = fmt.Sprintf(s, args...)
	print(s)
}

// writeln writes the specified string to the output appending a newline
// afterward.
func writeln(s string, args ...interface{}) {
	s = fmt.Sprintf(s, args...)
	println(s)
}

// writeLines prints a specified number of lines from the script.
func writeLines(script []string, start, end int) {

	size := len(script)

	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}

	for i := start; i < end; i++ {
		writeln("%d: %s", i+1, script[i])
	}
}
