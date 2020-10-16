package tree

import (
	"github.com/PaulioRandall/scarlet-go/number"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
)

// TODO: Model on https://github.com/golang/go/blob/a517c3422e808ae51533a0700e05d59e8a799136/src/go/ast/ast.go

type (

	// Node represents a node in a syntax graph or tree.
	Node interface {
		Pos() position.Snippet
		node()
	}

	// Assignee is a Node that represents something that can have value bound
	// to it, i.e. an identifier.
	Assignee interface {
		Node
		assignee()
	}

	// Expr (Expression) is a Node that represents a traditional programmers
	// expression, i.e. a statement that always returns a single result.
	Expr interface {
		Node
		expr()
	}

	// Literal is a Node that represents a literal value such as a bool, a number
	// or a string.
	Literal interface {
		Node
		literal()
	}

	// Stat (Statement) is a Node representing a traditional programmers
	// statement.
	Stat interface {
		Node
		stat()
	}
)

type (
	// Ident Node is an Expr representing an identifier.
	Ident struct {
		position.Snippet
		Val string // Identifier name as defined in source
	}

	// VoidLit Node is an Expr representing a void.
	VoidLit struct {
		position.Snippet
	}

	// BoolLit Node is an Expr representing a literal boolean.
	BoolLit struct {
		position.Snippet
		Val bool
	}

	// NumLit Node is an Expr representing a literal number.
	NumLit struct {
		position.Snippet
		Val number.Number
	}

	// StrLit Node is an Expr representing a literal string.
	StrLit struct {
		position.Snippet
		Val string
	}

	// SingleAssign Node is a Stat representing a single assignment.
	SingleAssign struct {
		position.Snippet
		Left  Assignee
		Infix position.Snippet
		Right Expr
	}

	// MultiAssign Node is a Stat representing a multiple assignment.
	MultiAssign struct {
		position.Snippet
		Left  []Assignee // Ordered left to right
		Infix position.Snippet
		Right []Expr // Ordered left to right
	}

	// BinaryExpr Node is an Expr representing an operation with two operands.
	BinaryExpr struct {
		position.Snippet
		Left  Expr
		Op    token.Token
		OpPos position.Snippet
		Right Expr
	}
)

func (n Ident) Pos() position.Snippet        { return n.Snippet }
func (n VoidLit) Pos() position.Snippet      { return n.Snippet }
func (n BoolLit) Pos() position.Snippet      { return n.Snippet }
func (n NumLit) Pos() position.Snippet       { return n.Snippet }
func (n StrLit) Pos() position.Snippet       { return n.Snippet }
func (n SingleAssign) Pos() position.Snippet { return n.Snippet }
func (n MultiAssign) Pos() position.Snippet  { return n.Snippet }
func (n BinaryExpr) Pos() position.Snippet   { return n.Snippet }

func (n Ident) node()        {}
func (n VoidLit) node()      {}
func (n BoolLit) node()      {}
func (n NumLit) node()       {}
func (n StrLit) node()       {}
func (n SingleAssign) node() {}
func (n MultiAssign) node()  {}
func (n BinaryExpr) node()   {}

func (n Ident) assignee() {}

func (n Ident) expr()      {}
func (n VoidLit) expr()    {}
func (n BoolLit) expr()    {}
func (n NumLit) expr()     {}
func (n StrLit) expr()     {}
func (n BinaryExpr) expr() {}

func (n BoolLit) literal() {}
func (n NumLit) literal()  {}
func (n StrLit) literal()  {}

func (n SingleAssign) stat() {}
func (n MultiAssign) stat()  {}
