package sanitiser2

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type Iterator interface {
	HasNext() bool
	HasPrev() bool
	Next() lexeme.Lexeme
	Item() lexeme.Lexeme
	LookBehind() lexeme.Lexeme
	Remove() lexeme.Lexeme
}

func SanitiseAll(itr Iterator) {
	for itr.HasNext() {
		curr := itr.Next()

		if curr.IsRedundant() {
			itr.Remove() // Remove tokens redundant to the compiler
			continue
		}

		if !itr.HasPrev() {
			if curr.IsTerminator() {
				itr.Remove() // Remove terminators at the start of the scroll
			}
			continue // No action for the first valid token
		}

		switch prev := itr.LookBehind(); {
		case prev.IsTerminator() && curr.IsTerminator():
			itr.Remove() // Remove successive terminators

		case prev.IsOpener() && curr.Type() == lexeme.NEWLINE:
			itr.Remove()

		case prev.Type() == lexeme.DELIM && curr.Type() == lexeme.NEWLINE:
			itr.Remove()
		}

		/*
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
