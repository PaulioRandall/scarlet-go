package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
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
		case curr.Is(lexeme.PR_REDUNDANT):
			remove(curr)

		case curr.Prev == nil && curr.Is(lexeme.PR_TERMINATOR):
			remove(curr)

		case curr.Prev == nil:

		case curr.Prev.Is(lexeme.PR_TERMINATOR) && curr.Is(lexeme.PR_TERMINATOR):
			remove(curr)

		case curr.Prev.Has(lexeme.PR_PARENTHESIS, lexeme.PR_OPENER) && curr.Is(lexeme.PR_NEWLINE):
			remove(curr)

		case curr.Prev.Is(lexeme.PR_SEPARATOR) && curr.Is(lexeme.PR_NEWLINE):
			remove(curr)

		case curr.Prev.Is(lexeme.PR_SEPARATOR) && curr.Has(lexeme.PR_PARENTHESIS, lexeme.PR_CLOSER):
			remove(curr.Prev)
		}
	}

	return first
}
