package token

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
)

type Token interface {
	GenType() GenType
	SubType() SubType
	Entity
	Raw() string
	Value() string
	Snippet
	String() string
}

type Snippet interface {
	Begin() (int, int)
	End() (int, int)
}

type Tok struct {
	Gen              GenType
	Sub              SubType
	RawProps         []Prop
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

func (tk Tok) Props() []Prop {
	return tk.RawProps
}

func (tk Tok) Is(o Prop) bool {

	for _, p := range tk.RawProps {
		if p == o {
			return true
		}
	}

	return false
}

func (tk Tok) IsNot(o Prop) bool {
	return !tk.Is(o)
}

func (tk Tok) Raw() string {
	return tk.RawStr
}

func (tk Tok) Value() string {

	if tk.Is(PR_SPELL) {
		return tk.RawStr[1:]
	}

	if tk.Is(PR_STRING) {
		if len(tk.RawStr) == 2 {
			return ""
		}

		return tk.RawStr[1 : len(tk.RawStr)-1]
	}

	return tk.RawStr
}

func (tk Tok) Begin() (int, int) {
	return tk.Line, tk.ColBegin
}

func (tk Tok) End() (int, int) {
	return tk.Line, tk.ColEnd
}

func (tk Tok) String() string {

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		tk.Line+1,
		tk.ColBegin,
		tk.Line+1,
		tk.ColEnd,
		tk.Gen.String(),
		tk.Sub.String(),
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
