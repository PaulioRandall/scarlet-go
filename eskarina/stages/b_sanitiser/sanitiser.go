package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func SanitiseAll(con *lexeme.Container) *lexeme.Container {

	if con.Empty() {
		return con
	}

	w := con.To().Window()

	for w.Next() || w.Ahead() != nil {

		switch {
		case w.Curr().Tok.IsRedundant():
			w.Remove()

		case w.Behind() == nil && w.Curr().Tok.IsTerminator():
			w.Remove()

		case w.Behind() == nil:

		case w.Behind().Tok.IsTerminator() && w.Curr().Tok.IsTerminator():
			w.Remove()

		case w.Behind().Tok == lexeme.LEFT_PAREN && w.Curr().Tok == lexeme.NEWLINE:
			w.Remove()

		case w.Behind().Tok == lexeme.SEPARATOR && w.Curr().Tok == lexeme.NEWLINE:
			w.Remove()

		case w.Behind().Tok == lexeme.SEPARATOR && w.Curr().Tok == lexeme.RIGHT_PAREN:
			w.Prev()
			w.Remove()
			w.Next()
		}
	}

	return w.To().Container()
}
