package err

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// PrintErr prints an error using fmt library.
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
		fPrintln("[ERROR] %s", msg)

	} else {
		// +1 converts from index to count
		fPrintLines(script, line-LINES_BEFORE, line)
		fPrintln("---")
		fPrintLines(script, line, line+1)
		fPrintln("--- [%d:%d->%d] %s", line+1, col, col+length, msg)
		fPrintLines(script, line+1, line+1+LINES_AFTER)
	}

	if cause != nil {
		fPrintln("[CAUSE] %s", cause.Error())
	}
}

func loadScript(file string) []string {

	bytes, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	code := string(bytes)
	strings.ReplaceAll(code, "\r", "")
	return strings.Split(code, "\n")
}

func fPrintln(s string, args ...interface{}) {
	fmt.Printf(s, args...)
	fmt.Println()
}

func fPrintLines(script []string, start, end int) {

	size := len(script)

	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}

	for i := start; i < end; i++ {
		fPrintln("%d: %s", i+1, script[i])
	}
}
