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

	sn := sanitiser{
		itr: token.NewIterator(in),
	}

	for !sn.itr.Empty() {
		tk := sn.read()
		out = append(out, tk)
	}

	return out
}

type sanitiser struct {
	itr  *token.TokenIterator
	prev token.Token
}

func (sn *sanitiser) read() (_ token.Token) {

	for tk := sn.itr.Peek(); !sn.itr.Empty(); tk = sn.itr.Peek() {
		sn.itr.Skip()

		if !isRedundantType(tk.Type, sn.prev.Type) {
			tk = formatToken(tk)
			sn.prev = tk
			return tk
		}
	}

	return
}

// isRedundantType returns true if l is considered redundant to parsing.
func isRedundantType(l, prev token.TokenType) bool {

	if l == token.WHITESPACE || l == token.COMMENT {
		return true
	}

	if l != token.NEWLINE && l != token.TERMINATOR {
		return false
	}

	return prev == token.LIST_OPEN ||
		prev == token.DELIM ||
		prev == token.TERMINATOR ||
		prev == token.BLOCK_OPEN ||
		prev == token.UNDEFINED
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
