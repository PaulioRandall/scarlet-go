package scanner2

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

func scan(lex *lexeme) error {

	switch {
	case lex.accept('\r'), lex.match('\n'):
		lex.ty = TK_NEWLINE
		return lex.expect('\n')

	case unicode.IsSpace(lex.get()):
		whitespace(lex)
		return nil

	case lex.accept('_'):
		lex.ty = TK_VOID
		return nil
	}

	return err.New("Unknown symbol", err.Pos(lex.scn.line, lex.scn.col))
}

func whitespace(lex *lexeme) {

	lex.ty = TK_WHITESPACE

	for ru := lex.get(); unicode.IsSpace(ru); ru = lex.get() {

		if ru == '\r' || ru == '\n' {
			return
		}

		lex.next()
	}
}
