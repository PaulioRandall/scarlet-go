package processor

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// Runtime is a handler for performing memory related and context dependent
// instructions such as access to the value stack and acces to scope variables.
type Runtime interface {
	spell.Runtime
}

const (
	GENERAL_ERROR int = 1
)

func Statement(env Runtime, s tree.Stat) {
	switch v := s.(type) {
	case tree.SingleAssign:
		SingleAssign(env, v)
	case tree.MultiAssign:
		MultiAssign(env, v)
	case tree.AsymAssign:
		AsymAssign(env, v)
	case tree.SpellCall:
		SpellCall(env, v)
	default:
		panic("SANITY CHECK! Unknown tree.Stat type")
	}
}

func SingleAssign(env Runtime, n tree.SingleAssign) {
	assign(env, n.Left, Expression(env, n.Right))
}

func MultiAssign(env Runtime, n tree.MultiAssign) {
	vals := Expressions(env, n.Right)
	for i, a := range n.Left {
		assign(env, a, vals[i])
	}
}

func AsymAssign(env Runtime, n tree.AsymAssign) {
	vals := MultiReturn(env, n.Right)
	for i, a := range n.Left {
		assign(env, a, vals[i])
	}
}

func assign(env Runtime, a tree.Assignee, v value.Value) {
	switch id, ok := Assignee(env, a); {
	case !ok:
		return
	case v == nil:
		env.Unbind(id)
	default:
		env.Bind(id, v)
	}
}

func SpellCall(env Runtime, n tree.SpellCall) []value.Value {

	s, ok := env.Spellbook().Lookup(n.Name)
	if !ok {
		panic("SANITY CHECK! Unknown spell '" + n.Name + "'")
	}

	in := Expressions(env, n.Args)
	out := spell.NewOutput(s.Outputs)
	s.Spell(env, in, out)
	return out.Slice()
}

func Assignee(env Runtime, n tree.Assignee) (value.Ident, bool) {
	switch v := n.(type) {
	case tree.Ident:
		return value.Ident(v.Val), true
	case tree.AnonIdent:
		return value.Ident(""), false
	default:
		panic("SANITY CHECK! Unknown tree.Assignee type")
	}
}

func MultiReturn(env Runtime, n tree.Expr) []value.Value {
	switch v := n.(type) {
	case tree.SpellCall:
		return SpellCall(env, v)
	}
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
		return Ident(env, v)
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

func Ident(env Runtime, n tree.Ident) value.Value {
	return env.Fetch(value.Ident(n.Val))
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

	case token.LT:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.Less(rNum.Number))

	case token.MT:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.More(rNum.Number))

	case token.LTE:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.LessOrEqual(rNum.Number))

	case token.MTE:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum.Number.MoreOrEqual(rNum.Number))

	case token.EQU:
		return value.Bool(l.Equal(r))

	case token.NEQ:
		return value.Bool(!l.Equal(r))

	default:
		panic("SANITY CHECK! Unknown tree.binaryExpr type")
	}
}

func SpellCallExpr(env Runtime, n tree.SpellCall) value.Value {

	s, ok := env.Spellbook().Lookup(n.Name)
	if !ok {
		panic("SANITY CHECK! Unknown spell '@" + n.Name + "'")
	}
	if s.Outputs != 1 {
		panic("SANITY CHECK! '@" + n.Name + "' is not a single result expression")
	}

	in := Expressions(env, n.Args)
	out := spell.NewOutput(1)
	s.Spell(env, in, out)
	return out.Get(0)
}
