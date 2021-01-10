package ast

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

// Abstract node types
type (
	Node interface {
		NodeType() NodeType
		Snippet() scroll.Snippet
	}

	Expr interface {
		Node
		expr()
	}

	Stmt interface {
		Node
		stmt()
	}
)

// Concrete node types
type (
	node struct {
		nt NodeType
		sn scroll.Snippet
	}

	Ident struct {
		node
		lx token.Lexeme
	}

	Lit struct {
		node
		lx token.Lexeme
	}

	Define struct {
		node
		id Ident
		ex Expr
	}

	Assign struct {
		node
		id Ident
		ex Expr
	}
)

func (n node) NodeType() NodeType      { return n.nt }
func (n node) Snippet() scroll.Snippet { return n.sn }

func (n Ident) expr() {}
func (n Lit) expr()   {}

func (n Assign) stmt() {}
func (n Define) stmt() {}

func (n Ident) Lexeme() token.Lexeme { return n.lx }
func (n Lit) Lexeme() token.Lexeme   { return n.lx }

func (n Assign) Ident() Ident { return n.id }
func (n Assign) Expr() Expr   { return n.ex }
func (n Define) Ident() Ident { return n.id }
func (n Define) Expr() Expr   { return n.ex }

func _enforceTypes() {
	var _ Expr = Ident{}
	var _ Expr = Lit{}

	var _ Stmt = Define{}
	var _ Stmt = Assign{}
}
