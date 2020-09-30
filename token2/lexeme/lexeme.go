package lexeme

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/token"
)

type Lexeme struct {
	token.Token
	token.Snippet
	Val string
}

func New(val string, tk token.Token, snip token.Snippet) Lexeme {
	return Lexeme{
		Token:   tk,
		Snippet: snip,
		Val:     val,
	}
}

func Tok(val string, tk token.Token) Lexeme {

	sizeBytes := len(val)
	sizeRunes := len([]rune(val))

	snip := token.Snippet{
		End: token.Position{
			SrcOffset: sizeBytes,
			ColByte:   sizeBytes,
			ColRune:   sizeRunes,
		},
	}

	return New(val, tk, snip)
}

func (l Lexeme) String() string {
	return fmt.Sprintf("%s %s %q",
		l.Snippet.String(),
		l.Token.String(),
		l.Val,
	)
}
