// sanitiser package was created to remove redundant tokens, such as comment and
// whitespace, and to apply formatting o values such as trimming off the quotes
// from string literals and templates.
//
// Key decisions: N/A
package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

// SanitiseAll creates a sanitisers all tokens from s returning a new array.
func SanitiseAll(in []lexeme.Token) (out []lexeme.Token) {

	sn := sanitiser{
		itr: lexeme.NewIterator(in),
	}

	for !sn.itr.Empty() {
		tk := sn.read()
		out = append(out, tk)
	}

	return out
}

type sanitiser struct {
	itr  *lexeme.TokenIterator
	prev lexeme.Token
}

func (sn *sanitiser) read() (_ lexeme.Token) {

	for tk := sn.itr.Peek(); !sn.itr.Empty(); tk = sn.itr.Peek() {
		sn.itr.Skip()

		if !isRedundantLexeme(tk.Lexeme, sn.prev.Lexeme) {
			tk = formatToken(tk)
			sn.prev = tk
			return tk
		}
	}

	return
}

// isRedundantLexeme returns true if l is considered redundant to parsing.
func isRedundantLexeme(l, prev lexeme.Lexeme) bool {

	if l == lexeme.LEXEME_WHITESPACE || l == lexeme.LEXEME_COMMENT {
		return true
	}

	if l != lexeme.LEXEME_NEWLINE && l != lexeme.LEXEME_TERMINATOR {
		return false
	}

	return prev == lexeme.LEXEME_LIST_OPEN ||
		prev == lexeme.LEXEME_DELIM ||
		prev == lexeme.LEXEME_TERMINATOR ||
		prev == lexeme.LEXEME_DO ||
		prev == lexeme.LEXEME_UNDEFINED
}

// Applies any special formatting to the token such as converting its lexeme
// type or trimming runes off its value.
func formatToken(tk lexeme.Token) lexeme.Token {

	switch tk.Lexeme {
	case lexeme.LEXEME_NEWLINE:
		// Non-redundant newline tokens are expression and statement terminators
		// in disguise.
		tk.Lexeme = lexeme.LEXEME_TERMINATOR

	case lexeme.LEXEME_STRING, lexeme.LEXEME_TEMPLATE:
		// Removes prefix and suffix from tk.Value
		s := tk.Value
		tk.Value = s[1 : len(s)-1]
	}

	return tk
}
