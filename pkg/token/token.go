package token

import (
	"fmt"
	"strconv"
)

type Token interface {
	Type() TokenType
	Value() string
	Line() int
	Col() int
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

func NewToken(ty TokenType, v string, line, col int) Token {
	return tok{ty, v, line, col}
}

func ToString(tk Token) string {

	if tk == nil {
		return `NIL-TOKEN`
	}

	if v, ok := tk.(fmt.Stringer); ok {
		return v.String()
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

func PrettyPrint(tks []Token) {

	for _, tk := range tks {
		s := tk.Type().String()
		fmt.Print(s + " ")
	}

	fmt.Println()
}
