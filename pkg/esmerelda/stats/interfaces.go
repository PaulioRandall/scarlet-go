package stats

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type Expr interface {
	Snippet
	Kind() Kind
}

type Snippet interface {
	fmt.Stringer
	Begin() (line, col int)
	End() (line, col int)
}

type Exit interface {
	Expr
	Tk() Token
	Code() Expr
}

func ExitString(e Exit) string {

	b := builder{}

	b.add(0, "[Exit] ")
	b.add(0, e.Tk().String())

	b.newline()
	b.add(1, "Code: ")
	b.newline()
	b.add(2, e.Code().String())

	return b.String()
}

type Void interface {
	Expr
	Tk() Token
}

func VoidString(v Void) string {

	b := builder{}

	b.add(0, "[Void] ")
	b.add(0, v.Tk().String())

	return b.String()
}

type Identifier interface {
	Expr
	Tk() Token
}

func IdentifierString(id Identifier) string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.add(0, id.Tk().String())

	return b.String()
}

type Literal interface {
	Expr
	Tk() Token
}

func LiteralString(l Literal) string {

	b := builder{}

	b.add(0, "[Literal] ")
	b.add(0, l.Tk().String())

	return b.String()
}

type ContainerItem interface {
	Expr
	Container() Expr
	Key() Expr
}

func ContainerItemString(c ContainerItem) string {

	b := builder{}

	b.add(0, "[ContainerItem] ")

	b.newline()
	b.add(1, "Container: ")
	b.newline()
	b.add(2, c.Container().String())

	b.newline()
	b.add(1, "Key: ")
	b.newline()
	b.add(2, c.Key().String())

	return b.String()
}

type Negation interface {
	Expr
	Expr() Expr
}

func NegationString(n Negation) string {

	b := builder{}

	b.add(0, "[Negation]")

	b.newline()
	b.add(1, n.Expr().String())

	return b.String()
}

type Operation interface {
	Expr
	Operator() Token
	Left() Expr
	Right() Expr
}

func OperationString(o Operation) string {

	b := builder{}

	b.add(0, "[Operation] ")
	b.add(0, o.Operator().String())

	b.newline()
	b.add(1, "Left: ")
	b.newline()
	b.add(2, o.Left().String())

	b.newline()
	b.add(1, "Right: ")
	b.newline()
	b.add(2, o.Right().String())

	return b.String()
}

type Assign interface {
	Expr
	Const() bool
	Target() Expr
	Source() Expr
}

func AssignString(a Assign) string {

	b := builder{}

	b.add(0, "[Assign] ")
	if a.Const() {
		b.add(0, "Const")
	}

	b.newline()
	b.add(1, "Target: ")
	b.newline()
	b.add(1, a.Target().String())

	b.newline()
	b.add(1, "Source: ")
	b.newline()
	b.add(2, a.Source().String())

	return b.String()
}

type AssignBlock interface {
	Expr
	Const() bool
	Targets() []Expr
	Sources() []Expr
	Count() int
}

func AssignBlockString(bk AssignBlock) string {

	b := builder{}
	b.add(0, "[AssignBlock] ")
	if bk.Const() {
		b.add(0, "Const")
	}

	for i := 0; i < bk.Count(); i++ {

		b.newline()
		b.add(1, "Target: ")
		b.newline()
		b.add(1, bk.Targets()[i].String())

		b.newline()
		b.add(1, "Source: ")
		b.newline()
		b.add(2, bk.Sources()[i].String())
	}

	return b.String()
}

type ExprFunc interface {
	Expr
	Inputs() []Token
	Expr() Expr
}

func ExprFuncString(e ExprFunc) string {

	b := builder{}

	b.add(0, "[ExprFunc] ")

	if len(e.Inputs()) > 0 {

		b.newline()
		b.add(1, "Inputs: ")

		for _, in := range e.Inputs() {
			b.newline()
			b.add(2, in.String())
		}
	}

	b.newline()
	b.add(1, e.Expr().String())

	return b.String()
}

type FuncDef interface {
	Expr
	Inputs() []Token
	Outputs() []Token
	Body() Expr
}

