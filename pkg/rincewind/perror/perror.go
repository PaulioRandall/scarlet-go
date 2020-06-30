package perror

import (
	"fmt"

	"github.com/pkg/errors"
)

func ProgPanic(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic("PROGRAMMERS ERROR! " + msg)
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type perr struct {
	msg                 string
	lineBegin, colBegin int
	lineEnd, colEnd     int
}

func (e perr) Error() string {
	return e.msg
}

func (e perr) Begin() (int, int) {
	return e.lineBegin, e.colBegin
}

func (e perr) End() (int, int) {
	return e.lineEnd, e.colEnd
}

func New(msg string) error {

	e := perr{
		msg:       msg,
		lineBegin: -1,
		colBegin:  -1,
		lineEnd:   -1,
		colEnd:    -1,
	}

	return errors.WithStack(e)
}

func NewByPos(msg string, line, col int) error {

	e := perr{
		msg:       msg,
		lineBegin: line,
		colBegin:  col,
		lineEnd:   line,
		colEnd:    col + 1,
	}

	return errors.WithStack(e)
}

func NewByStr(msg string, line, col, len int) error {

	e := perr{
		msg:       msg,
		lineBegin: line,
		colBegin:  col,
		lineEnd:   line,
		colEnd:    col + len,
	}

	return errors.WithStack(e)
}

func NewBySnippet(msg string, snip Snippet) error {

	e := perr{
		msg: msg,
	}

	e.lineBegin, e.colBegin = snip.Begin()
	e.lineEnd, e.colEnd = snip.End()

	return errors.WithStack(e)
}

func NewAfterSnippet(msg string, snip Snippet) error {

	e := perr{
		msg:     msg,
		lineEnd: -1,
		colEnd:  -1,
	}

	e.lineBegin, e.colBegin = snip.End()
	return errors.WithStack(e)
}
