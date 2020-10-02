package lexeme

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// Lexeme represents a token from source code including its position and raw
// value.
type Lexeme struct {
	token.Token
	position.Snippet
	Val string
}

// Make returns a new Lexeme. It's convenience to help avoid construction
// errors.
func Make(val string, tk token.Token, snip position.Snippet) Lexeme {
	return Lexeme{
		Token:   tk,
		Snippet: snip,
		Val:     val,
	}
}

// Make returns a new Lexeme with a zero start position and calculated end
// position. It's convenience to help avoid construction errors.
func MakeTok(val string, tk token.Token) Lexeme {

	sizeBytes := len(val)
	sizeRunes := len([]rune(val))

	snip := position.Snippet{
		End: position.Position{
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