func FuncDefString(f FuncDef) string {

	b := builder{}

	b.add(0, "[FuncDef] ")

	if len(f.Inputs()) > 0 {

		b.newline()
		b.add(1, "Inputs: ")

		for _, in := range f.Inputs() {
			b.newline()
			b.add(2, in.String())
		}
	}

	if len(f.Outputs()) > 0 {

		b.newline()
		b.add(1, "Outputs: ")

		for _, out := range f.Outputs() {
			b.newline()
			b.add(2, out.String())
		}
	}

	b.newline()
	b.add(1, f.Body().String())

	return b.String()
}

type FuncCall interface {
	Expr
	Function() Expr
	Arguments() []Expr
}

func FuncCallString(f FuncCall) string {

	b := builder{}

	b.add(0, "[FuncCall]")

	b.newline()
	b.add(1, "Function:")
	b.newline()
	b.add(2, f.Function().String())

	b.newline()
	b.add(1, "Arguments:")
	for _, a := range f.Arguments() {
		b.newline()
		b.add(2, a.String())
	}

	return b.String()
}

type Block interface {
	Expr
	Stats() []Expr
}

func BlockString(bk Block) string {

	b := builder{}

	for _, a := range bk.Stats() {
		b.add(0, a.String())
	}

	return b.String()
}

type Watch interface {
	Expr
	Identifiers() []Token
	Body() Block
}

func WatchString(w Watch) string {

	b := builder{}

	b.add(0, "[Watch] ")

	b.newline()
	b.add(1, "Identifiers:")
	for _, id := range w.Identifiers() {
		b.newline()
		b.add(2, id.String())
	}

	b.newline()
	b.add(1, BlockString(w.Body()))

	return b.String()
}

type Guard interface {
	Expr
	Condition() Expr
	Body() Block
}

func GuardString(g Guard) string {

	b := builder{}

	b.add(0, "[Guard] ")

	b.newline()
	b.add(1, g.Condition().String())

	b.newline()
	b.add(1, BlockString(g.Body()))

	return b.String()
}

type WhenCase interface {
	Expr
	Object() Expr
	Body() Block
}

func WhenCaseString(mc WhenCase) string {

	b := builder{}

	b.add(0, "[WhenCase] ")

	b.newline()
	b.add(1, "Object:")
	b.newline()
	b.add(2, mc.Object().String())

	b.newline()
	b.add(1, "Body:")
	b.newline()
	b.add(2, BlockString(mc.Body()))

	return b.String()
}

type When interface {
	Expr
	Subject() Token
	Init() Expr
	Cases() []Expr
}

func WhenString(m When) string {

	b := builder{}

	b.add(0, "[When]")
	b.newline()

	b.add(1, "Target ")
	b.add(1, m.Subject().String())

	b.newline()
	b.add(1, "Init:")
	b.newline()
	b.add(2, m.Init().String())

	for _, mc := range m.Cases() {
		b.newline()
		b.add(2, mc.String())
	}

	return b.String()
}

type Loop interface {
	Expr
	Initialiser() Assign
	Guard() Guard
}

func LoopString(l Loop) string {

	b := builder{}

	b.add(0, "[Loop]")

	b.newline()
	b.add(1, "Initialiser:")
	b.newline()
	b.add(2, AssignString(l.Initialiser()))

	b.newline()
	b.add(1, "Guard:")
	b.newline()
	b.add(2, GuardString(l.Guard()))

	return b.String()
}

type SpellCall interface {
	Expr
	Spell() Token
	Args() []Expr
}

func SpellCallString(s SpellCall) string {

	b := builder{}

	b.add(0, "[SpellCall]")
	b.add(2, s.Spell().String())

	b.newline()
	b.add(1, "Args:")
	for _, a := range s.Args() {
		b.newline()
		b.add(2, a.String())
	}

	return b.String()
}

type Exists interface {
	Expr
	Subject() Expr
}

func ExistsString(e Exists) string {

	b := builder{}

	b.add(0, "[Exists]")
	b.newline()
	b.add(1, e.Subject().String())

	return b.String()
}
