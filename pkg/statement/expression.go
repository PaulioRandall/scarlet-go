package statement

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Expression interface {
	Token() token.Token

	String(indent int) string
}

type Identifier struct {
	Source token.Token
}

func (id Identifier) Token() token.Token {
	return id.Source
}

func (id Identifier) String(i int) string {
	return indent(i) + "[Identifier] " + id.Source.String()
}

type Value struct {
	Source token.Token
}

func (v Value) Token() token.Token {
	return v.Source
}

func (v Value) String(i int) string {
	return indent(i) + "[Value] " + v.Source.String()
}

type Arithmetic struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (a Arithmetic) Token() token.Token {
	return a.Operator
}

func (a Arithmetic) String(i int) string {

	str := indent(i) + "[Arithmetic] " + a.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += a.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += a.Right.String(i + 2)

	return str
}

type Relation struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (r Relation) Token() token.Token {
	return r.Operator
}

func (r Relation) String(i int) string {

	str := indent(i) + "[Relation] " + r.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += r.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += r.Right.String(i + 2)

	return str
}

type Equality struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (e Equality) Token() token.Token {
	return e.Operator
}

func (e Equality) String(i int) string {

	str := indent(i) + "[Equality] " + e.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += e.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += e.Right.String(i + 2)

	return str
}

type Logic struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (l Logic) Token() token.Token {
	return l.Operator
}

func (l Logic) String(i int) string {

	str := indent(i) + "[Logic] " + l.Operator.String() + newline()
	str += indent(i+1) + "Left:" + newline()
	str += l.Left.String(i+2) + newline()
	str += indent(i+1) + "Right: " + newline()
	str += l.Right.String(i + 2)

	return str
}

type List struct {
	Start token.Token
	Exprs []Expression
	End   token.Token
}

func (l List) Token() token.Token {
	return l.Start
}

func (l List) String(i int) string {

	str := indent(i) + "[List] " + l.Start.String() + newline()
	for _, ex := range l.Exprs {
		str += ex.String(i+1) + newline()
	}
	str += indent(i+1) + l.End.String()

	return str
}

type FuncCall struct {
	ID     token.Token
	Input  []token.Token
	Output []token.Token
}

func (f FuncCall) Token() token.Token {
	return f.ID
}

func (f FuncCall) String(i int) string {

	str := indent(i) + "[FuncCall] " + f.ID.String() + newline()

	str += indent(i+1) + "Input:" + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	str += indent(i+1) + "Output: " + newline()
	for _, tk := range f.Input {
		str += tk.String() + newline()
	}

	return str
}

// NewValueExpression returns either a Value or Identifier expression depending
// on the token type.
func NewValueExpression(tk token.Token) Expression {
	switch tk.Type {
	case token.ID:
		return Identifier{tk}
	default:
		return Value{tk}
	}
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
		return Arithmetic{left, op, right}

	case token.LESS_THAN,
		token.LESS_THAN_OR_EQUAL,
		token.MORE_THAN,
		token.MORE_THAN_OR_EQUAL:
		return Relation{left, op, right}

	case token.EQUAL, token.NOT_EQUAL:
		return Equality{left, op, right}

	case token.AND, token.OR:
		return Logic{left, op, right}
	}

	panic(string("[NewMathExpression] Unknown operator '" + string(op.Type) + "'"))
}
