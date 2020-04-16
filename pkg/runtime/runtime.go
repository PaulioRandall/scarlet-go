package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Run(stats []statement.Statement) Context {
	ctx := Context{make(map[string]Value), nil}
	ExeStatements(&ctx, stats)
	return ctx
}

func ExeStatements(ctx *Context, stats []statement.Statement) {
	for _, s := range []statement.Statement(stats) {
		ExeStatement(ctx, s)
	}
}

func ExeStatement(ctx *Context, stat statement.Statement) {

	values := EvalExpressions(ctx, stat.Exprs)

	if stat.IDs != nil {

		if len(stat.IDs) > len(values) {
			panic(err("ExeStatement", stat.Assign,
				"Missing expression values to populate variables... have %d, want %d",
				len(stat.IDs), len(values),
			))

		} else if len(stat.IDs) < len(values) {
			panic(err("ExeStatement", stat.Assign,
				"Too many expression values to populate variables... have %d, want %d",
				len(stat.IDs), len(values),
			))
		}
	}

	for i, id := range stat.IDs {
		ctx.Set(id.Value, values[i])
	}
}

func EvalExpressions(ctx *Context, exprs []statement.Expression) []Value {

	var values []Value

	for _, expr := range exprs {
		v := EvalExpression(ctx, expr)
		values = append(values, v)
	}

	return values
}

func EvalExpression(ctx *Context, expr statement.Expression) Value {
	switch v := expr.(type) {
	case statement.Identifier:
		return ctx.Get(v.Source.Value)

	case statement.Value:
		return valueOf(v.Source)

	case statement.Arithmetic:
		return EvalArithmetic(ctx, v)

	case statement.Relation:
		return EvalRelation(ctx, v)

	case statement.Equality:
		return EvalEquality(ctx, v)

	case statement.Logic:
		return EvalLogic(ctx, v)

	case statement.List:
		return EvalList(ctx, v)
	}

	panic(err("EvalExpression", expr.Token(), "Unknown expression type"))
}

func EvalArithmetic(ctx *Context, a statement.Arithmetic) Value {

	leftExpr := EvalExpression(ctx, a.Left)
	leftInt, isLeftInt := leftExpr.(Int)
	leftFloat, isLeftFloat := leftExpr.(Float)

	rightExpr := EvalExpression(ctx, a.Right)
	rightInt, isRightInt := rightExpr.(Int)
	rightFloat, isRightFloat := rightExpr.(Float)

	switch {
	case isLeftFloat && isRightFloat:
		return floatArithmetic(a.Operator, leftFloat, rightFloat)
	case isLeftInt && isRightFloat:
		return floatArithmetic(a.Operator, leftInt.ToFloat(), rightFloat)
	case isLeftFloat && isRightInt:
		return floatArithmetic(a.Operator, leftFloat, rightInt.ToFloat())
	case isLeftInt && isRightInt:
		if a.Operator.Type == token.DIVIDE {
			return floatArithmetic(a.Operator, leftInt.ToFloat(), rightInt.ToFloat())
		}

		return intArithmetic(a.Operator, leftInt, rightInt)
	}

	if !isLeftInt && !isLeftFloat {
		panic(err("EvalArithmetic", a.Left.Token(), "Expected Int or Float"))
	} else {
		panic(err("EvalArithmetic", a.Right.Token(), "Expected Int or Float"))
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

	op := l.Operator

	leftExpr := EvalExpression(ctx, l.Left)
	left, ok := leftExpr.(Bool)
	if !ok {
		panic(err("EvalLogic", op, "Expected Bool value on left"))
	}

	rightExpr := EvalExpression(ctx, l.Right)
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

	op := r.Operator

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

	left := resolve(r.Left)
	right := resolve(r.Right)

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

	op := eq.Operator
	left := EvalExpression(ctx, eq.Left)
	right := EvalExpression(ctx, eq.Right)

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

func EvalList(ctx *Context, list statement.List) Value {
	items := EvalExpressions(ctx, list.Exprs)
	return List(items)
}
