package err

import (
	"github.com/pkg/errors"
)

type Err interface {
	error
	Line() int // DEAD
	Col() int  // DEAD
	Len() int  // DEAD
	Begin() (int, int)
	End() (int, int)
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type sErr struct {
	msg         string
	line        int
	col         int
	len         int
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

// DEAD
func (e sErr) Line() int {
	return e.line
}

// DEAD
func (e sErr) Col() int {
	return e.col
}

// DEAD
func (e sErr) Len() int {
	return e.len
}

type Option func(*sErr)

// DEAD
type Lexeme interface {
	Value() string
	Line() int // DEAD
	Col() int  // DEAD
}

// DEAD
func Panic(msg string, ops ...Option) {
	er := New(msg, ops...)
	panic(er)
}

func New(msg string, ops ...Option) error {
	e := error(newErr(msg, ops...))
	return errors.WithStack(e)
}

func newErr(msg string, ops ...Option) Err {

	s := sErr{
		msg:  msg,
		line: -1,
		col:  -1,
	}

	applyOptions(&s, ops...)
	return s
}

func Wrap(e error, ops ...Option) error {
	e = wrap(e, ops...)
	return errors.WithStack(e)
}

func wrap(e error, ops ...Option) Err {

	s := sErr{
		msg:  e.Error(),
		line: -1,
		col:  -1,
	}

	applyOptions(&s, ops...)
	return s
}

func applyOptions(e *sErr, ops ...Option) {
	for _, op := range ops {
		op(e)
	}
}

func Pos(line, col int) Option {
	return func(s *sErr) {
		s.line = line
		s.col = col
	}
}

func Len(len int) Option {
	return func(s *sErr) {
		s.len = len
	}
}

func At(lex Lexeme) Option {
	l, c, ln := lex.Line(), lex.Col(), len(lex.Value())

	return func(s *sErr) {
		s.line = l
		s.col = c
		s.len = ln
	}
}

func After(lex Lexeme) Option {
	l := lex.Line()
	c := lex.Col() + len(lex.Value())

	return func(s *sErr) {
		s.line = l
		s.col = c
		s.len = 0
	}
}
