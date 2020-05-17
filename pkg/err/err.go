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
		script  = loadScript(scriptFile)
		linePre = digitCount(len(script))
		msg     = e.Error()
		line    = e.Line()
		col     = e.Col()
	)

	if line < 0 || col < 0 {
		fPrintln(w, "[ERROR] %s", msg)

	} else {
		// +1 converts from index to count
		fPrintLines(w, script, linePre, line-LINES_BEFORE, line)
		fPrintLines(w, script, linePre, line, line+1)
		printErrPtr(w, e, linePre)
		fPrintLines(w, script, linePre, line+1, line+1+LINES_AFTER)
	}
}

func digitCount(i int) (n int) {
	for i != 0 {
		i /= 10
		n++
	}
	return
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

func fPrintLines(w io.Writer, script []string, linePre, start, end int) {

	size := len(script)

	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}

	for i := start; i < end; i++ {
		s := craftLine(linePre, i+1, script[i])
		fPrintln(w, s)
	}
}

func craftLine(linePre, lineNum int, txt string) string {

	if lineNum < 1 {
		n := preLen(linePre)
		pre := strings.Repeat(" ", n)
		return fmt.Sprintf("%s%s", pre, txt)
	}

	preSpace := linePre - digitCount(lineNum)
	pre := strings.Repeat(" ", preSpace)
	return fmt.Sprintf("%s%d: %s", pre, lineNum, txt)
}

func preLen(linePre int) int {
	return linePre + 2
}

func printErrPtr(w io.Writer, e Err, linePre int) {

	var (
		msg  = e.Error()
		col  = e.Col()
		size = e.Len()
	)

	if size == 0 {
		size = 1
	}

	// `¯\_(ツ)_/¯`

	pre := strings.Repeat(" ", col)

	s := pre + strings.Repeat(`^`, size)
	s = fmt.Sprintf("%s [%d:%d]", s, col, col+size)
	s = craftLine(linePre, 0, s)
	fPrintln(w, s)

	s = craftLine(linePre, 0, msg)
	fPrintln(w, s)
}
