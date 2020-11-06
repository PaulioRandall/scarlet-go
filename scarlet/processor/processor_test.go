package processor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

func numValue(n string) value.Num { return value.Num{Number: number.New(n)} }
func ident(id string) tree.Ident  { return tree.Ident{Val: id} }
func anon() tree.AnonIdent        { return tree.AnonIdent{} }
func numLit(n string) tree.NumLit { return tree.NumLit{Val: number.New(n)} }
func boolLit(b bool) tree.BoolLit { return tree.BoolLit{Val: b} }
func strLit(s string) tree.StrLit { return tree.StrLit{Val: s} }
func binExpr(l tree.Expr, op token.Token, r tree.Expr) tree.BinaryExpr {
	return tree.BinaryExpr{Left: l, Op: op, Right: r}
}

type expRuntime struct {
	exitCode int
	exitFlag bool
	err      error
}

func assertRuntime(t *testing.T, exp *expRuntime, act *testRuntime) {
	require.Equal(t, exp.err, act.err)
	require.Equal(t, exp.exitFlag, act.exitFlag)
	require.Equal(t, exp.exitCode, act.exitCode)
}

func assertOutput(t *testing.T, exp interface{}, act interface{}) {
	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.IsType(t, exp, act)

	if want, ok := exp.(value.Num); ok {
		have := act.(value.Num)
		require.True(t, want.Equal(have))
		return
	}

	require.Equal(t, exp, act)
}

