package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type Iterator interface {
	ToContainer() *lexeme.Container
	Prev() bool
	Next() bool
	Curr() *lexeme.Lexeme
	Remove() *lexeme.Lexeme
	Behind() *lexeme.Lexeme
	Ahead() *lexeme.Lexeme
}

func SanitiseAll(con *lexeme.Container) *lexeme.Container {

	if con.Empty() {
		return con
	}

	itr := Iterator(con.ToItinerant())

	for itr.Next() || itr.Ahead() != nil {

		switch {
		case itr.Curr().Tok.IsRedundant():
			itr.Remove()

		case itr.Behind() == nil && itr.Curr().Tok.IsTerminator():
			itr.Remove()

		case itr.Behind() == nil:

		case itr.Behind().Tok.IsTerminator() && itr.Curr().Tok.IsTerminator():
			itr.Remove()

		case itr.Behind().Tok == lexeme.LEFT_PAREN && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Behind().Tok == lexeme.SEPARATOR && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Behind().Tok == lexeme.SEPARATOR && itr.Curr().Tok == lexeme.RIGHT_PAREN:
			itr.Prev()
			itr.Remove()
			itr.Next()
		}
	}

	return itr.ToContainer()
}
