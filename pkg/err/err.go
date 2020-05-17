package err

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Err interface {
	error
	Line() int // index
	Col() int  // index
	Len() int
}

type goErr struct {
	error
}

func (ge goErr) Error() string {
	return ge.error.Error()
}

func (ge goErr) Line() int {
	return 0
}

func (ge goErr) Col() int {
	return 0
}

func (ge goErr) Len() int {
	return 0
}

func Try(f func()) (err Err) {

	func() {
		defer func() {

			switch v := recover().(type) {
			case nil:
				err = nil
			case Err:
				err = v
			case string:
				err = goErr{errors.New(v)}
			case error:
				err = goErr{v}
			default:
				s := `¯\_(ツ)_/¯ Something panicked, but I don't understand the error`
				err = goErr{errors.New(s)}
			}

		}()

		f()
	}()

	return
}

func Print(w io.Writer, e Err, scriptFile string) {

	const (
		LINES_BEFORE = 5
		LINES_AFTER  = 4
	)

	var (
		script = loadScript(scriptFile)
		// `¯\_(ツ)_/¯`
		msg    = e.Error()
		line   = e.Line()
		col    = e.Col()
		length = e.Len()
	)

	if line < 0 || col < 0 {
		fPrintln(w, "[ERROR] %s", msg)

	} else {
		// +1 converts from index to count
		fPrintLines(w, script, line-LINES_BEFORE, line)
		fPrintln(w, "---")
		fPrintLines(w, script, line, line+1)
		fPrintln(w, "--- [%d:%d->%d] %s", line+1, col, col+length, msg)
		fPrintLines(w, script, line+1, line+1+LINES_AFTER)
	}
}

func loadScript(f string) []string {

	bs, e := ioutil.ReadFile(f)
	if e != nil {
		panic(e)
	}

	s := string(bs)
	strings.ReplaceAll(s, "\r", "")
	return strings.Split(s, "\n")
}

func fPrintln(w io.Writer, s string, args ...interface{}) {
	fmt.Fprintf(w, s, args...)
	fmt.Fprintln(w)
}

func fPrintLines(w io.Writer, script []string, start, end int) {

	size := len(script)

	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}

	for i := start; i < end; i++ {
		fPrintln(w, "%d: %s", i+1, script[i])
	}
}
