package parser

import (
	"github.com/PaulioRandall/scarlet-go/number"
	"github.com/PaulioRandall/scarlet-go/token2/position"
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

	// Stat (Statement) is a Node representing a traditional programmers
	// statement.
	Stat interface {
		Node
		stat() // Constrains assignment by statement nodes only
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

	/*
		// BinaryOp Node is an Expr representing a binary operation.
		BinaryOp struct {
			Left   Expr
			Op     token.Token
			OpSnip position.Snippet
			Right  Expr
		}

		// BoolOp Node is an Expr representing a boolean operation.
		BoolOp struct {
			BinaryOp
		}

		// ArithOp Node is an Expr representing an arithmetic operation.
		ArithOp struct {
			BinaryOp
		}

		// SpellCall Node is an Expr representing a spell call.
		SpellCall struct {
			Prefix   position.Snippet
			NameSnip position.Snippet
			Name     string // Identifier name as defined in source
			LParen   position.Snippet
			Args     []Expr // Ordered left to right
			RParen   position.Snippet
		}

		// ExprStat Node is a Stat representing an Expr.
		ExprStat struct {
			Expr Expr
		}
	*/

	// SingleAssign Node is a Stat representing a single assignment.
	SingleAssign struct {
		Left  Expr
		Infix position.Snippet
		Right Expr
	}

	// MultiAssign Node is a Stat representing a multiple assignment.
	MultiAssign struct {
		Left  []Expr // Ordered left to right
		Infix position.Snippet
		Right []Expr // Ordered left to right
	}

/*
	// Block Node is a Stat representing a block of code with its own scope.
	Block struct {
		LCurly position.Snippet
		Stats  []Stat
		RCurly position.Snippet
	}

	// Guard Node is a Stat representing a block of code that is conditionally
	// executed, i.e. 'if' statement.
	Guard struct {
		LSquare position.Snippet
		Cond    Expr
		RSquare position.Snippet
		Body    Block
	}

	// WhileLoop Node is a Stat representing a while loop.
	WhileLoop struct {
		Loop  position.Snippet
		Guard Guard
	}
*/
)

func (n Ident) expr()   {}
func (n VoidLit) expr() {}
func (n BoolLit) expr() {}
func (n NumLit) expr()  {}
func (n StrLit) expr()  {}

/*
func (n BoolOp) expr()    {}
func (n ArithOp) expr()   {}
func (n SpellCall) expr() {}
*/

func (n SingleAssign) stat() {}
func (n MultiAssign) stat()  {}

/*
func (n ExprStat) stat()    {}
func (n Block) stat()       {}
func (n Guard) stat()       {}
func (n WhileLoop) stat()   {}
*/

func (n Ident) Snip() position.Snippet   { return n.Snippet }
func (n VoidLit) Snip() position.Snippet { return n.Snippet }
func (n BoolLit) Snip() position.Snippet { return n.Snippet }
func (n NumLit) Snip() position.Snippet  { return n.Snippet }
func (n StrLit) Snip() position.Snippet  { return n.Snippet }

func (n SingleAssign) Snip() position.Snippet {
	return position.SuperSnippet(n.Left.Snip(), n.Right.Snip())
}

func (n MultiAssign) Snip() position.Snippet {
	return position.SuperSnippet(
		exprListSnippet(n.Left),
		exprListSnippet(n.Right),
	)
}

/*
func (n BinaryOp) Snip() position.Snippet {
	return position.SuperSnippet(n.Left.Snip(), n.Right.Snip())
}
func (n SpellCall) Snip() position.Snippet {
	return position.SuperSnippet(n.Prefix, n.RParen)
}
func (n ExprStat) Snip() position.Snippet { return n.Expr.Snip() {
func (n Block) Snip() position.Snippet {
	return position.SuperSnippet(n.LCurly, n.RCurly)
}
func (n Guard) Snip() position.Snippet {
	return position.SuperSnippet(n.LSquare, n.Body.Snip())
}
func (n WhileLoop) Snip() position.Snippet {
	return position.SuperSnippet(n.Loop, n.Guard.Snip())
}
*/

func exprListSnippet(nodes []Expr) position.Snippet {
	var r position.Snippet
	for i, s := range nodes {
		if i == 0 {
			r = s.Snip()
		} else {
			r = position.SuperSnippet(r, s.Snip())
		}
	}
	return r
}
