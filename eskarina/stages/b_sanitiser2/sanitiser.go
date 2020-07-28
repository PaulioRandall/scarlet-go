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
		case curr.Tok.IsRedundant():
			remove(curr)

		case curr.Prev == nil && curr.Tok.IsTerminator():
			remove(curr)

		case curr.Prev == nil:

		case curr.Prev.Tok.IsTerminator() && curr.Tok.IsTerminator():
			remove(curr)

		case curr.Prev.Tok == lexeme.LEFT_PAREN && curr.Tok == lexeme.NEWLINE:
			remove(curr)

		case curr.Prev.Tok == lexeme.SEPARATOR && curr.Tok == lexeme.NEWLINE:
			remove(curr)

		case curr.Prev.Tok == lexeme.SEPARATOR && curr.Tok == lexeme.RIGHT_PAREN:
			remove(curr.Prev)
		}
	}

	return first
}
