package err

import (
	"github.com/pkg/errors"
)

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type Lexeme interface {
	Line() int
	Col() int
	Size() int
}

type sErr struct {
	msg         string
	sLine, sCol int
	eLine, eCol int
}

func (e sErr) Error() string {
	return e.msg
}

func (e sErr) Begin() (int, int) {
	return e.sLine, e.sCol
}

func (e sErr) End() (int, int) {
	return e.eLine, e.eCol
}

func New(msg string) error {

	e := sErr{
		msg:   msg,
		sLine: -1,
		sCol:  -1,
		eLine: -1,
		eCol:  -1,
	}

	return errors.WithStack(e)
}

func NewByPos(msg string, line, col int) error {

	e := sErr{
		msg:   msg,
		sLine: line,
		sCol:  col,
		eLine: line,
		eCol:  col + 1,
	}

	return errors.WithStack(e)
}

func NewByStr(msg string, line, col, len int) error {

	e := sErr{
		msg:   msg,
		sLine: line,
		sCol:  col,
		eLine: line,
		eCol:  col + len,
	}

	return errors.WithStack(e)
}

func NewByLexeme(msg string, lex Lexeme) error {

	e := sErr{
		msg:   msg,
		sLine: lex.Line(),
		sCol:  lex.Col(),
		eLine: lex.Line(),
		eCol:  lex.Line() + lex.Size(),
	}

	return errors.WithStack(e)
}

func NewBySnippet(msg string, snip Snippet) error {

	e := sErr{
		msg: msg,
	}

	e.sLine, e.sCol = snip.Begin()
	e.eLine, e.eCol = snip.End()

	return errors.WithStack(e)
}
