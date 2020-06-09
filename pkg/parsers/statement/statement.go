package statement

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Statement interface {
	Expression
}

type NewExpression interface {
	Expression
	Kind() Kind
}

type Expression interface {
	Snippet
}

type Snippet interface {
	fmt.Stringer
	Begin() (line, col int)
	End() (line, col int)
}

type Void interface {
	Tk() Token
}

func VoidString(v Void) string {

	b := builder{}

	b.add(0, "[Void] ")
	b.addToken(0, v.Tk())

	return b.String()
}

type Identifier interface {
	Tk() Token
}

func IdentifierString(id Identifier) string {

	b := builder{}

	b.add(0, "[Identifier] ")
	b.addToken(0, id.Tk())

	return b.String()
}

type Literal interface {
	Tk() Token
}

func LiteralString(l Literal) string {

	b := builder{}

	b.add(0, "[Literal] ")
	b.addToken(0, l.Tk())

	return b.String()
}

type ListAccessor interface {
	ID() Expression
	Index() Expression
}

func ListAccessorString(l ListAccessor) string {

	b := builder{}

	b.add(0, "[ListAccessor] ")

	b.newline()
	b.add(1, "ID: ")
	b.newline()
	b.add(2, l.ID().String())

	b.newline()
	b.add(1, "Index: ")
	b.newline()
	b.add(2, l.Index().String())

	return b.String()
}

type ListConstructor interface {
	Open() Token
	Close() Token
	Items() []Expression
}

func ListConstructorString(l ListConstructor) string {

	b := builder{}

	b.add(0, "[List] ")

	for _, item := range l.Items() {
		b.newline()
		b.add(1, item.String())
	}

	return b.String()
}

type Negation interface {
	Expr() Expression
}

func NegationString(n Negation) string {

	b := builder{}

	b.add(0, "[Negation]")

	b.newline()
	b.add(1, n.Expr().String())

	return b.String()
}

type Assignment interface {
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
