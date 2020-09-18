package checker

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

type checkErr struct {
	error
	lex lexeme.Lexeme
	msg string
}

func newErr(lex *lexeme.Lexeme, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return checkErr{
		lex: *lex,
		msg: msg,
	}
}

func (e checkErr) Error() string {
	return e.msg
}

func (e checkErr) Pos() (lineIdx, colIdx int) {
	return e.lex.Line, e.lex.Col
}

func (e checkErr) Len() int {
	return len(e.lex.Raw)
}

type checkEofErr struct {
	error
	msg string
}

func newEofErr(msg string, args ...interface{}) error {
	return checkEofErr{
		msg: fmt.Sprintf(msg, args...),
	}
}

func (e checkEofErr) Error() string {
	return e.msg
}

func (e checkEofErr) Eof() bool {
	return true
}
