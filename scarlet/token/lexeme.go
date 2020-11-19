// Defines the Lexeme struct and utility functions. Lexemes tie together text
// Snippet from source code with a Token. Lexemes are passed along the parsing
// pipeline instead of Tokens so the text value can be used for non-keywords
// and grammer while the position is used for logging and error messages.

package token

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
)

// Lexeme represents a token from source code including its position and raw
// value.
type Lexeme struct {
	Token
	Snippet // @Retire
	position.Range
	Val string
}

// Value returns the original text that makes up the Lexeme.
func (l Lexeme) Value() string {
	return l.Val
}

// Make returns a new Lexeme. It's convenience to help avoid construction
// errors.
func Make(val string, tk Token, snip Snippet) Lexeme {
	return Lexeme{
		Token:   tk,
		Snippet: snip,
		Val:     val,
	}
}

// Make returns a new Lexeme with a zero start position and calculated end
//  It's convenience to help avoid construction errors.
func MakeTok(val string, tk Token) Lexeme {

	sizeBytes := len(val)
	sizeRunes := len([]rune(val))

	snip := Snippet{
		End: UTF8Pos{
			Offset:  sizeBytes,
			ColByte: sizeBytes,
			ColRune: sizeRunes,
		},
	}

	return Make(val, tk, snip)
}

func (l Lexeme) String() string {
	return fmt.Sprintf("%s %s %q",
		l.Snippet.String(),
		l.Token.String(),
		l.Val,
	)
}
