package parser

import (
	"github.com/PaulioRandall/scarlet-go/number"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// TODO: Model on https://github.com/golang/go/blob/a517c3422e808ae51533a0700e05d59e8a799136/src/go/ast/ast.go

type (

	// Node represents a node in a syntax graph or tree.
	Node interface {
		Snip() position.Snippet
	}

	// Expr (Expression) is a Node that represents a traditional programmers
	// expression, i.e. a statement that always returns a single result.
	Expr interface {
		Node
		expr() // Constrains assignment by expression nodes only
	}

	// Stat (Statement) is a Node that representing a traditional programmers
	// statement.
	Stat interface {
		Node
		stat() // Constrains assignment by statement nodes only
	}

	// Assign (Assignment) is a Node that representing an assignment.
	Assign interface {
		Node
		assign() // Constrains assignment by assignment nodes only
	}
)

type (

	// Ident Node is an Expr representing an identifier.
	Ident struct {
		position.Snippet
		Value string // Identifier name as defined in source
	}

	// BoolLit Node is an Expr representing a literal boolean.
	BoolLit struct {
		position.Snippet
		Value bool
	}

	// NumLit Node is an Expr representing a literal number.
	NumLit struct {
		position.Snippet
		Value number.Number
	}

	// StrLit Node is an Expr representing a literal string.
	StrLit struct {
		position.Snippet
		Value string
	}

	// BinaryOp Node is an Expr representing a binary operation.
	BinaryOp struct {
		position.Snippet
		Left   Expr
		Op     token.Token
		OpSnip position.Snippet
		Right  Expr
	}

	// MultiAssign Node is an Assign and Stat representing a multiple assignment.
	MultiAssign struct {
		position.Snippet
		Left  []Expr
		Infix token.Token
		Right []Expr
	}
)

func (n Ident) expr()    {}
func (n BoolLit) expr()  {}
func (n NumLit) expr()   {}
func (n StrLit) expr()   {}
func (n BinaryOp) expr() {}

func (n MultiAssign) stat() {}

func (n MultiAssign) assign() {}

func (n Ident) Snip() position.Snippet       { return n.Snippet }
func (n BoolLit) Snip() position.Snippet     { return n.Snippet }
func (n NumLit) Snip() position.Snippet      { return n.Snippet }
func (n StrLit) Snip() position.Snippet      { return n.Snippet }
func (n BinaryOp) Snip() position.Snippet    { return n.Snippet }
func (n MultiAssign) Snip() position.Snippet { return n.Snippet }
