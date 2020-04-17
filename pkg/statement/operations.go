package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// Operation represents an mathematical operation, an expression with a left
// side, opertor, and right side.
type Operation interface {
	Expression

	// Left returns the expression for obtaining the left hand value.
	Left() Expression

	// Operator returns the token representing the operator.
	Operator() token.Token

	// Right returns the expression for obtaining the right hand value.
	Right() Expression

	// precedence returns the precedence of the operation so it may be compared
	// with other operations for evaluation priority.
	Precedence() int
}

// baseOp is the base structure embeded in Operations to avoid redefining
// idententical methods.
type baseOp struct {
	name  string
	left  Expression
	op    token.Token
	right Expression
}

// Left satisfies the Operation interface.
func (bo baseOp) Left() Expression {
	return bo.left
}

// Operator satisfies the Operation interface.
func (bo baseOp) Operator() token.Token {
	return bo.op
}

// Right satisfies the Operation interface.
func (bo baseOp) Right() Expression {
	return bo.right
}

// Precedence satisfies the Operation interface.
func (bo baseOp) Precedence() int {
	return Precedence(bo.op.Type)
}

// Token satisfies the Expression interface.
func (bo baseOp) Token() token.Token {
	return bo.op
}

// String satisfies the Expression interface.
func (bo baseOp) String(i int) string {

	str := indent(i) + "[" + bo.name + "] " + bo.op.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += bo.left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += bo.right.String(i + 2)

	return str
}

// Arithmetic operations are for simple maths such as ADD and MULTIPLY.
type Arithmetic struct {
	baseOp
}

// Relation operations are for relational expressions such as LESS_THAN.
type Relation struct {
	baseOp
}

// Equality operations are for EQUAL and NOT_EQUAL expressions.
type Equality struct {
	baseOp
}

// Logic operations are for boolean algebra such as AND and OR
type Logic struct {
	baseOp
}

// Precedence returns the precedences of the token type.
func Precedence(tt token.TokenType) int {
	switch tt {
	case token.MULTIPLY, token.DIVIDE, token.REMAINDER: // Multiplicative
		return 6
	case token.ADD, token.SUBTRACT: // Additive
		return 5
	case token.LESS_THAN, token.LESS_THAN_OR_EQUAL: // Relational
		fallthrough
	case token.MORE_THAN, token.MORE_THAN_OR_EQUAL: // Relational
		return 4
	case token.EQUAL, token.NOT_EQUAL: // Equality
		return 3
	case token.AND:
		return 2
	case token.OR:
		return 1
	default:
		return 0
	}
}

// NewOperation returns either an Arithmetic, Relation, Equality, or Logic
// expression depending on the operator token type.
func NewOperation(left Expression, op token.Token, right Expression) Expression {

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

	panic(string("[NewOperation] Unknown operator '" + string(op.Type) + "'"))
}
