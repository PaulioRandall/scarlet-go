package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// operation represents an algebraic operation.
type operation struct {
	left     Expr
	operator token.Token
	right    Expr
}

// toString returns the string representation of the operation.
func (op operation) toString(opType string) (s string) {

	s += opType + " Operation (" + op.operator.String() + ")\n"

	s += "\tLeft\n"
	s += "\t" + strings.ReplaceAll(op.left.String(), "\n", "\t")

	s += "\tRight\n"
	s += "\t" + strings.ReplaceAll(op.right.String(), "\n", "\t")

	return
}

// mathOperation represents a number based operation.
type mathOperation struct {
	Expr
	operation
}

// String satisfies the Expr interface.
func (mop mathOperation) String() string {
	return mop.toString("Math")
}

// Eval satisfies the Expr interface.
func (op mathOperation) Eval(ctx Context) (_ Value) {

	lv := op.left.Eval(ctx)
	rv := op.right.Eval(ctx)

	// TODO: Check both operands are numeric

	if lv.k == INT && rv.k == INT {
		return op.evalInt(
			op.operator,
			lv.v.(int64),
			rv.v.(int64),
		)
	}

	var left float64
	var right float64

	if lv.k == INT {
		left = float64(lv.v.(int64))
	} else {
		left = lv.v.(float64)
	}

	if rv.k == INT {
		right = float64(rv.v.(int64))
	} else {
		right = rv.v.(float64)
	}

	return op.evalReal(op.operator, left, right)
}

// evalInt evaluates an operation involving two integer operands.
func (op mathOperation) evalInt(tk token.Token, l, r int64) (_ Value) {

	var a int64

	switch tk.Kind {
	case token.ADD:
		a = l + r
	case token.SUBTRACT:
		a = l - r
	case token.MULTIPLY:
		a = l * r
	case token.DIVIDE:
		if r == 0 {
			panic(bard.NewHorror(tk, nil, "Cannot divide by zero"))
		}
		a = l / r
	case token.MOD:
		a = l % r
	default:
		goto COMPARISON
	}

	return Value{
		k: INT,
		v: a,
	}

COMPARISON:

	var b bool

	switch tk.Kind {
	case token.EQU:
		b = l == r
	case token.NEQ:
		b = l != r
	case token.LT:
		b = l < r
	case token.LT_OR_EQU:
		b = l <= r
	case token.MT:
		b = l > r
	case token.MT_OR_EQU:
		b = l >= r
	default:
		panic(bard.NewHorror(tk, nil, "SANITY CHECK! Unknown int math operator"))
	}

	return Value{
		k: BOOL,
		v: b,
	}
}

// evalReal evaluates an operation involving two real operands.
func (op mathOperation) evalReal(tk token.Token, l, r float64) (_ Value) {

	var a float64

	switch tk.Kind {
	case token.ADD:
		a = l + r
	case token.SUBTRACT:
		a = l - r
	case token.MULTIPLY:
		a = l * r
	case token.DIVIDE:
		if r == 0 {
			panic(bard.NewHorror(tk, nil, "Cannot divide by zero"))
		}
		a = l / r
	default:
		goto COMPARISON
	}

	return Value{
		k: REAL,
		v: a,
	}

COMPARISON:

	var b bool

	switch tk.Kind {
	case token.EQU:
		b = l == r
	case token.NEQ:
		b = l != r
	case token.LT:
		b = l < r
	case token.LT_OR_EQU:
		b = l <= r
	case token.MT:
		b = l > r
	case token.MT_OR_EQU:
		b = l >= r
	default:
		panic(bard.NewHorror(tk, nil, "SANITY CHECK! Unknown real math operator"))
	}

	return Value{
		k: BOOL,
		v: b,
	}
}

// boolOperation represents a boolean based operation.
type boolOperation struct {
	Expr
	operation
}

// String satisfies the Expr interface.
func (bop boolOperation) String() string {
	return bop.toString("Bool")
}

// Eval satisfies the Expr interface.
func (op boolOperation) Eval(ctx Context) (_ Value) {

	lv := op.left.Eval(ctx)
	rv := op.right.Eval(ctx)

	// TODO: Check both operands are booleans

	l := lv.v.(bool)
	r := rv.v.(bool)
	tk := op.operator
	var b bool

	switch tk.Kind {
	case token.AND:
		b = l && r
	case token.OR:
		b = l || r
	default:
		panic(bard.NewHorror(tk, nil, "SANITY CHECK! Unknown bool operator"))
	}

	return Value{
		k: BOOL,
		v: b,
	}
}
