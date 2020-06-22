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

	switch s.peekSym() {
	case '_':
		lex.add(s.nextSym())
		lex.ty = TK_VOID
	}

	return err.New("Unknown symbol", err.Pos(s.line, s.col))
}
