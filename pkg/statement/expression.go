package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
)

type Expression interface {
	Token() lexeme.Token

	String(indent int) string
}

type Value struct {
	Source lexeme.Token
}

func (v Value) Token() lexeme.Token {
	return v.Source
}

func (v Value) String(i int) string {
	return indent(i) + "[Value] " + v.Source.String()
}

type Arithmetic struct {
	Left     Expression
	Operator lexeme.Token
	Right    Expression
}

func (a Arithmetic) Token() lexeme.Token {
	return a.Operator
}

func (a Arithmetic) String(i int) string {

	str := indent(i) + "[Arithmetic] " + a.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += a.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += a.Right.String(i+2) + newline()

	return str
}

type Logic struct {
	Left     Expression
	Operator lexeme.Token
	Right    Expression
}

func (l Logic) Token() lexeme.Token {
	return l.Operator
}

func (l Logic) String(i int) string {

	str := indent(i) + "[Logic] " + l.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += l.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += l.Right.String(i+2) + newline()

	return str
}

type FuncCall struct {
	ID     lexeme.Token
	Input  []lexeme.Token
	Output []lexeme.Token
}

func (f FuncCall) Token() lexeme.Token {
	return f.ID
}

func (f FuncCall) String(i int) string {

	str := indent(i) + "[FuncCall] " + f.ID.String() + newline()

	str += indent(i+1) + "Input:" + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	str += indent(i+1) + "Output: " + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	return str
}
