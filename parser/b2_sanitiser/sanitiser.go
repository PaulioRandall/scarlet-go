package sanitiser2

import (
	"github.com/PaulioRandall/scarlet-go/token/container"
	//"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

func SanitiseAll(con *container.Container) {

	if con.Empty() {
		return
	}

	for itr := con.Iterator(); itr.HasNext(); {
		switch itr.Next(); {
		case itr.Item().Type().IsRedundant():
			itr.Remove()
			/*
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
			*/
		}
	}
}
