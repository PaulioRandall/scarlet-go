package token

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
)

// Lexeme represents a token from source including its position and value.
type Lexeme struct {
	Token
	scroll.Snippet
}

// Value returns the original text that makes up the Lexeme.
func (l Lexeme) Value() string {
	return l.Snippet.Text
}

// MakeLex returns a new Lexeme.
func MakeLex(tk Token, sn scroll.Snippet) Lexeme {
	return Lexeme{
		Token:   tk,
		Snippet: sn,
	}
}

// MakeLex2 returns a new Lexeme with a zero start position.
func MakeLex2(tk Token, v string) Lexeme {

	sizeBytes := len(v)
	sizeRunes := len([]rune(v))

	sn := scroll.Snippet{
		End: scroll.Position{
			Offset:  sizeBytes,
			ColByte: sizeBytes,
			ColRune: sizeRunes,
		},
	}

	return MakeLex(tk, sn)
}

func (l Lexeme) String() string {
	return fmt.Sprintf("%s %s", l.Token.String(), l.Snippet.String())
}
