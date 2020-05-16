package alpha

import (
	"github.com/shopspring/decimal"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalOperation(ctx *alphaContext, op st.Operation) result {

	tk := op.Operator

	switch tk.Type {
	case token.ADD:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return numberLiteral(left.Add(right))

	case token.SUBTRACT:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return numberLiteral(left.Sub(right))

	case token.MULTIPLY:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return numberLiteral(left.Mul(right))

	case token.DIVIDE:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return numberLiteral(left.Div(right))

	case token.REMAINDER:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return numberLiteral(left.Mod(right))

	case token.LESS_THAN:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return boolLiteral(left.LessThan(right))

	case token.MORE_THAN:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return boolLiteral(left.GreaterThan(right))

	case token.LESS_THAN_OR_EQUAL:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return boolLiteral(left.LessThanOrEqual(right))

	case token.MORE_THAN_OR_EQUAL:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return boolLiteral(left.GreaterThanOrEqual(right))

	case token.AND:
		left, right := evalBools(ctx, op.Left, op.Right)
		return boolLiteral(left && right)

	case token.OR:
		left, right := evalBools(ctx, op.Left, op.Right)
		return boolLiteral(left || right)

	case token.EQUAL:
		left, right := evalValues(ctx, op.Left, op.Right)
		return boolLiteral(equal(left, right))

	case token.NOT_EQUAL:
		left, right := evalValues(ctx, op.Left, op.Right)
		return boolLiteral(!equal(left, right))
	}

	panic(err("evalOperation", tk, "Unknown operation type"))
}

func evalValues(ctx *alphaContext, left, right st.Expression) (result, result) {
	l := evalExpression(ctx, left)
	r := evalExpression(ctx, right)

	if _, ok := l.(voidLiteral); ok {
		panic(err("evalValues", left.Token(),
			"Left side evaluated to voidLiteral, but it's not allowed here"))
	}

	if _, ok := r.(voidLiteral); ok {
		panic(err("evalValues", right.Token(),
			"Right side evaluated to voidLiteral, but it's not allowed here"))
	}

	return l, r
}

func equal(left, right result) bool {

	l, ok := left.(numberLiteral)
	if !ok {
		return left == right
	}

	r, ok := right.(numberLiteral)
	if !ok {
		return left == right
	}

	return decimal.Decimal(l).Equal(decimal.Decimal(r))
}

func evalNumbers(ctx *alphaContext,
	left, right st.Expression,
) (decimal.Decimal, decimal.Decimal) {

	return evalNumber(ctx, left), evalNumber(ctx, right)
}

func evalNumber(ctx *alphaContext, ex st.Expression) decimal.Decimal {

	v := evalExpression(ctx, ex)
	v = expectOneValue(v, ex.Token())

	if v, ok := v.(numberLiteral); ok {
		return decimal.Decimal(v)
	}

	panic(err("evalNumber", ex.Token(), "Expected Number as result"))
}

func evalBools(ctx *alphaContext, left, right st.Expression) (bool, bool) {
	return evalBool(ctx, left), evalBool(ctx, right)
}

func evalBool(ctx *alphaContext, ex st.Expression) bool {
	if v, ok := evalExpression(ctx, ex).(boolLiteral); ok {
		return bool(v)
	}

	panic(err("EvalBool", ex.Token(), "Expected Bool as result"))
}
