package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func SanitiseAll(con *lexeme.Container) *lexeme.Container {

	if con.Empty() {
		return con
	}

	itr := con.Iterator()

	for itr.Next() || itr.After() != nil {

		switch {
		case itr.Curr().Tok.IsRedundant():
			itr.Remove()

		case itr.Before() == nil && itr.Curr().Tok.IsTerminator():
			itr.Remove()

		case itr.Before() == nil:

		case itr.Before().Tok.IsTerminator() && itr.Curr().Tok.IsTerminator():
			itr.Remove()

		case itr.Before().Tok == lexeme.LEFT_PAREN && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Before().Tok == lexeme.SEPARATOR && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Before().Tok == lexeme.SEPARATOR && itr.Curr().Tok == lexeme.RIGHT_PAREN:
			itr.Prev()
			itr.Remove()
			itr.Next()
		}
	}

	return con
}
