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

func MagicToken(
	gen GenType,
	sub SubType,
	tk Token,
) Token {

	r := tok{
		gen: gen,
		sub: sub,
		raw: tk.Raw(),
	}

	r.line, r.begin = tk.Begin()
	_, r.end = tk.End()

	return r
}

func New_(gen GenType, sub SubType, raw string, line, col int) tok {
	return tok{
		gen:   gen,
		sub:   sub,
		raw:   raw,
		line:  line,
		begin: col,
		end:   col + len(raw),
	}
}

type tok struct {
	gen        GenType
	sub        SubType
	raw        string
	line       int
	begin, end int
}

func (tk tok) GenType() GenType {
	return tk.gen
}

func (tk tok) SubType() SubType {
	return tk.sub
}

func (tk tok) Raw() string {
	return tk.raw
}

func (tk tok) Value() string {

	switch {
	case tk.gen == GE_SPELL:
		return tk.raw[1:]

	case tk.sub == SU_STRING:
		if len(tk.raw) == 2 {
			return ""
		}

		return tk.raw[1 : len(tk.raw)-1]
	}

	return tk.raw
}

func (tk tok) Begin() (int, int) {
	return tk.line, tk.begin
}

func (tk tok) End() (int, int) {
	return tk.line, tk.end
}

func (tk tok) String() string {

	// +1 converts from line index to number
	return fmt.Sprintf(`%d:%d %d:%d %s:%s %q`,
		tk.line+1,
		tk.begin,
		tk.line+1,
		tk.end,
		tk.gen.String(),
		tk.sub.String(),
		tk.Value(),
	)
}
