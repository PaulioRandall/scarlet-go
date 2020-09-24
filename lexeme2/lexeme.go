package lexeme2

import (
	"fmt"
)

type lexeme struct {
	TokenType
	raw  string
	line int
	col  int
}

func (l lexeme) Type() TokenType {
	return l.TokenType
}

func (l lexeme) Raw() string {
	return l.raw
}

func (l lexeme) Line() int {
	return l.line
}

func (l lexeme) Col() int {
	return l.col
}

func (l lexeme) Len() int {
	return len(l.raw)
}

func (lex lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.line,
		lex.col,
		lex.TokenType.String(),
		lex.raw,
	)
}
