package recursive

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

func Print(ss Statements) {
	s := ss.String()
	s = strings.ReplaceAll(s, "\t", "  ")
	println(s)
	println(lexeme.LEXEME_EOF)
	println()
}

type Statements []Statement

// PrintAll pretty prints t.
func (ss *Statements) String() string {

	slice := []Statement(*ss)
	str := "[Statements]"

	for _, s := range slice {
		str += newline()
		str += s.String(1)
	}

	return str
}

type Statement struct {
	IDs    []lexeme.Token
	Assign lexeme.Token
	Exprs  []Expression
}

func (s *Statement) String(i int) string {

	str := indent(i) + "[Statement]"
	if s.Assign != (lexeme.Token{}) {
		str += " " + s.Assign.String()
	}
	str += newline()

	str += indent(i+1) + "IDs:" + newline()
	for _, tk := range s.IDs {
		str += indent(i+2) + tk.String() + newline()
	}

	str += indent(i+1) + "Exprs:"
	for _, expr := range s.Exprs {
		str += newline()
		str += expr.String(i + 2)
	}

	return str
}

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

func indent(indent int) string {
	return strings.Repeat("\t", indent)
}

func newline() string {
	return "\n"
}
