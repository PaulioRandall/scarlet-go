package ast

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

// Abstract node types
type (
	Node interface {
		Snippet() scroll.Snippet
		node()
	}

	Expr interface {
		Node
		expr()
	}

	Stmt interface {
		Node
		stmt()
	}

	Literal interface {
		Expr
		literal()
	}

	Binding interface {
		Stmt
		Base() BaseBinding
	}
)

// Concrete node types
type (
	Base struct {
		Snip scroll.Snippet
	}

	Ident struct {
		Base
		Lex token.Lexeme
	}

	BoolLit struct {
		Base
		Val bool
	}

	NumLit struct {
		Base
		Val float64
	}

	StrLit struct {
		Base
		Val string
	}

	BaseBinding struct {
		Base
		Op    token.Lexeme
		Left  []Ident
		Right []Expr
	}

	Define struct {
		BaseBinding
	}

	Assign struct {
		BaseBinding
	}
)

func (n Base) Snippet() scroll.Snippet { return n.Snip }
func (n Base) node()                   {}

func (n Ident) expr()   {}
func (n BoolLit) expr() {}
func (n NumLit) expr()  {}
func (n StrLit) expr()  {}

func (n Define) stmt() {}
func (n Assign) stmt() {}

func (n BoolLit) literal() {}
func (n NumLit) literal()  {}
func (n StrLit) literal()  {}

func (n Define) Base() BaseBinding { return n.BaseBinding }
func (n Assign) Base() BaseBinding { return n.BaseBinding }

func _enforceTypes() {

	var _ Expr = Ident{}

	var _ Literal = BoolLit{}
	var _ Literal = NumLit{}
	var _ Literal = StrLit{}

	var _ Binding = Define{}
	var _ Binding = Assign{}
}
