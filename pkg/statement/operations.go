package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Operation interface {
	Expression

	Left() Expression

	Operator() token.Token

	Right() Expression
}

// baseOp is the base structure embeded in Operations to avoid redefining
// idententical methods.
type baseOp struct {
	name  string
	left  Expression
	op    token.Token
	right Expression
}

func (bo baseOp) Left() Expression {
	return bo.left
}

func (bo baseOp) Operator() token.Token {
	return bo.op
}

func (bo baseOp) Right() Expression {
	return bo.right
}

func (bo baseOp) Token() token.Token {
	return bo.op
}

func (bo baseOp) String(i int) string {

	str := indent(i) + "[" + bo.name + "] " + bo.op.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += bo.left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += bo.right.String(i + 2)

	return str
}

type Arithmetic struct {
	baseOp
}

type Relation struct {
	baseOp
}

type Equality struct {
	baseOp
}

type Logic struct {
	baseOp
}

// NewMathExpression returns either an Arithmetic, Relation, or Logic
// expression depending on the operator token type.
func NewMathExpression(left Expression, op token.Token, right Expression) Expression {

	switch op.Type {
	case token.ADD,
		token.SUBTRACT,
		token.MULTIPLY,
		token.DIVIDE,
		token.REMAINDER:
		return Arithmetic{baseOp{`Arithmetic`, left, op, right}}

	case token.LESS_THAN,
		token.LESS_THAN_OR_EQUAL,
		token.MORE_THAN,
		token.MORE_THAN_OR_EQUAL:
		return Relation{baseOp{`Relation`, left, op, right}}

	case token.EQUAL, token.NOT_EQUAL:
		return Equality{baseOp{`Equality`, left, op, right}}

	case token.AND, token.OR:
		return Logic{baseOp{`Logic`, left, op, right}}
	}

	panic(string("[NewMathExpression] Unknown operator '" + string(op.Type) + "'"))
}
