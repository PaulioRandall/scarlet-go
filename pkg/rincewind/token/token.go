package token

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type Token interface {
	GenType() GenType
	SubType() SubType
	Raw() string
	Value() string
	Snippet
	String() string
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

func Retype(tk Token, gen GenType, sub SubType) Tok {

	r := Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: tk.Raw(),
	}

	r.Line, r.ColBegin = tk.Begin()
	_, r.ColEnd = tk.End()

	return r
}

type Tok struct {
	Gen              GenType
	Sub              SubType
	RawStr           string
	Line             int
	ColBegin, ColEnd int
}

func (tk Tok) GenType() GenType {
	return tk.Gen
}

func (tk Tok) SubType() SubType {
	return tk.Sub
}

func (tk Tok) Raw() string {
	return tk.RawStr
}

func (tk Tok) Value() string {
	return Value(tk)
}

func (tk Tok) Begin() (int, int) {
	return tk.Line, tk.ColBegin
}

func (tk Tok) End() (int, int) {
	return tk.Line, tk.ColEnd
}

func (tk Tok) String() string {
	return String(tk)
}

func Value(tk Token) string {

	switch {
	case tk.GenType() == GE_SPELL:
		return tk.Raw()[1:]

	case tk.SubType() == SU_STRING:
		if len(tk.Raw()) == 2 {
			return ""
		}

		return tk.Raw()[1 : len(tk.Raw())-1]
	}

	return tk.Raw()
}

func String(tk Token) string {

	line, begin := tk.Begin()
	_, end := tk.End()

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		line+1,
		begin,
		line+1,
		end,
		tk.GenType().String(),
		tk.SubType().String(),
		tk.Value(),
	)
}

/*
func (ty TokenType) Precedence() int {
	switch ty {
	case TK_MULTIPLY, TK_DIVIDE, TK_REMAINDER:
		return 6 // Multiplicative

	case TK_PLUS, TK_MINUS:
		return 5 // Additive

	case TK_LESS_THAN, TK_LESS_THAN_OR_EQUAL, TK_MORE_THAN, TK_MORE_THAN_OR_EQUAL:
		return 4 // Relational

	case TK_EQUAL, TK_NOT_EQUAL:
		return 3 // Equalitive

	case TK_AND:
		return 2

	case TK_OR:
		return 1
	}

	return 0
}
*/
