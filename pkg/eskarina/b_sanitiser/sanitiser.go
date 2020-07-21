package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func SanitiseAll(first *lexeme.Lexeme) *lexeme.Lexeme {

	if first == nil {
		return nil
	}

	remove := func(lex *lexeme.Lexeme) {
		if first == lex {
			first = lex.Next
		}
		lex.Remove()
	}

	for lex := first; lex != nil; lex = lex.Next {

		switch {
		case lex.Has(prop.PR_REDUNDANT):
			remove(lex)

		case lex.Prev == nil && lex.Has(prop.PR_TERMINATOR):
			remove(lex)

		case lex.Prev == nil:

		case lex.Prev.Has(prop.PR_TERMINATOR) && lex.Has(prop.PR_TERMINATOR):
			remove(lex)

		case lex.Prev.Is(prop.PR_PARENTHESIS, prop.PR_OPENER) && lex.Has(prop.PR_NEWLINE):
			remove(lex)

		case lex.Prev.Has(prop.PR_SEPARATOR) && lex.Has(prop.PR_NEWLINE):
			remove(lex)

		case lex.Prev.Has(prop.PR_SEPARATOR) && lex.Is(prop.PR_PARENTHESIS, prop.PR_CLOSER):
			remove(lex.Prev)
		}
	}

	return first
}
