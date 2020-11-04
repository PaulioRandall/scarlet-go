package tree

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"
)

// TODO: Model on https://github.com/golang/go/blob/a517c3422e808ae51533a0700e05d59e8a799136/src/go/ast/ast.go

type (
	// Node represents a node in a syntax graph or tree.
	Node interface {
		Pos() token.Snippet
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
		Expr
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
		token.Snippet
		Val string // Identifier name as defined in source
	}

	// VoidLit Node is an Expr representing a void.
	VoidLit struct {
		token.Snippet
	}

	// BoolLit Node is an Expr representing a literal boolean.
	BoolLit struct {
		token.Snippet
		Val bool
	}

	// NumLit Node is an Expr representing a literal number.
	NumLit struct {
		token.Snippet
		Val number.Number
	}

	// StrLit Node is an Expr representing a literal string.
	StrLit struct {
		token.Snippet
		Val string
	}

	// SingleAssign Node is a Stat representing a single assignment.
	SingleAssign struct {
		token.Snippet
		Left  Assignee
		Infix token.Snippet
		Right Expr
	}

	// MultiAssign Node is a Stat representing a multiple assignment.
	MultiAssign struct {
		token.Snippet
		Asym  bool       // True if a single spell or function call exists in Right
		Left  []Assignee // Ordered left to right
		Infix token.Snippet
		Right []Expr // Ordered left to right
	}

	// BinaryExpr Node is an Expr representing an operation with two operands.
	BinaryExpr struct {
		token.Snippet
		Left  Expr
		Op    token.Token
		OpPos token.Snippet
		Right Expr
	}

	// SpellCall Node is Expr representing a spell invocation.
	SpellCall struct {
		token.Snippet
		Name string
		Args []Expr
	}
)

func (n Ident) Pos() token.Snippet        { return n.Snippet }
func (n VoidLit) Pos() token.Snippet      { return n.Snippet }
func (n BoolLit) Pos() token.Snippet      { return n.Snippet }
func (n NumLit) Pos() token.Snippet       { return n.Snippet }
func (n StrLit) Pos() token.Snippet       { return n.Snippet }
func (n SingleAssign) Pos() token.Snippet { return n.Snippet }
func (n MultiAssign) Pos() token.Snippet  { return n.Snippet }
func (n BinaryExpr) Pos() token.Snippet   { return n.Snippet }
func (n SpellCall) Pos() token.Snippet    { return n.Snippet }

func (n Ident) node()        {}
func (n VoidLit) node()      {}
func (n BoolLit) node()      {}
func (n NumLit) node()       {}
func (n StrLit) node()       {}
func (n SingleAssign) node() {}
func (n MultiAssign) node()  {}
func (n BinaryExpr) node()   {}
func (n SpellCall) node()    {}

func (n Ident) assignee() {}

func (n Ident) expr()      {}
func (n VoidLit) expr()    {}
func (n BoolLit) expr()    {}
func (n NumLit) expr()     {}
func (n StrLit) expr()     {}
func (n BinaryExpr) expr() {}
func (n SpellCall) expr()  {}

func (n BoolLit) literal() {}
func (n NumLit) literal()  {}
func (n StrLit) literal()  {}

func (n SingleAssign) stat() {}
func (n MultiAssign) stat()  {}
func (n SpellCall) stat()    {}
