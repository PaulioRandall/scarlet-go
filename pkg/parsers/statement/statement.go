package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Statement interface {
	Expression
}

type Expression interface {
	Begin() (line, col int)
	End() (line, col int)
	String() string
}

type Identifier struct {
	TK Token
}

func (id Identifier) Begin() (int, int) {
	tk := id.TK
	return tk.Line(), tk.Col()
}

func (id Identifier) End() (int, int) {
	tk := id.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (id Identifier) String() string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.addToken(0, id.TK)

	return b.String()
}

type Literal struct {
	TK Token
}

func (l Literal) Begin() (int, int) {
	return l.TK.Line(), l.TK.Col()
}

func (l Literal) End() (int, int) {
	tk := l.TK
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (l Literal) String() string {

	b := builder{}

	b.add(0, "[Literal] ")
	b.addToken(0, l.TK)

	return b.String()
}

type List struct {
	Open, Close Token
	Items       []Expression
}

func (l List) Begin() (int, int) {
	return l.Open.Line(), l.Open.Col()
}

func (l List) End() (int, int) {
	tk := l.Close
	return tk.Line(), tk.Col() + len(tk.Value())
}

func (l List) String() string {

	b := builder{}

	b.add(0, "[List] ")

	for _, item := range l.Items {
		b.newline()
		b.add(1, item.String())
	}

	return b.String()
}

type Assignment struct {
	Target Token
	Source Expression
}

func (a Assignment) Begin() (int, int) {
	tk := a.Target
	return tk.Line(), tk.Col()
}

func (a Assignment) End() (int, int) {
	return a.Source.End()
}

func (a Assignment) String() string {

	b := builder{}

	b.add(0, "[Assignment] ")

	b.newline()
	b.add(1, "Target: ")
	b.addToken(1, a.Target)

	b.newline()
	b.add(1, "Source: ")
	b.newline()
	b.add(2, a.Source.String())

	return b.String()
}

type AssignmentBlock struct {
	Stats []Assignment
}

func (ab AssignmentBlock) Begin() (int, int) {
	return ab.Stats[0].Begin()
}

func (ab AssignmentBlock) End() (int, int) {
	i := len(ab.Stats) - 1
	return ab.Stats[i].End()
}

func (ab AssignmentBlock) String() string {

	b := builder{}

	for _, a := range ab.Stats {
		b.add(0, a.String())
	}

	return b.String()
}

type Negation struct {
	Expr Expression
}

func (n Negation) Begin() (int, int) {
	return n.Expr.Begin()
}

func (n Negation) End() (int, int) {
	return n.Expr.End()
}

func (n Negation) String() string {

	b := builder{}

	b.add(0, "[Negation] ")

	b.newline()
	b.add(1, n.Expr.String())

	return b.String()
}
