package lexeme

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token/token"
)

type Lexeme struct {
	token.Token
	Val  string
	Line int
	Col  int
	End  int
}

func New(val string, tk token.Token, line, col int) Lexeme {
	return Lexeme{
		Token: tk,
		Val:   val,
		Line:  line,
		Col:   col,
		End:   col + len(val),
	}
}

func Tok(val string, tk token.Token) Lexeme {
	return Lexeme{
		Token: tk,
		Val:   val,
	}
}

func (l Lexeme) TokenVal() token.Token {
	return l.Token
}

func (l Lexeme) Value() string {
	return l.Val
}

func (l Lexeme) LineIdx() int {
	return l.Line
}

func (l Lexeme) ColIdx() int {
	return l.Col
}

func (l Lexeme) EndIdx() int {
	return l.End
}

func (l Lexeme) String() string {
	return fmt.Sprintf("[%d] %d:%d %s %q",
		l.Line,
		l.Col,
		l.End,
		l.Token.String(),
		l.Val,
	)
}
