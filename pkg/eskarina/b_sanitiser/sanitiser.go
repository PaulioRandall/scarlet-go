package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func SanitiseAll(first *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := first; lex != nil; lex = lex.Next {
		if !sanitise(lex) {
			first = lex
			break
		}
	}

	for lex := first; lex != nil; lex = lex.Next {
		sanitise(lex)
	}

	return first
}

func sanitise(lex *lexeme.Lexeme) bool {

	switch {
	case lex.Has(prop.PR_REDUNDANT):
		lex.Remove()

	case lex.Prev == nil && lex.Has(prop.PR_TERMINATOR):
		lex.Remove()

	case lex.Prev == nil:
		return false

	case lex.Prev.Has(prop.PR_TERMINATOR) && lex.Has(prop.PR_TERMINATOR):
		lex.Remove()

	case lex.Prev.Is(prop.PR_PARENTHESIS, prop.PR_OPENER) && lex.Has(prop.PR_NEWLINE):
		lex.Remove()

	case lex.Prev.Has(prop.PR_SEPARATOR) && lex.Has(prop.PR_NEWLINE):
		lex.Remove()

	case lex.Prev.Has(prop.PR_SEPARATOR) && lex.Is(prop.PR_PARENTHESIS, prop.PR_CLOSER):
		lex.Prev.Remove()

	default:
		return false
	}

	return true
}
