package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// SanitiseAll removes redundant tokens, such as comment and whitespace, as well
// as applying formatting to values, e.g trimming off the quotes from string
// literals and templates.
func SanitiseAll(in []token.Token) (out []token.Token) {

	itr := token.NewIterator(in)
	var prev token.Token

	for prev.Type != token.EOF && !itr.Empty() {
		p := sanitise(itr, prev)

		if p != (token.Token{}) {
			out = append(out, p)
			prev = p
		}
	}

	return out
}

func sanitise(itr *token.TokenIterator, prev token.Token) token.Token {

	for !isParsableToken(prev, itr.Next()) {
		if itr.Empty() {
			return token.Token{}
		}
	}

	return formatToken(itr.Past())
}

func isParsableToken(prev, next token.Token) bool {

	past := prev.Type
	curr := next.Type

	switch curr {

	case token.UNDEFINED, token.WHITESPACE, token.COMMENT:
		return false

	case token.NEWLINE, token.TERMINATOR:

		switch past {
		// Sometimes the extra terminator or newline is redundant.
		// Removing them makes parsing easier.
		case token.TERMINATOR,
			token.DELIM,
			token.BLOCK_OPEN,
			token.BLOCK_CLOSE,
			token.MATCH,
			token.LIST,
			token.UNDEFINED:

			return false
		}
	}

	return true
}

func formatToken(tk token.Token) token.Token {

	switch tk.Type {
	case token.NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		tk.Type = token.TERMINATOR

	case token.STRING, token.TEMPLATE:
		// Avoid issues later by removing the quote marks.
		s := tk.Value
		tk.Value = s[1 : len(s)-1]
	}

	return tk
}
