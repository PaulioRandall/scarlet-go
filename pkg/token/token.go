package token

import (
	"fmt"
	"strconv"
)

type Token interface {
	Morpheme() Morpheme
	Value() string
	Line() int
	Col() int
}

type tok struct {
	m Morpheme
	v string
	l int
	c int
}

func (tk tok) Morpheme() Morpheme {
	return tk.m
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

func NewToken(m Morpheme, v string, line, col int) Token {
	return tok{m, v, line, col}
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
	m := tk.Morpheme()

	switch m {
	case TEMPLATE, TERMINATOR, NEWLINE, WHITESPACE:
		s = strconv.QuoteToGraphic(v)

	case STRING:
		s = "`" + v + "`"

	default:
		s = v
	}

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %s %v`,
		tk.Line()+1,
		tk.Col(),
		m.String(),
		s,
	)
}

func PrettyPrint(tks []Token) {

	for _, tk := range tks {
		s := tk.Morpheme().String()
		fmt.Print(s + " ")
	}

	fmt.Println()
}
