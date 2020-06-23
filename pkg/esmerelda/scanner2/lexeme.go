package scanner2

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type lexeme struct {
	scn *scanner
	ty  TokenType
	tok []rune
}

func (l *lexeme) add(ru rune) {
	l.tok = append(l.tok, ru)
}

func (l *lexeme) get() rune {
	return l.scn.peekSym()
}

func (l *lexeme) match(ru rune) bool {

	if l.scn.peekSym() != ru {
		return false
	}

	return true
}

func (l *lexeme) accept(ru rune) bool {

	if l.match(ru) {
		l.add(l.scn.nextSym())
		return true
	}

	return false
}

func (l *lexeme) expect(exp rune) error {

	if l.accept(exp) {
		return nil
	}

	m := fmt.Sprintf("Expected %q, but got %q", exp, l.scn.peekSym())
	return err.New(m, err.Pos(l.scn.line, l.scn.col))
}

func (l *lexeme) next() rune {
	ru := l.scn.nextSym()
	l.add(ru)
	return ru
}
