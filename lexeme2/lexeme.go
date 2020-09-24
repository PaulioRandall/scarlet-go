package lexeme2

import (
	"fmt"
)

type Snippet interface {
	At() (line, start, end int)
}

type Lexeme struct {
	Tok  Token
	Raw  string
	Line int
	Col  int
}

func (lex Lexeme) At() (line, start, end int) {
	return lex.Line, lex.Col, len(lex.Raw)
}

func (lex Lexeme) String() string {
	return fmt.Sprintf("%d:%d %s %q",
		lex.Line,
		lex.Col,
		lex.Tok.String(),
		lex.Raw,
	)
}
