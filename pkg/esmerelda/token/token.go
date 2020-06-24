package token

import (
	"fmt"
	"strconv"
)

type Token interface {
	fmt.Stringer
	Type() TokenType
	Value() string
	Line() int
	Col() int
	Size() int
}

func NewToken(ty TokenType, v string, line, col int) Token {
	return tok{ty, v, line, col}
}

type tok struct {
	ty TokenType
	v  string
	l  int
	c  int
}

func (tk tok) Type() TokenType {
	return tk.ty
}

func (tk tok) Value() string {
	return tk.v
}

func (tk tok) Line() int {
	return tk.l
}

func (tk tok) Col() int {
	return tk.c
}

func (tk tok) Size() int {
	return len(tk.v)
}

func (tk tok) String() string {
	return toString(tk)
}

func toString(tk Token) string {

	if tk == nil {
		return `NIL`
	}

	var s interface{}
	v := tk.Value()
	ty := tk.Type()

	switch ty {
	case TK_STRING, TK_TERMINATOR, TK_NEWLINE, TK_WHITESPACE:
		s = strconv.QuoteToGraphic(v)

	default:
		s = v
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`,
		tk.Line()+1,
		tk.Col(),
		ty.String(),
		s,
	)
}
