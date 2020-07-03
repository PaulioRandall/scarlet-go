package perror

import (
	"fmt"

	"github.com/pkg/errors"
)

func ProgPanic(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	panic("PROGRAMMERS ERROR! " + msg)
}

type Error interface {
	error
	Snippet
	Code() string
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type perr struct {
	msg                 string
	lineBegin, colBegin int
	lineEnd, colEnd     int
	code                string
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

func (e perr) Code() string {
	return e.code
}

func New(code string, msg string) error {

	e := perr{
		msg:       msg,
		lineBegin: -1,
		colBegin:  -1,
		lineEnd:   -1,
		colEnd:    -1,
		code:      code,
	}

	return errors.WithStack(e)
}

func NewByPos(code string, msg string, line, col int) error {

	e := perr{
		msg:       msg,
		lineBegin: line,
		colBegin:  col,
		lineEnd:   line,
		colEnd:    col + 1,
		code:      code,
	}

	return errors.WithStack(e)
}

func NewByStr(code string, msg string, line, col, len int) error {

	e := perr{
		msg:       msg,
		lineBegin: line,
		colBegin:  col,
		lineEnd:   line,
		colEnd:    col + len,
		code:      code,
	}

	return errors.WithStack(e)
}

func NewBySnippet(code string, msg string, snip Snippet) error {

	if snip == nil {
		return New(code, msg)
	}

	e := perr{
		msg:  msg,
		code: code,
	}

	e.lineBegin, e.colBegin = snip.Begin()
	e.lineEnd, e.colEnd = snip.End()

	return errors.WithStack(e)
}

func NewAfterSnippet(code string, msg string, snip Snippet) error {

	e := perr{
		msg:     msg,
		lineEnd: -1,
		colEnd:  -1,
		code:    code,
	}

	e.lineBegin, e.colBegin = snip.End()
	return errors.WithStack(e)
}

func Unwrap(e error) Error {
	err := (Error)(nil)
	errors.As(e, &err)
	return err
}
