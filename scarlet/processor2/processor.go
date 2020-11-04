package processor2

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// Runtime is a handler for performing memory related and context dependent
// instructions such as access to the value stack and acces to scope variables.
type Runtime interface {

	// Spellbook returns the book containing the spells available during runtime.
	Spellbook() spell.Book

	// Bind sets the value of a variable overwriting any existing value.
	Bind(value.Ident, value.Value)

	// Fetch returns the value associated with the specified identifier.
	Fetch(value.Ident) value.Value

	// Fail sets the error and exit status a non-recoverable error occurs
	// during execution.
	Fail(int, error)

	// Exit causes the program to exit with the specified exit code.
	Exit(int)

	// GetErr returns the error if set else returns nil.
	GetErr() error

	// GetExitFlag returns true if the program should stop execution after
	// finishing any instruction currently being executed.
	GetExitFlag() bool
}

const (
	GENERAL_ERROR int = 1
)

func Execute(env Runtime, s tree.Stat) {
	switch v := s.(type) {
	case tree.SingleAssign:
		SingleAssign(env, v)
	case tree.MultiAssign:
		MultiAssign(env, v)
	case tree.SpellCall:
		SpellCall(env, v)
	default:
		panic("SANITY CHECK! Unknown tree.Stat type")
	}
}

func SingleAssign(env Runtime, n tree.SingleAssign) {
	l := Assignee(env, n.Left)
	r := Expression(env, n.Right)
	env.Bind(l, r)
}

func MultiAssign(env Runtime, n tree.MultiAssign) {

	var vals []value.Value

	if n.Asym {
		vals = MultiReturn(env, n.Right[0])
	} else {
		vals = Expressions(env, n.Right)
	}

	for i, v := range n.Left {
		l := Assignee(env, v)
		r := vals[i]
		env.Bind(l, r)
	}
}

func SpellCall(env Runtime, n tree.SpellCall) {
	// TODO
}

func Assignee(env Runtime, n tree.Assignee) value.Ident {
	switch v := n.(type) {
	case tree.Ident:
		return value.Ident(v.Val)
	default:
		panic("SANITY CHECK! Unknown tree.Assignee type")
	}
}

func MultiReturn(env Runtime, n tree.Expr) []value.Value {
	// TODO
	return nil
}

func Expressions(env Runtime, n []tree.Expr) []value.Value {
	r := make([]value.Value, len(n))
	for i, v := range n {
		r[i] = Expression(env, v)
	}
	return r
}

func Expression(env Runtime, n tree.Expr) value.Value {
	switch v := n.(type) {
	case tree.Ident:
		panic("Not implemented yet!")
	case tree.Literal:
		return Literal(env, v)
	case tree.BinaryExpr:
		return BinaryExpr(env, v)
	case tree.SpellCall:
		return SpellCallExpr(env, v)
	default:
		panic("SANITY CHECK! Unknown tree.Expr type")
	}
}

func Literal(env Runtime, n tree.Literal) value.Value {
	switch v := n.(type) {
	case tree.BoolLit:
		return value.Bool(v.Val)
	case tree.NumLit:
		return value.Num{Number: v.Val}
	case tree.StrLit:
		return value.Str(v.Val[1 : len(v.Val)-1])
	default:
		panic("SANITY CHECK! Unknown tree.Literal type")
	}
}

func BinaryExpr(env Runtime, n tree.BinaryExpr) value.Value {

	l, r := Expression(env, n.Left), Expression(env, n.Right)

	switch n.Op {
	case token.ADD:
		lNum, rNum := l.(value.Num), r.(value.Num)
		lNum.Number = lNum.Number.Copy()
		lNum.Number.Add(rNum.Number)
		return lNum

	case token.SUB:
		lNum, rNum := l.(value.Num), r.(value.Num)
		lNum.Number = lNum.Number.Copy()
		lNum.Number.Sub(rNum.Number)
		return lNum

	case token.MUL:
		lNum, rNum := l.(value.Num), r.(value.Num)
		lNum.Number = lNum.Number.Copy()
		lNum.Number.Mul(rNum.Number)
		return lNum

	case token.DIV:
		lNum, rNum := l.(value.Num), r.(value.Num)
		lNum.Number = lNum.Number.Copy()
		lNum.Number.Div(rNum.Number)
		return lNum

	case token.REM:
		lNum, rNum := l.(value.Num), r.(value.Num)
		lNum.Number = lNum.Number.Copy()
		lNum.Number.Mod(rNum.Number)
		return lNum

	case token.AND:
		return l.(value.Bool) && r.(value.Bool)

	case token.OR:
		return l.(value.Bool) || r.(value.Bool)

	case token.LESS:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.Less(rNum.Number))

	case token.MORE:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.More(rNum.Number))

	case token.LESS_EQUAL:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.LessOrEqual(rNum.Number))

	case token.MORE_EQUAL:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.MoreOrEqual(rNum.Number))

	case token.EQUAL:
		return value.Bool(l.Equal(r))

	case token.NOT_EQUAL:
		return value.Bool(!l.Equal(r))

	default:
		panic("SANITY CHECK! Unknown tree.binaryExpr type")
	}
}

func SpellCallExpr(env Runtime, n tree.SpellCall) value.Value {
	// TODO
	return nil
}
