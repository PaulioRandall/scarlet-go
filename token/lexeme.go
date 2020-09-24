package token

import (
	"fmt"
)

type Lexeme struct {
	TokenType
	raw  string
	line int
	col  int
}

func New(raw string, tt TokenType, line, col int) Lexeme {
	return Lexeme{
		TokenType: tt,
		raw:       raw,
		line:      line,
		col:       col,
	}
}

func Tok(raw string, tt TokenType) Lexeme {
	return Lexeme{
		TokenType: tt,
		raw:       raw,
	}
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
