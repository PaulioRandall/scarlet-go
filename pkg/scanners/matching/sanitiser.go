package matching

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func sanitiseNext(itr *token.TokenIterator, prev token.Token) token.Token {

	var tk token.Token

	for tk == (token.Token{}) {
		if itr.Empty() {
			return token.Token{}
		}

		tk = sanitise(prev, itr.Next())
	}

	return tk
}

func sanitise(prev, next token.Token) token.Token {

	if isParsableToken(prev, next) {
		return formatToken(next)
	}

	return token.Token{}
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
