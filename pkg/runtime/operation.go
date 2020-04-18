package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func EvalArithmetic(ctx *Context, a statement.Arithmetic) Value {

	leftExpr := EvalExpression(ctx, a.Left())
	leftInt, isLeftInt := leftExpr.(Int)
	leftFloat, isLeftFloat := leftExpr.(Float)

	rightExpr := EvalExpression(ctx, a.Right())
	rightInt, isRightInt := rightExpr.(Int)
	rightFloat, isRightFloat := rightExpr.(Float)

	op := a.Operator()

	switch {
	case isLeftFloat && isRightFloat:
		return floatArithmetic(op, leftFloat, rightFloat)
	case isLeftInt && isRightFloat:
		return floatArithmetic(op, leftInt.ToFloat(), rightFloat)
	case isLeftFloat && isRightInt:
		return floatArithmetic(op, leftFloat, rightInt.ToFloat())
	case isLeftInt && isRightInt:
		if op.Type == token.DIVIDE {
			return floatArithmetic(op, leftInt.ToFloat(), rightInt.ToFloat())
		}

		return intArithmetic(op, leftInt, rightInt)
	}

	if !isLeftInt && !isLeftFloat {
		panic(err("EvalArithmetic", a.Left().Token(), "Expected Int or Float"))
	} else {
		panic(err("EvalArithmetic", a.Right().Token(), "Expected Int or Float"))
	}
}

func floatArithmetic(op token.Token, a, b Float) Value {

	x := float64(a)
	y := float64(b)

	switch op.Type {
	case token.ADD:
		return Float(x + y)
	case token.SUBTRACT:
		return Float(x - y)
	case token.MULTIPLY:
		return Float(x * y)
	case token.DIVIDE:
		return Float(x / y)
	}

	panic(err("floatArithmetic", op, "Unknown arithmetic Float operator"))
}

func intArithmetic(op token.Token, a, b Int) Value {

	x := int64(a)
	y := int64(b)

	switch op.Type {
	case token.ADD:
		return Int(x + y)
	case token.SUBTRACT:
		return Int(x - y)
	case token.MULTIPLY:
		return Int(x * y)
	case token.REMAINDER:
		return Int(x % y)
	}

	panic(err("intArithmetic", op, "Unknown arithmetic Int operator"))
}

func EvalLogic(ctx *Context, l statement.Logic) Value {

	op := l.Operator()

	leftExpr := EvalExpression(ctx, l.Left())
	left, ok := leftExpr.(Bool)
	if !ok {
		panic(err("EvalLogic", op, "Expected Bool value on left"))
	}

	rightExpr := EvalExpression(ctx, l.Right())
	right, ok := rightExpr.(Bool)
	if !ok {
		panic(err("EvalLogic", op, "Expected Bool value on right"))
	}

	a := bool(left)
	b := bool(right)

	switch op.Type {
	case token.AND:
		return Bool(a && b)
	case token.OR:
		return Bool(a || b)
	}

	panic(err("EvalLogic", op, "Unknown logical operator"))
}

func EvalRelation(ctx *Context, r statement.Relation) Value {

	op := r.Operator()

	resolve := func(expr statement.Expression) float64 {
		v := EvalExpression(ctx, expr)

		if vInt, ok := v.(Int); ok {
			return float64(vInt.ToFloat())
		}

		if vFloat, ok := v.(Float); ok {
			return float64(vFloat)
		}

		panic(err("EvalRelation", op, "Expected Int or Float value on left"))
	}

	left := resolve(r.Left())
	right := resolve(r.Right())

	switch op.Type {
	case token.LESS_THAN:
		return Bool(left < right)
	case token.MORE_THAN:
		return Bool(left > right)
	case token.LESS_THAN_OR_EQUAL:
		return Bool(left <= right)
	case token.MORE_THAN_OR_EQUAL:
		return Bool(left >= right)
	}

	panic(err("EvalRelation", op, "Unknown relational operator"))
}

func EvalEquality(ctx *Context, eq statement.Equality) Value {

	op := eq.Operator()
	left := EvalExpression(ctx, eq.Left())
	right := EvalExpression(ctx, eq.Right())

	intToFloat := func(v Value) Value {
		if vInt, ok := v.(Int); ok {
			return vInt.ToFloat()
		}
		return v
	}

	left = intToFloat(left)
	right = intToFloat(right)

	switch op.Type {
	case token.EQUAL:
		return Bool(left == right)
	case token.NOT_EQUAL:
		return Bool(left != right)
	}

	panic(err("EvalEquality", op, "Unknown equality operator"))
}
