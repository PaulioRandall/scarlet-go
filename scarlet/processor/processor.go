package processor

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
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
	case tree.Guard:
		Guard(env, v)
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
	for i, v := range in {
		if v == nil {
			idx := strconv.Itoa(i + 1)
			panic("SANITY CHECK! Invalid nil argument " + idx + " for '@" + n.Name + "'")
		}
	}

	out := spell.NewOutput(s.Outputs)

	if !env.GetExitFlag() {
		s.Spell(env, in, out)
	}

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
	case tree.UnaryExpr:
		return UnaryExpr(env, v)
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
		return value.Num(v.Val)
	case tree.StrLit:
		return value.Str(v.Val[1 : len(v.Val)-1])
	default:
		panic("SANITY CHECK! Unknown tree.Literal type")
	}
}

func UnaryExpr(env Runtime, n tree.UnaryExpr) value.Value {
	switch n.Op {
	case tree.OP_EXIST:
		if id, ok := n.Term.(tree.Ident); ok {
			return env.Exists(value.Ident(id.Val))
		}
		return value.Bool(Expression(env, n.Term) != nil)

	default:
		panic("SANITY CHECK! Unknown tree.binaryExpr type")
	}
}

func BinaryExpr(env Runtime, n tree.BinaryExpr) value.Value {

	l, r := Expression(env, n.Left), Expression(env, n.Right)

	switch n.Op {
	case tree.OP_ADD:
		return l.(value.Num) + r.(value.Num)

	case tree.OP_SUB:
		return l.(value.Num) - r.(value.Num)

	case tree.OP_MUL:
		return l.(value.Num) * r.(value.Num)

	case tree.OP_DIV:
		return l.(value.Num) / r.(value.Num)

	case tree.OP_REM:
		x, y := l.(value.Num), r.(value.Num)
		for x >= y {
			x -= y
		}
		return x

	case tree.OP_AND:
		return l.(value.Bool) && r.(value.Bool)

	case tree.OP_OR:
		return l.(value.Bool) || r.(value.Bool)

	case tree.OP_LT:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum < rNum)

	case tree.OP_MT:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum > rNum)

	case tree.OP_LTE:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum <= rNum)

	case tree.OP_MTE:
		lNum, rNum := l.(value.Num), r.(value.Num)
		return value.Bool(lNum >= rNum)

	case tree.OP_EQU:
		return value.Bool(l.Equal(r))

	case tree.OP_NEQ:
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
	for i, v := range in {
		if v == nil {
			idx := strconv.Itoa(i + 1)
			panic("SANITY CHECK! Invalid nil argument " + idx + " for '@" + n.Name + "'")
		}
	}

	out := spell.NewOutput(1)
	s.Spell(env, in, out)
	return out.Get(0)
}

func Guard(env Runtime, g tree.Guard) {

	cond, ok := Expression(env, g.Cond).(value.Bool)
	if !ok {
		panic("SANITY CHECK! Expected boolean result")
	}

	if !cond {
		return
	}

	for _, v := range g.Body.Stmts {
		s, ok := v.(tree.Stat)
		if !ok {
			panic("SANITY CHECK! Result of expression ignored")
		}
		Statement(env, s)
	}
}
