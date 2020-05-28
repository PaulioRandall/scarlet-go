package alpha

import (
	"github.com/shopspring/decimal"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func evalOperation(ctx *alphaContext, op Operation) result {

	tk := op.Operator

	switch tk.Morpheme() {
	case ADD:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return numberLiteral(left.Add(right))

	case SUBTRACT:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return numberLiteral(left.Sub(right))

	case MULTIPLY:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return numberLiteral(left.Mul(right))

	case DIVIDE:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return numberLiteral(left.Div(right))

	case REMAINDER:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return numberLiteral(left.Mod(right))

	case LESS_THAN:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return boolLiteral(left.LessThan(right))

	case MORE_THAN:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return boolLiteral(left.GreaterThan(right))

	case LESS_THAN_OR_EQUAL:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return boolLiteral(left.LessThanOrEqual(right))

	case MORE_THAN_OR_EQUAL:
		left, right := evalNumber(ctx, op.Left), evalNumber(ctx, op.Right)
		return boolLiteral(left.GreaterThanOrEqual(right))

	case AND:
		left, right := evalBool(ctx, op.Left), evalBool(ctx, op.Right)
		return boolLiteral(left && right)

	case OR:
		left, right := evalBool(ctx, op.Left), evalBool(ctx, op.Right)
		return boolLiteral(left || right)

	case EQUAL:
		left, right := evalExpression(ctx, op.Left), evalExpression(ctx, op.Right)
		return boolLiteral(equal(left, right))

	case NOT_EQUAL:
		left, right := evalExpression(ctx, op.Left), evalExpression(ctx, op.Right)
		return boolLiteral(!equal(left, right))
	}

	err.Panic("Unknown operation", err.At(tk))
	return nil
}

func evalNegation(ctx *alphaContext, n Negation) result {

	v := evalExpression(ctx, n.Expr)

	if num, ok := v.(numberLiteral); ok {
		d := decimal.Decimal(num)
		return numberLiteral(d.Neg())
	}

	if b, ok := v.(boolLiteral); ok {
		return boolLiteral(!bool(b))
	}

	err.Panic("Not a numeric or boolean expression", err.At(n.Token()))
	return nil
}

func equal(left, right result) bool {

	nl, lok := left.(numberLiteral)
	nr, rok := right.(numberLiteral)

	if lok && rok {
		return decimal.Decimal(nl).Equal(decimal.Decimal(nr))
	}

	return left == right
}

func evalNumber(ctx *alphaContext, ex Expression) decimal.Decimal {

	v := evalExpression(ctx, ex)
	v = expectOneValue(v, ex.Token())
	n, ok := v.(numberLiteral)

	if !ok {
		err.Panic("Not a numeric expression", err.At(ex.Token()))
	}

	return decimal.Decimal(n)
}

func evalBool(ctx *alphaContext, ex Expression) bool {

	v, ok := evalExpression(ctx, ex).(boolLiteral)

	if !ok {
		err.Panic("Not a boolean expression", err.At(ex.Token()))
	}

	return bool(v)
}
