package parser

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/token"
)

// opKind represents the kind of an operation.
type opKind string

const (
	NOT_OPERATOR opKind = `NOT_OPERATOR`
	ARITHMETIC   opKind = `Arithmetic`
	COMPARISON   opKind = `Comparison`
	BOOLEAN      opKind = `Boolean`
)

// identifyOperatorKind returns the operation kind of the token. If the token
// is not a known operator then NOT_OPERATOR is returned.
func (p *Parser) identifyOperatorKind(k token.Kind) opKind {
	switch k {
	case token.ADD, token.SUBTRACT, token.MULTIPLY, token.DIVIDE, token.MOD:
		return ARITHMETIC
	case token.EQU, token.NEQ:
		return COMPARISON
	case token.LT, token.LT_OR_EQU, token.MT, token.MT_OR_EQU:
		return COMPARISON
	case token.AND, token.OR:
		return BOOLEAN
	default:
		return NOT_OPERATOR
	}
}

// parseOperation parses an arithmetic operation.
func (p *Parser) parseOperation(left Expr) Expr {

	op := p.take()
	kind := p.identifyOperatorKind(op.Kind)
	right := p.parseExpr()

	return operation{
		kind:     kind,
		left:     left,
		operator: op,
		right:    right,
	}
}

// operation represents an algebraic operation.
type operation struct {
	kind     opKind
	left     Expr
	operator token.Token
	right    Expr
}

// String satisfies the Expr interface.
func (op operation) String() (s string) {

	s += string(op.kind) + " Operation (" + op.operator.String() + ")"

	s += "\n\tLeft"
	s += "\n\t\t" + strings.ReplaceAll(op.left.String(), "\n", "\n\t\t")

	s += "\n\tRight"
	s += "\n\t\t" + strings.ReplaceAll(op.right.String(), "\n", "\n\t\t")

	return
}

// Eval satisfies the Expr interface.
func (op operation) Eval(ctx Context) (_ Value) {
	switch op.kind {
	case ARITHMETIC, COMPARISON:
		return op.evalNumeric(ctx)
	case BOOLEAN:
		return op.evalBoolean(ctx)
	default:
		panic(bard.NewHorror(op.operator, nil, "SANITY CHECK! Unknown operator"))
	}
}

// evalNumeric evaluates an arithmetic or comparison ooeration.
func (op operation) evalNumeric(ctx Context) Value {

	lv := op.left.Eval(ctx)
	rv := op.right.Eval(ctx)

	// TODO: Check both operands are numeric

	if op.operator.Kind != token.DIVIDE && lv.k == INT && rv.k == INT {
		return op.evalInt(lv.v.(int64), rv.v.(int64))
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

	return op.evalReal(left, right)
}

// evalInt evaluates an operation involving two integer operands.
func (op operation) evalInt(l, r int64) (v Value) {

	if op.kind == ARITHMETIC {
		v.k = INT

		switch op.operator.Kind {
		case token.ADD:
			v.v = l + r
		case token.SUBTRACT:
			v.v = l - r
		case token.MULTIPLY:
			v.v = l * r
		case token.MOD:
			v.v = l % r
		default:
			panic(bard.NewHorror(op.operator, nil,
				"SANITY CHECK! Unknown integer arithmetic operator",
			))
		}

		return
	}

	v.k = BOOL

	switch op.operator.Kind {
	case token.EQU:
		v.v = l == r
	case token.NEQ:
		v.v = l != r
	case token.LT:
		v.v = l < r
	case token.LT_OR_EQU:
		v.v = l <= r
	case token.MT:
		v.v = l > r
	case token.MT_OR_EQU:
		v.v = l >= r
	default:
		panic(bard.NewHorror(op.operator, nil,
			"SANITY CHECK! Unknown int boolean operator",
		))
	}

	return
}

// evalReal evaluates an operation involving two real operands.
func (op operation) evalReal(l, r float64) (v Value) {

	if op.kind == ARITHMETIC {
		v.k = REAL

		switch op.operator.Kind {
		case token.ADD:
			v.v = l + r
		case token.SUBTRACT:
			v.v = l - r
		case token.MULTIPLY:
			v.v = l * r
		case token.DIVIDE:
			if r == 0 {
				panic(bard.NewHorror(op.operator, nil, "Cannot divide by zero"))
			}
			v.v = l / r
		default:
			panic(bard.NewHorror(op.operator, nil,
				"SANITY CHECK! Unknown real arithmetic operator",
			))
		}

		return
	}

	v.k = BOOL

	switch op.operator.Kind {
	case token.EQU:
		v.v = l == r
	case token.NEQ:
		v.v = l != r
	case token.LT:
		v.v = l < r
	case token.LT_OR_EQU:
		v.v = l <= r
	case token.MT:
		v.v = l > r
	case token.MT_OR_EQU:
		v.v = l >= r
	default:
		panic(bard.NewHorror(op.operator, nil,
			"SANITY CHECK! Unknown real boolean operator",
		))
	}

	return
}

// evalNumeric evaluates an arithmetic or comparison ooeration.
func (op operation) evalBoolean(ctx Context) (v Value) {

	lv := op.left.Eval(ctx)
	rv := op.right.Eval(ctx)

	// TODO: Check both operands are booleans

	l := lv.v.(bool)
	r := rv.v.(bool)

	v.k = BOOL

	switch op.operator.Kind {
	case token.AND:
		v.v = l && r
	case token.OR:
		v.v = l || r
	default:
		panic(bard.NewHorror(op.operator, nil,
			"SANITY CHECK! Unknown bool operator",
		))
	}

	return
}
