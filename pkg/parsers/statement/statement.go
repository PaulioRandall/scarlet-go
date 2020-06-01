package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func startPos(tk Token) (line int, col int) {
	return tk.Line(), tk.Col()
}

func endPos(tk Token) (line int, col int) {
	return tk.Line(), tk.Col() + len(tk.Value())
}

type Statement interface {
	Expression
}

type Expression interface {
	Begin() (line, col int)
	End() (line, col int)
	String() string
}

type Void struct {
	TK Token
}

func (v Void) Begin() (int, int) {
	return startPos(v.TK)
}

func (v Void) End() (int, int) {
	return endPos(v.TK)
}

func (v Void) String() string {

	b := builder{}

	b.add(0, "[Void] ")
	b.addToken(0, v.TK)

	return b.String()
}

type Identifier struct {
	TK Token
}

func (id Identifier) Begin() (int, int) {
	return startPos(id.TK)
}

func (id Identifier) End() (int, int) {
	return endPos(id.TK)
}

func (id Identifier) String() string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.addToken(0, id.TK)

	return b.String()
}

type ListAccessor struct {
	List  Expression
	Index Expression
}

func (la ListAccessor) Begin() (int, int) {
	return la.List.Begin()
}

func (la ListAccessor) End() (int, int) {
	return la.Index.End()
}

func (la ListAccessor) String() string {

	b := builder{}

	b.add(0, "[ListAccessor] ")

	b.newline()
	b.add(1, "List: ")
	b.newline()
	b.add(2, la.List.String())

	b.newline()
	b.add(1, "Index: ")
	b.newline()
	b.add(2, la.Index.String())

	return b.String()
}

type Literal struct {
	TK Token
}

func (l Literal) Begin() (int, int) {
	return startPos(l.TK)
}

func (l Literal) End() (int, int) {
	return endPos(l.TK)
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
	return startPos(l.Open)
}

func (l List) End() (int, int) {
	return endPos(l.Close)
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
	Target Expression
	Source Expression
}

func (a Assignment) Begin() (int, int) {
	return a.Target.Begin()
}

func (a Assignment) End() (int, int) {
	return a.Source.End()
}

func (a Assignment) String() string {

	b := builder{}

	b.add(0, "[Assignment] ")

	b.newline()
	b.add(1, "Target: ")
	b.newline()
	b.add(1, a.Target.String())

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
