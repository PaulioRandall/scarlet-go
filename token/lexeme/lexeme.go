package lexeme

import (
	"fmt"
)

type Lexeme struct {
	TokenType
	val  string
	line int
	col  int
}

func New(val string, tt TokenType, line, col int) Lexeme {
	return Lexeme{
		TokenType: tt,
		val:       val,
		line:      line,
		col:       col,
	}
}

func Tok(val string, tt TokenType) Lexeme {
	return Lexeme{
		TokenType: tt,
		val:       val,
	}
}

func (l Lexeme) Type() TokenType {
	return l.TokenType
}

func (l Lexeme) Val() string {
	return l.val
}

func (l Lexeme) Line() int {
	return l.line
}

func (l Lexeme) Col() int {
	return l.col
}

func (l Lexeme) Len() int {
	return len(l.val)
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.line,
		lex.col,
		lex.TokenType.String(),
		lex.val,
	)
}
