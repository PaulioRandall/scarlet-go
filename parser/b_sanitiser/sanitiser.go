package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

func SanitiseAll(con *lexeme.Container) {

	if con.Empty() {
		return
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

		case itr.Before().Tok == lexeme.L_PAREN && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Before().Tok == lexeme.L_CURLY && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Before().Tok == lexeme.DELIM && itr.Curr().Tok == lexeme.NEWLINE:
			itr.Remove()

		case itr.Before().Tok == lexeme.DELIM && itr.Curr().Tok == lexeme.R_PAREN:
			itr.Prev()
			itr.Remove()
			itr.Next()

		case itr.Before().Tok.IsTerminator() && itr.Curr().Tok == lexeme.R_CURLY:
			itr.Prev()
			itr.Remove()
			itr.Next()
		}
	}
}
