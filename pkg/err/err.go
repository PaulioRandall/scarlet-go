package err

import (
	"errors"
	"fmt"
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

// TODO: Write to io.writier
func Print(e Err, file string) {

	const (
		LINES_BEFORE = 5
		LINES_AFTER  = 4
	)

	var (
		script = loadScript(file)
		// `¯\_(ツ)_/¯`
		msg    = e.Error()
		line   = e.Line()
		col    = e.Col()
		length = e.Len()
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
