package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalOperation(ctx *alphaContext, op statement.Operation) Value {

	tk := op.Operator

	switch tk.Type {
	case token.ADD:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Number(left + right)

	case token.SUBTRACT:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Number(left - right)

	case token.MULTIPLY:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Number(left * right)

	case token.DIVIDE:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Number(left / right)

	case token.REMAINDER:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Number(int64(left) % int64(right))

	case token.LESS_THAN:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Bool(left < right)

	case token.MORE_THAN:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Bool(left > right)

	case token.LESS_THAN_OR_EQUAL:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Bool(left <= right)

	case token.MORE_THAN_OR_EQUAL:
		left, right := evalNumbers(ctx, op.Left, op.Right)
		return Bool(left >= right)

	case token.AND:
		left, right := evalBools(ctx, op.Left, op.Right)
		return Bool(left && right)

	case token.OR:
		left, right := evalBools(ctx, op.Left, op.Right)
		return Bool(left || right)

	case token.EQUAL:
		left, right := evalValues(ctx, op.Left, op.Right)
		return Bool(left == right)

	case token.NOT_EQUAL:
		left, right := evalValues(ctx, op.Left, op.Right)
		return Bool(left != right)
	}

	panic(err("evalOperation", tk, "Unknown operation type"))
}

func evalValues(ctx *alphaContext, left, right statement.Expression) (Value, Value) {
	l := evalExpression(ctx, left)
	r := evalExpression(ctx, right)

	if _, ok := l.(Void); ok {
		panic(err("evalValues", left.Token(),
			"Left side evaluated to Void, but it's not allowed here"))
	}

	if _, ok := r.(Void); ok {
		panic(err("evalValues", right.Token(),
			"Right side evaluated to Void, but it's not allowed here"))
	}

	return l, r
}

func evalNumbers(ctx *alphaContext, left, right statement.Expression) (float64, float64) {
	return evalNumber(ctx, left), evalNumber(ctx, right)
}

func evalNumber(ctx *alphaContext, ex statement.Expression) float64 {

	v := evalExpression(ctx, ex)
	v = expectOneValue(v, ex.Token())

	if v, ok := v.(Number); ok {
		return float64(v)
	}

	panic(err("evalNumber", ex.Token(), "Expected Number as result"))
}

func evalBools(ctx *alphaContext, left, right statement.Expression) (bool, bool) {
	return evalBool(ctx, left), evalBool(ctx, right)
}

func evalBool(ctx *alphaContext, ex statement.Expression) bool {
	if v, ok := evalExpression(ctx, ex).(Bool); ok {
		return bool(v)
	}

	panic(err("EvalBool", ex.Token(), "Expected Bool as result"))
}
