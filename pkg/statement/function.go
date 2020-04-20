package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type FuncDef struct {
	Open   token.Token
	Input  []token.Token
	Output []token.Token
	Body   Block
}

func (f FuncDef) Token() token.Token {
	return f.Open
}

func (f FuncDef) String(i int) string {

	str := indent(i) + "[FuncDef] " + f.Open.String() + newline()

	str += indent(i+1) + "Input:" + newline()
	for _, tk := range f.Input {
		str += indent(i+2) + tk.String() + newline()
	}

	str += indent(i+1) + "Output: "
	for _, tk := range f.Output {
		str += newline() + indent(i+2) + tk.String()
	}

	return str
}

type FuncCall struct {
	ID    Expression
	Input []Expression
}

func (f FuncCall) Token() token.Token {
	return f.ID.Token()
}

func (f FuncCall) String(i int) string {

	str := indent(i) + "[FuncCall] " + newline()
	str += indent(i+1) + "ID:" + newline()
	str += f.ID.String(i+2) + newline()

	str += indent(i+1) + "Input:"
	for _, ex := range f.Input {
		str += newline() + ex.String(i+2)
	}

	return str
}
