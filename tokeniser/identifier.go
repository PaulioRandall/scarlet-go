package tokeniser

import (
	"unicode"

	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
	"github.com/PaulioRandall/scarlet-go/tokeniser/source"
)

// idThunk returns a TokeThunk that returns an ID token along with the next
// thunk to use `after`.
func idThunk(s *source.Source, after TokenThunk) TokenThunk {
	return func() (token.Token, TokenThunk, perror.Perror) {
		t := s.SliceBy(countWord)
		return t, after, nil
	}
}

// countWord counts the number of runes in the next word.
func countWord(runes []rune) (n int, k token.Kind) {
	k = token.ID

	for _, ru := range runes {
		if !unicode.IsLetter(ru) {
			break
		}

		n++
	}

	return
}
