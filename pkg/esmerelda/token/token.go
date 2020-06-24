package token

import (
	"fmt"
	"strconv"
)

type Token interface {
	fmt.Stringer
	Type() TokenType
	Value() string
	Begin() (int, int)
	End() (int, int)
}

func NewToken(ty TokenType, v string, line, col int) Token {
	return tok{
		ty:       ty,
		val:      v,
		line:     line,
		colStart: col,
		colEnd:   col + len(v),
	}
}

type tok struct {
	ty               TokenType
	val              string
	line             int
	colStart, colEnd int
}

func (tk tok) Type() TokenType {
	return tk.ty
}

func (tk tok) Value() string {
	return tk.val
}

func (tk tok) Begin() (int, int) {
	return tk.line, tk.colStart
}

func (tk tok) End() (int, int) {
	return tk.line, tk.colEnd
}

func (tk tok) String() string {
	return toString(tk)
}

func toString(tk Token) string {

	if tk == nil {
		return `NIL`
	}

	sLine, sCol := tk.Begin()
	eLine, eCol := tk.End()
	s := strconv.QuoteToGraphic(tk.Value())

	// +1 for line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s %s`,
		sLine+1,
		sCol,
		eLine+1,
		eCol,
		tk.Type().String(),
		s,
	)
}
