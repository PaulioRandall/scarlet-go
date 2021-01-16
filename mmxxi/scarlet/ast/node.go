package ast

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

type ValType int

const (
	T_UNDEFINED ValType = iota
	T_INFER             // Inferred
	//T_USER // User defined
	T_BOOL
	T_NUM
	T_STR
	//T_LIST
	//T_MAP
	//T_EFUNC
	//T_FUNC
)

func (vt ValType) String() string {
	switch vt {
	case T_INFER:
		return "T_INFER"
	case T_BOOL:
		return "T_BOOL"
	case T_NUM:
		return "T_NUM"
	case T_STR:
		return "T_STR"
	}

	return "UNDEFINED"
}

// Abstract node types
type (
	Node interface {
		Snippet() scroll.Snippet
		node()
	}

	Expr interface {
		Node
		ValueType() ValType
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

	BaseExpr struct {
		Base
		ValType ValType
	}

	Ident struct {
		BaseExpr
		Lex token.Lexeme
	}

	BoolLit struct {
		BaseExpr
		Val bool
	}

	NumLit struct {
		BaseExpr
		Val float64
	}

	StrLit struct {
		BaseExpr
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

func (n BaseExpr) ValueType() ValType { return n.ValType }
func (n BaseExpr) expr()              {}

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
