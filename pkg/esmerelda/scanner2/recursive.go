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

	ru := l.scn.peekSym()
	if l.accept(exp) {
		m := fmt.Sprintf("Expected '%v', but got '%v'", exp, ru)
		return err.New(m, err.Pos(l.scn.line, l.scn.col))
	}

	return nil
}

func (l *lexeme) scan() error {

	switch {
	case l.accept('\n'):
		l.ty = TK_NEWLINE
		return nil

	case l.accept('\r'):
		l.accept('\n')
		l.ty = TK_NEWLINE
		return nil

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
