package source

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cookies"
	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/scanner/count"
	"github.com/PaulioRandall/scarlet-go/token"
)

// scan identifies and returns the next token in the source. The token must
// be one that appears at the start of a statement within the top level of a
// source file.
func (s *source) scan() (token.Token, token.ScanToken, perror.Perror) {

	var k token.Kind
	var n int

	if len(s.runes) == 0 {
		return token.Empty(), nil, nil
	}

	switch ru := s.runes[0]; {
	case cookies.NewlineRunes(s.runes, 0) != 0:
		return s.scanNewline(), s.scan, nil

	case unicode.IsSpace(ru):
		k = token.WHITESPACE
		n = count.CountSpaces(s.runes)
		return s.scanRunes(n, k), s.scan, nil

	case unicode.IsLetter(ru):
		n = count.CountWord(s.runes)
		return s.scanWord(n), s.scan, nil
	}

	return token.Empty(), nil, perror.New(
		"Unknown token",
		s.line,
		s.col,
		s.col,
	)
}
