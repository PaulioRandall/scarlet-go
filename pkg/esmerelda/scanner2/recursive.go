package scanner2

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type lexeme struct {
	ty  TokenType
	tok []rune
}

func (l *lexeme) add(ru rune) {
	l.tok = append(l.tok, ru)
}

func (s *Scanner) next(lex *lexeme) error {

	switch /*ru := s.peekSym();*/ {
	case s.symbol(lex):
		return nil
	}

	return err.New("Unknown symbol", err.Pos(s.line, s.col))
}

func (s *Scanner) symbol(lex *lexeme) bool {

	switch s.peekSym() {
	case '_':
		lex.add(s.nextSym())
		lex.ty = TK_VOID
		return true
	}

	return false
}
