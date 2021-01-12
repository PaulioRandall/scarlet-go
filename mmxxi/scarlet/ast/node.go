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
	Ident struct {
		Snip scroll.Snippet
		Lex  token.Lexeme
	}

	Lit struct {
		Snip scroll.Snippet
		Lex  token.Lexeme
	}

	Assign struct {
		Snip   scroll.Snippet
		Op     token.Lexeme
		Idents []Ident
		Exprs  []Expr
	}
)

func (n Ident) NodeType() NodeType  { return IDENT }
func (n Lit) NodeType() NodeType    { return LITERAL }
func (n Assign) NodeType() NodeType { return ASSIGN }

func (n Ident) Snippet() scroll.Snippet  { return n.Snip }
func (n Lit) Snippet() scroll.Snippet    { return n.Snip }
func (n Assign) Snippet() scroll.Snippet { return n.Snip }

func (n Ident) expr() {}
func (n Lit) expr()   {}

func (n Assign) stmt() {}

func _enforceTypes() {
	var _ Expr = Ident{}
	var _ Expr = Lit{}

	var _ Stmt = Assign{}
}
