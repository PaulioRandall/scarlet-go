// sanitiser package was created to remove redundant tokens, such as comment and
// whitespace, and to apply formatting o values such as trimming off the quotes
// from string literals and templates.
//
// Key decisions: N/A
package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// SanitiseAll creates a sanitisers all tokens from s returning a new array.
func SanitiseAll(in []token.Token) (out []token.Token) {

	itr := token.NewIterator(in)
	var prev token.Token

	for prev.Type != token.EOF && !itr.Empty() {
		prev = sanitise(itr, prev)
		out = append(out, prev)
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

// isParsableToken returns true if next is a toke of value to the parser.
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
		case token.LIST_OPEN,
			token.DELIM,
			token.TERMINATOR,
			token.BLOCK_OPEN,
			token.UNDEFINED:

			return false
		}
	}

	return true
}

// Applies any special formatting to the token such as converting its token
// type or trimming runes off its value.
func formatToken(tk token.Token) token.Token {

	switch tk.Type {
	case token.NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		tk.Type = token.TERMINATOR

	case token.STRING, token.TEMPLATE:
		// Removes prefix and suffix from tk.Value
		s := tk.Value
		tk.Value = s[1 : len(s)-1]
	}

	return tk
}
