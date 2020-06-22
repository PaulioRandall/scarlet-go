package scanner2

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type SymItr interface {
	Next() (rune, bool)
}

type Scanner struct {
	buffer
	line int
	col  int
}

func New(s SymItr) *Scanner {

	b := buffer{
		itr: s,
	}
	b.bufferNext()

	return &Scanner{
		buffer: b,
	}
}

func (s *Scanner) Next() (Token, error) {

	if s.empty() {
		return nil, nil
	}

	lex := lexeme{}
	if e := s.next(&lex); e != nil {
		return nil, e
	}

	return s.tokenise(lex), nil
}

func (s *Scanner) tokenise(lex lexeme) Token {

	if lex.ty == TK_UNDEFINED {
		panic("PROGRAMMERS ERROR! Token type missing")
	}

	val := string(lex.tok)

	tk := NewToken(
		lex.ty,
		val,
		s.line,
		s.col,
	)

	// TODO: update scanner

	return tk
}
