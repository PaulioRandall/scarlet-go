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

	for curr, next := first, first; curr != nil; curr = next {
		next = next.Next

		switch {
		case curr.Has(prop.PR_REDUNDANT):
			remove(curr)

		case curr.Prev == nil && curr.Has(prop.PR_TERMINATOR):
			remove(curr)

		case curr.Prev == nil:

		case curr.Prev.Has(prop.PR_TERMINATOR) && curr.Has(prop.PR_TERMINATOR):
			remove(curr)

		case curr.Prev.Is(prop.PR_PARENTHESIS, prop.PR_OPENER) && curr.Has(prop.PR_NEWLINE):
			remove(curr)

		case curr.Prev.Has(prop.PR_SEPARATOR) && curr.Has(prop.PR_NEWLINE):
			remove(curr)

		case curr.Prev.Has(prop.PR_SEPARATOR) && curr.Is(prop.PR_PARENTHESIS, prop.PR_CLOSER):
			remove(curr.Prev)
		}
	}

	return first
}
