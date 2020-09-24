package lexeme2

import (
	"fmt"
)

type Lexeme struct {
	TokenType
	raw  string
	line int
	col  int
}

func (l Lexeme) Type() TokenType {
	return l.TokenType
}

func (l Lexeme) Raw() string {
	return l.raw
}

func (l Lexeme) Line() int {
	return l.line
}

func (l Lexeme) Col() int {
	return l.col
}

func (l Lexeme) Len() int {
	return len(l.raw)
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.line,
		lex.col,
		lex.TokenType.String(),
		lex.raw,
	)
}
