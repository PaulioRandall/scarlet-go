package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/token"
)

// ScanToken is a recursive descent function that returns the next token
// followed by the callable (tail) function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken)

// no_tok returns an empty Token.
func no_tok() token.Token {
	return token.Token{}
}

// fileScope identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (s *source) fileScope() (t token.Token, f ScanToken) {

	var k token.Kind
	var n int

	switch ru := s.runes[0]; {
	case unicode.IsSpace(ru):
		k, n = token.WHITESPACE, countSpaces(s.runes)
	default:
		return no_tok(), s.fileScope
	}

	return s.scan(n, k), s.fileScope
}
