package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/code"
	"github.com/PaulioRandall/scarlet-go/token/inst"
	"github.com/PaulioRandall/scarlet-go/token/number"
	"github.com/PaulioRandall/scarlet-go/token/token"
	"github.com/PaulioRandall/scarlet-go/token/tree"
	"github.com/PaulioRandall/scarlet-go/token/value"

	"github.com/stretchr/testify/require"
)

func requireInsts(t *testing.T, exp, act []inst.Inst) {
	require.Equal(t, len(exp), len(act))
	for i, v := range act {
		require.Equal(t, exp[i], v)
	}
}

func numValue(n string) value.Num {
	return value.Num{Number: number.New(n)}
}

func TestCompile_SingleAssign(t *testing.T) {

	// x := 1
	in := tree.SingleAssign{
		Left:  tree.Ident{Val: "x"},
		Right: tree.NumLit{Val: number.New("1")},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, "text"
	in := tree.MultiAssign{
		Left: []tree.Assignee{
			tree.Ident{Val: "x"},
			tree.Ident{Val: "y"},
			tree.Ident{Val: "z"},
		},
		Right: []tree.Expr{
			tree.BoolLit{Val: true},
			tree.NumLit{Val: number.New("1")},
			tree.StrLit{Val: "text"},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: value.Bool(true)},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("y")},
		inst.Inst{Code: code.STACK_PUSH, Data: value.Str("text")},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("z")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_BinaryExpr_1(t *testing.T) {

	// x := 1 + 2
	in := tree.SingleAssign{
		Left: tree.Ident{Val: "x"},
		Right: tree.BinaryExpr{
			Left:  tree.NumLit{Val: number.New("1")},
			Op:    token.ADD,
			Right: tree.NumLit{Val: number.New("2")},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: code.BIN_OP_ADD},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_BinaryExpr_2(t *testing.T) {

	// x := 1 + 2 * 3
	in := tree.SingleAssign{
		Left: tree.Ident{Val: "x"},
		Right: tree.BinaryExpr{
			Left: tree.NumLit{Val: number.New("1")},
			Op:   token.ADD,
			Right: tree.BinaryExpr{
				Left:  tree.NumLit{Val: number.New("2")},
				Op:    token.MUL,
				Right: tree.NumLit{Val: number.New("3")},
			},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("3")},
		inst.Inst{Code: code.BIN_OP_MUL},
		inst.Inst{Code: code.BIN_OP_ADD},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_BinaryExpr_3(t *testing.T) {

	// x, y := true && false, 1 + 2 * 3
	in := tree.MultiAssign{
		Left: []tree.Assignee{
			tree.Ident{Val: "x"},
			tree.Ident{Val: "y"},
		},
		Right: []tree.Expr{
			tree.BinaryExpr{
				Left:  tree.BoolLit{Val: true},
				Op:    token.AND,
				Right: tree.BoolLit{Val: false},
			},
			tree.BinaryExpr{
				Left: tree.NumLit{Val: number.New("1")},
				Op:   token.ADD,
				Right: tree.BinaryExpr{
					Left:  tree.NumLit{Val: number.New("2")},
					Op:    token.MUL,
					Right: tree.NumLit{Val: number.New("3")},
				},
			},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: value.Bool(true)},
		inst.Inst{Code: code.STACK_PUSH, Data: value.Bool(false)},
		inst.Inst{Code: code.BIN_OP_AND},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: code.STACK_PUSH, Data: numValue("3")},
		inst.Inst{Code: code.BIN_OP_MUL},
		inst.Inst{Code: code.BIN_OP_ADD},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("y")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}
