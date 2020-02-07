package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// Operand is an expression that always returns a Value that can be used as an
// operand in an operation.
type Operand Expr

// operation represents an algebraic operation.
type operation struct {
	left     Operand
	operator token.Token
	right    Operand
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
	Operand
	operation
}

// String satisfies the Expr interface.
func (mop mathOperation) String() string {
	return mop.toString("Math")
}

// Eval satisfies the Expr interface.
func (op mathOperation) Eval(ctx Context) (_ Value) {

	var left float64
	var right float64
	var answer float64

	lv := op.left.Eval(ctx)
	rv := op.right.Eval(ctx)

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

	switch op.operator.Kind {
	case token.MATH_ADD:
		answer = left + right
	}

	return Value{
		k: REAL,
		v: answer,
	}
}
