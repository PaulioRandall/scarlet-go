package statement

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Snippet
	Kind() Kind
}

type Snippet interface {
	fmt.Stringer
	Begin() (line, col int)
	End() (line, col int)
}

type Void interface {
	Expression
	Tk() Token
}

func VoidString(v Void) string {

	b := builder{}

	b.add(0, "[Void] ")
	b.addToken(0, v.Tk())

	return b.String()
}

type Identifier interface {
	Expression
	Tk() Token
}

func IdentifierString(id Identifier) string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.addToken(0, id.Tk())

	return b.String()
}

type Literal interface {
	Expression
	Tk() Token
}

func LiteralString(l Literal) string {

	b := builder{}

	b.add(0, "[Literal] ")
	b.addToken(0, l.Tk())

	return b.String()
}

type CollectionAccessor interface {
	Expression
	Collection() Expression
	Key() Expression
}

func CollectionAccessorString(c CollectionAccessor) string {

	b := builder{}

	b.add(0, "[CollectionAccessor] ")

	b.newline()
	b.add(1, "Collection: ")
	b.newline()
	b.add(2, c.Collection().String())

	b.newline()
	b.add(1, "Key: ")
	b.newline()
	b.add(2, c.Key().String())

	return b.String()
}

type Negation interface {
	Expression
	Expr() Expression
}

func NegationString(n Negation) string {

	b := builder{}

	b.add(0, "[Negation]")

	b.newline()
	b.add(1, n.Expr().String())

	return b.String()
}

type Operation interface {
	Expression
	Operator() Token
	Left() Expression
	Right() Expression
}

func OperationString(o Operation) string {

	b := builder{}

	b.add(0, "[Operation] ")
	b.addToken(0, o.Operator())

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

type Assignment interface {
	Expression
	Target() Expression
	Source() Expression
}

func AssignmentString(a Assignment) string {

	b := builder{}

	b.add(0, "[Assignment] ")

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

type AssignmentBlock interface {
	Expression
	Assignments() []Assignment
}

func AssignmentBlockString(bk AssignmentBlock) string {

	b := builder{}

	for _, a := range bk.Assignments() {
		b.add(0, a.String())
	}

	return b.String()
}

type ExpressionFunction interface {
	Expression
	Inputs() []Token
	Expr() Expression
}

func ExpressionFunctionString(e ExpressionFunction) string {

	b := builder{}

	b.add(0, "[ExpressionFunction] ")

	if len(e.Inputs()) > 0 {

		b.newline()
		b.add(1, "Inputs: ")

		for _, in := range e.Inputs() {
			b.newline()
			b.addToken(2, in)
		}
	}

	b.newline()
	b.add(1, e.Expr().String())

	return b.String()
}

type Parameters interface {
	Expression
	Inputs() []Token
	Outputs() []Token
}

func ParametersString(p Parameters) string {

	b := builder{}

	b.add(0, "[Parameters] ")

	if len(p.Inputs()) > 0 {

		b.newline()
		b.add(1, "Inputs: ")

		for _, in := range p.Inputs() {
			b.newline()
			b.addToken(2, in)
		}
	}

	if len(p.Outputs()) > 0 {

		b.newline()
		b.add(1, "Outputs: ")

		for _, out := range p.Outputs() {
			b.newline()
			b.addToken(2, out)
		}
	}

	return b.String()
}

type Function interface {
	Expression
	Params() Parameters
	Body() Block
}

func FunctionString(f Function) string {

	b := builder{}

	b.add(0, "[Function] ")

	b.newline()
	b.add(1, ParametersString(f.Params()))

	b.newline()
	b.add(1, BlockString(f.Body()))

	return b.String()
}

type FunctionCall interface {
	Expression
	Function() Expression
	Arguments() []Expression
}

func FunctionCallString(f FunctionCall) string {

	b := builder{}

	b.add(0, "[FunctionCall]")

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
	Expression
	Stats() []Expression
}

func BlockString(bk Block) string {

	b := builder{}

	for _, a := range bk.Stats() {
		b.add(0, a.String())
	}

	return b.String()
}

type Watch interface {
	Expression
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
		b.addToken(2, id)
	}

	b.newline()
	b.add(1, BlockString(w.Body()))

	return b.String()
}

type Guard interface {
	Expression
	Condition() Expression
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
	Guard
}

func WhenCaseString(mc WhenCase) string {

	b := builder{}

	b.add(0, "[WhenCase] ")

	b.newline()
	b.add(1, "Condition:")
	b.newline()
	b.add(2, mc.Condition().String())

	b.newline()
	b.add(1, "Body:")
	b.newline()
	b.add(2, BlockString(mc.Body()))

	return b.String()
}

type When interface {
	Expression
	Initialiser() Assignment
	Cases() []WhenCase
}

func WhenString(m When) string {

	b := builder{}

	b.add(0, "[When]")

	b.newline()
	b.add(1, "Initialiser:")
	b.newline()
	b.add(1, AssignmentString(m.Initialiser()))

	for _, mc := range m.Cases() {
		b.newline()
		b.add(2, WhenCaseString(mc))
	}

	return b.String()
}

type Loop interface {
	Expression
	Initialiser() Assignment
	Guard() Guard
}

func LoopString(l Loop) string {

	b := builder{}

	b.add(0, "[Loop]")

	b.newline()
	b.add(1, "Initialiser:")
	b.newline()
	b.add(2, AssignmentString(l.Initialiser()))

	b.newline()
	b.add(1, "Guard:")
	b.newline()
	b.add(2, GuardString(l.Guard()))

	return b.String()
}

type SpellCall interface {
	Expression
	Spell() Token
	Arguments() []Expression
}

func SpellCallString(s SpellCall) string {

	b := builder{}

	b.add(0, "[SpellCall]")
	b.addToken(2, s.Spell())

	b.newline()
	b.add(1, "Arguments:")
	for _, a := range s.Arguments() {
		b.newline()
		b.add(2, a.String())
	}

	return b.String()
}
