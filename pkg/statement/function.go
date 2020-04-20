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

	str += indent(i+1) + "Output: " + newline()
	for _, tk := range f.Output {
		str += indent(i+2) + tk.String() + newline()
	}

	return str
}

type FuncCall struct {
	ID    token.Token
	Input []token.Token
}

func (f FuncCall) Token() token.Token {
	return f.ID
}

func (f FuncCall) String(i int) string {

	str := indent(i) + "[FuncCall] " + f.ID.String() + newline()

	str += indent(i+1) + "Input:"
	for _, tk := range f.Input {
		str += newline() + tk.String()
	}

	return str
}
