package processor2

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

func numValue(n string) value.Num { return value.Num{Number: number.New(n)} }
func ident(id string) tree.Ident  { return tree.Ident{Val: id} }
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

func assertRuntime(t *testing.T, exp expRuntime, act *testRuntime) {
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
		env expRuntime
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
		assertRuntime(t, a.env, env)
		assertOutput(t, a.exp, act)
	}
}

func TestArithExpr(t *testing.T) {

	var assertions = []struct {
		env expRuntime
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(numLit("1"), token.ADD, numLit("2")),
			exp: numValue("3"),
		}, { //1
			in:  binExpr(numLit("4"), token.SUB, numLit("1")),
			exp: numValue("3"),
		}, { //2
			in:  binExpr(numLit("3"), token.MUL, numLit("4")),
			exp: numValue("12"),
		}, { //3
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
		assertRuntime(t, a.env, env)
		assertOutput(t, a.exp, act)
	}
}

func TestCompExpr(t *testing.T) {

	var assertions = []struct {
		env expRuntime
		in  tree.BinaryExpr
		exp value.Value
	}{
		{ // 0
			in:  binExpr(numLit("1"), token.LESS, numLit("2")),
			exp: value.Bool(true),
		}, { //1
			in:  binExpr(numLit("2"), token.LESS, numLit("2")),
			exp: value.Bool(false),
		}, { //2
			in:  binExpr(numLit("3"), token.LESS, numLit("2")),
			exp: value.Bool(false),
		}, { // 3
			in:  binExpr(numLit("1"), token.MORE, numLit("2")),
			exp: value.Bool(false),
		}, { // 4
			in:  binExpr(numLit("2"), token.MORE, numLit("2")),
			exp: value.Bool(false),
		}, { // 5
			in:  binExpr(numLit("3"), token.MORE, numLit("2")),
			exp: value.Bool(true),
		}, { // 6
			in:  binExpr(numLit("1"), token.LESS_EQUAL, numLit("2")),
			exp: value.Bool(true),
		}, { // 7
			in:  binExpr(numLit("2"), token.LESS_EQUAL, numLit("2")),
			exp: value.Bool(true),
		}, { // 8
			in:  binExpr(numLit("3"), token.LESS_EQUAL, numLit("2")),
			exp: value.Bool(false),
		}, { // 9
			in:  binExpr(numLit("1"), token.MORE_EQUAL, numLit("2")),
			exp: value.Bool(false),
		}, { // 10
			in:  binExpr(numLit("2"), token.MORE_EQUAL, numLit("2")),
			exp: value.Bool(true),
		}, { // 11
			in:  binExpr(numLit("3"), token.MORE_EQUAL, numLit("2")),
			exp: value.Bool(true),
		},
	}

	for i, a := range assertions {
		t.Logf("Assertion %d", i)
		env := newTestEnv()
		act := Expression(env, a.in)
		assertRuntime(t, a.env, env)
		assertOutput(t, a.exp, act)
	}
}
