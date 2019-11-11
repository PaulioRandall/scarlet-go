package scanner

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ScanToken is a recursive descent function that returns the next token
// followed by the callable (tail) function to get the token after next. If the
// function is null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken, perror.Perror)

// EmptyTok returns an empty Token.
func EmptyTok() token.Token {
	return token.Token{}
}

// fileScope identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (s *source) fileScope() (token.Token, ScanToken, perror.Perror) {

	var k token.Kind
	var n int

	if len(s.runes) == 0 {
		return EmptyTok(), nil, nil
	}

	switch ru := s.runes[0]; {
	case newlineRunes(s.runes, 0) != 0:
		return s.scanNewline(), s.fileScope, nil
	case unicode.IsSpace(ru):
		k, n = token.WHITESPACE, countSpaces(s.runes)
	default:
		return EmptyTok(), nil, perror.New(
			"Unknown token",
			s.line,
			s.col,
			s.col,
		)
	}

	return s.scan(n, k), s.fileScope, nil
}
