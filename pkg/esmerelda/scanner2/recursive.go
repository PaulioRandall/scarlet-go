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

func (l *lexeme) accept(ru rune) bool {

	if l.scn.peekSym() != ru {
		return false
	}

	l.add(ru)
	return true
}

func (l *lexeme) expect(exp rune) error {

	ru := l.scn.peekSym()
	if ru != exp {
		m := fmt.Sprintf("Expected '%v', but got '%v'", exp, ru)
		return err.New(m, err.Pos(l.scn.line, l.scn.col))
	}

	l.add(ru)
	return nil
}

func (l *lexeme) scan() error {

	switch {
	case l.symbol():
		return nil
	}

	return err.New("Unknown symbol", err.Pos(l.scn.line, l.scn.col))
}

func (l *lexeme) symbol() bool {

	switch {
	case l.accept('_'):
		l.ty = TK_VOID

	default:
		return false
	}

	return true
}