func TestLiteral(t *testing.T) {

	var assertions = []struct {
		in  tree.Literal
		exp value.Value
	}{
		{ // 0
			in:  boolLit(true),
			exp: value.Bool(true),
		}, { // 1
			in:  boolLit(false),
			exp: value.Bool(false),
		}, { // 2
			in:  numLit("1"),
			exp: numValue("1"),
		}, { // 3
			in:  strLit(`"abc"`),
			exp: value.Str(`abc`),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertOutput(t, a.exp, act)
	}
}

func TestArithBinExpr(t *testing.T) {

	var assertions = []struct {
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(numLit("1"), token.ADD, numLit("2")),
			exp: numValue("3"),
		}, { // 1
			in:  binExpr(numLit("4"), token.SUB, numLit("1")),
			exp: numValue("3"),
		}, { // 2
			in:  binExpr(numLit("3"), token.MUL, numLit("4")),
			exp: numValue("12"),
		}, { // 3
			in:  binExpr(numLit("12"), token.DIV, numLit("4")),
			exp: numValue("3"),
		}, { // 4
			in:  binExpr(numLit("5"), token.REM, numLit("3")),
			exp: numValue("2"),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertOutput(t, a.exp, act)
	}
}

func TestLogicBinExpr(t *testing.T) {

	var assertions = []struct {
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(boolLit(true), token.AND, boolLit(true)),
			exp: value.Bool(true),
		}, { // 1
			in:  binExpr(boolLit(true), token.AND, boolLit(false)),
			exp: value.Bool(false),
		}, { // 2
			in:  binExpr(boolLit(false), token.AND, boolLit(false)),
			exp: value.Bool(false),
		}, { // 3
			in:  binExpr(boolLit(true), token.OR, boolLit(true)),
			exp: value.Bool(true),
		}, { // 4
			in:  binExpr(boolLit(true), token.OR, boolLit(false)),
			exp: value.Bool(true),
		}, { // 5
			in:  binExpr(boolLit(false), token.OR, boolLit(false)),
			exp: value.Bool(false),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertOutput(t, a.exp, act)
	}
}

func TestCompBinExpr(t *testing.T) {

	var assertions = []struct {
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(numLit("1"), token.LT, numLit("2")),
			exp: value.Bool(true),
		}, { // 1
			in:  binExpr(numLit("2"), token.LT, numLit("2")),
			exp: value.Bool(false),
		}, { // 2
			in:  binExpr(numLit("3"), token.LT, numLit("2")),
			exp: value.Bool(false),
		}, { // 3
			in:  binExpr(numLit("1"), token.MT, numLit("2")),
			exp: value.Bool(false),
		}, { // 4
			in:  binExpr(numLit("2"), token.MT, numLit("2")),
			exp: value.Bool(false),
		}, { // 5
			in:  binExpr(numLit("3"), token.MT, numLit("2")),
			exp: value.Bool(true),
		}, { // 6
			in:  binExpr(numLit("1"), token.LTE, numLit("2")),
			exp: value.Bool(true),
		}, { // 7
			in:  binExpr(numLit("2"), token.LTE, numLit("2")),
			exp: value.Bool(true),
		}, { // 8
			in:  binExpr(numLit("3"), token.LTE, numLit("2")),
			exp: value.Bool(false),
		}, { // 9
			in:  binExpr(numLit("1"), token.MTE, numLit("2")),
			exp: value.Bool(false),
		}, { // 10
			in:  binExpr(numLit("2"), token.MTE, numLit("2")),
			exp: value.Bool(true),
		}, { // 11
			in:  binExpr(numLit("3"), token.MTE, numLit("2")),
			exp: value.Bool(true),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertOutput(t, a.exp, act)
	}
}

func TestEqualBinExpr(t *testing.T) {

	var assertions = []struct {
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(numLit("1"), token.EQU, numLit("1")),
			exp: value.Bool(true),
		}, { // 1
			in:  binExpr(numLit("1"), token.EQU, numLit("2")),
			exp: value.Bool(false),
		}, { // 2
			in:  binExpr(numLit("1"), token.EQU, strLit("abc")),
			exp: value.Bool(false),
		}, { // 3
			in:  binExpr(numLit("1"), token.NEQ, numLit("1")),
			exp: value.Bool(false),
		}, { // 4
			in:  binExpr(numLit("1"), token.NEQ, numLit("2")),
			exp: value.Bool(true),
		}, { // 5
			in:  binExpr(numLit("1"), token.NEQ, strLit("abc")),
			exp: value.Bool(true),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertOutput(t, a.exp, act)
	}
}

func TestExprs(t *testing.T) {

	var assertions = []struct {
		in  []tree.Expr
		exp []value.Value
	}{
		{ // 0
			in: []tree.Expr{
				numLit("1"),
				binExpr(numLit("1"), token.ADD, numLit("2")),
				binExpr(numLit("1"), token.EQU, strLit("abc")),
			},
			exp: []value.Value{
				numValue("1"),
				numValue("3"),
				value.Bool(false),
			},
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expressions(env, a.in)
		require.Equal(t, len(a.exp), len(act))
		for i := 0; i < len(a.exp); i++ {
			assertOutput(t, a.exp[i], act[i])
		}
	}
}

func TestSingleAssign_1(t *testing.T) {

	// x := 1
	in := tree.SingleAssign{
		Left:  ident("x"),
		Right: numLit("1"),
	}

	exp := newTestEnv()
	exp.scope[value.Ident("x")] = numValue("1")

	act := newTestEnv()
	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestSingleAssign_2(t *testing.T) {

	// _ := 1
	in := tree.SingleAssign{
		Left:  anon(),
		Right: numLit("1"),
	}

	exp := newTestEnv()

	act := newTestEnv()
	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestMultiAssign_1(t *testing.T) {

	// x, y, z := true, 1, "abc"
	in := tree.MultiAssign{
		Left:  []tree.Assignee{ident("x"), ident("y"), ident("z")},
		Right: []tree.Expr{boolLit(true), numLit("1"), strLit(`"abc"`)},
	}

	// x: true, y: 1, z: "abc"
	exp := newTestEnv()
	exp.scope[value.Ident("x")] = value.Bool(true)
	exp.scope[value.Ident("y")] = numValue("1")
	exp.scope[value.Ident("z")] = value.Str("abc")

	act := newTestEnv()
	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestMultiAssign_2(t *testing.T) {

	// _, y, _ := true, 1, "abc"
	in := tree.MultiAssign{
		Left:  []tree.Assignee{anon(), ident("y"), anon()},
		Right: []tree.Expr{boolLit(true), numLit("1"), strLit(`"abc"`)},
	}

	// y: 1
	exp := newTestEnv()
	exp.scope[value.Ident("y")] = numValue("1")

	act := newTestEnv()
	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestAsymAssign_1(t *testing.T) {

	// x, y, z := @Reverse(true, 123, "abc")
	in := tree.AsymAssign{
		Left: []tree.Assignee{ident("x"), ident("y"), ident("z")},
		Right: tree.SpellCall{
			Name: "reverse",
			Args: []tree.Expr{
				boolLit(true),
				numLit("123"),
				strLit(`"abc"`),
			},
		},
	}

	book := spell.Book{
		"reverse": spell.Inscription{
			Name:    "Reverse",
			Outputs: 3,
			Spell: func(env spell.Runtime, in []value.Value, out *spell.Output) {
				out.Set(0, value.Str("abc"))
				out.Set(1, numValue("1"))
				out.Set(2, value.Bool(true))
			},
		},
	}

	// x: "abc", y: 1, z: true
	exp := newTestEnv()
	exp.book = book
	exp.scope[value.Ident("x")] = value.Str("abc")
	exp.scope[value.Ident("y")] = numValue("1")
	exp.scope[value.Ident("z")] = value.Bool(true)

	act := newTestEnv()
	act.book = book

	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestAsymAssign_2(t *testing.T) {

	// _, y, _ := @Reverse(true, 123, "abc")
	in := tree.AsymAssign{
		Left: []tree.Assignee{anon(), ident("y"), anon()},
		Right: tree.SpellCall{
			Name: "reverse",
			Args: []tree.Expr{
				boolLit(true),
				numLit("123"),
				strLit(`"abc"`),
			},
		},
	}

	book := spell.Book{
		"reverse": spell.Inscription{
			Name:    "Reverse",
			Outputs: 3,
			Spell: func(env spell.Runtime, in []value.Value, out *spell.Output) {
				out.Set(0, value.Str("abc"))
				out.Set(1, numValue("1"))
				out.Set(2, value.Bool(true))
			},
		},
	}

	// y: 1
	exp := newTestEnv()
	exp.book = book
	exp.scope[value.Ident("y")] = numValue("1")

	act := newTestEnv()
	act.book = book

	Statement(act, in)
	require.Equal(t, exp, act)
}

func TestSpellCall_1(t *testing.T) {

	in := tree.SpellCall{
		Name: "Concat",
		Args: []tree.Expr{
			strLit(`"abc"`),
			strLit(`"123"`),
		},
	}

	expOut := []value.Value{value.Str("abc123")}

	book := spell.Book{
		"concat": spell.Inscription{
			Name:    "Concat",
			Outputs: 1,
			Spell: func(env spell.Runtime, in []value.Value, out *spell.Output) {
				require.Equal(t, 2, len(in))
				require.Equal(t, value.Str("abc"), in[0])
				require.Equal(t, value.Str("123"), in[1])
				require.Equal(t, 1, len(out.Slice()))
				out.Set(0, value.Str("abc123"))
			},
		},
	}

	exp := newTestEnv()
	exp.book = book

	act := newTestEnv()
	act.book = book

	out := SpellCall(act, in)
	require.Equal(t, exp, act)
	require.Equal(t, expOut, out)
}
