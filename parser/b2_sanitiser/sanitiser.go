package sanitiser2

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type Iterator interface {
	HasNext() bool
	HasPrev() bool
	Next() lexeme.Lexeme
	Item() lexeme.Lexeme
	Remove() lexeme.Lexeme
}

func SanitiseAll(itr Iterator) {
	for itr.HasNext() {
		switch itr.Next(); {
		case itr.Item().Type().IsRedundant():
			itr.Remove() // Remove tokens redundant to the compiler

		case !itr.HasPrev() && itr.Item().Type().IsTerminator():
			itr.Remove() // Remove terminators at the start of the scroll

		case !itr.HasPrev():
			continue

			/*
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
			*/
		}
	}
}
