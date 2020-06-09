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
