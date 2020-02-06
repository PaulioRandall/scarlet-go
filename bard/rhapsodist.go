package bard

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// rhapsodist is a Bard implementation that outputs information with reference
// to the source code.
type rhapsodist struct {
	Bard
	file string
}

// NewRhapsodist creates a new Bard that outputs information with reference to
// the specified source file.
func NewRhapsodist(file string) Bard {
	return rhapsodist{
		file: file,
	}
}

// CatchNightmare satisfies the Bard interface.
func (r rhapsodist) CatchNightmare(f func()) {

	isPanic := true

	defer func() {

		if !isPanic {
			return
		}

		e := recover()
		r.printError(e)
	}()

	f()
	isPanic = false
}

// write writes the specified string to the output.
func (r rhapsodist) write(s string, args ...interface{}) {
	s = fmt.Sprintf(s, args...)
	print(s)
}

// writeln writes the specified string to the output appending a newline
// afterward.
func (r rhapsodist) writeln(s string, args ...interface{}) {
	s = fmt.Sprintf(s, args...)
	println(s)
}

// printError priunts an error.
func (r rhapsodist) printError(e interface{}) {

	var (
		LINES_BEFORE = 4
		LINES_AFTER  = 3
	)

	printNightmare := func(e error, at string, line, col int) {

		if line < 0 || col < 0 {
			r.writeln("[ERROR] %s", e.Error())
			goto WRITE_CAUSE
		}

		// +1 converts from index to count
		r.writeFileLines(line-LINES_BEFORE, line)
		r.writeln("---")
		r.writeFileLines(line, line+1)
		r.writeln("--- %s", e.Error())
		r.writeFileLines(line+1, line+1+LINES_AFTER)

	WRITE_CAUSE:
		if n, ok := e.(Nightmare); ok && n.cause != nil {
			r.writeln("[CAUSE] %s", n.cause.Error())
		}
	}

	switch err := e.(type) {
	case Horror:
		printNightmare(err, err.tk.String(), err.tk.Line, err.tk.Col)
	case Terror:
		printNightmare(err, "%v:%v", err.line, err.col)
	case Nightmare:
		printNightmare(err, `¯\_(ツ)_/¯`, -1, -1)
	case error:
		printNightmare(err, `¯\_(ツ)_/¯`, -1, -1)
	default:
		panic("SANITY CHECK!" +
			" A panic was generated but the content was not an error of any sort",
		)
	}
}

// writeFileLines prints a specified number of lines from the source file.
func (r rhapsodist) writeFileLines(start, end int) {

	bytes, e := ioutil.ReadFile(r.file)
	if e != nil {
		panic(e)
	}

	code := string(bytes)
	lines := strings.Split(code, "\n")
	size := len(lines)

	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}

	for i := start; i < end; i++ {
		r.writeln("%d: %s", i+1, lines[i])
	}
}
