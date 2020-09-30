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

func New(val string, tk token.Token, line, col int) Lexeme {

	snip := token.Snippet{
		Position: token.Position{
			SrcOffset: 0, // TODO
			LineIdx:   line,
			ColByte:   0, // TODO
			ColRune:   col,
		},
		End: token.Position{
			SrcOffset: 0, // TODO
			LineIdx:   line,
			ColByte:   0, // TODO
			ColRune:   col + len(val),
		},
	}

	return Lexeme{
		Token:   tk,
		Snippet: snip,
		Val:     val,
	}
}

func Tok(val string, tk token.Token) Lexeme {
	return New(val, tk, 0, 0)
}

func (l Lexeme) Value() string {
	return l.Val
}

func (l Lexeme) String() string {
	return fmt.Sprintf("%s %s %q",
		l.Snippet.String(),
		l.Token.String(),
		l.Val,
	)
}
