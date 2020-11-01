package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

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
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
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
		inst.Inst{Code: inst.STACK_PUSH, Data: value.Bool(true)},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("y")},
		inst.Inst{Code: inst.STACK_PUSH, Data: value.Str("text")},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("z")},
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
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: inst.BIN_OP_ADD},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
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
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("3")},
		inst.Inst{Code: inst.BIN_OP_MUL},
		inst.Inst{Code: inst.BIN_OP_ADD},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
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
		inst.Inst{Code: inst.STACK_PUSH, Data: value.Bool(true)},
		inst.Inst{Code: inst.STACK_PUSH, Data: value.Bool(false)},
		inst.Inst{Code: inst.BIN_OP_AND},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("2")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("3")},
		inst.Inst{Code: inst.BIN_OP_MUL},
		inst.Inst{Code: inst.BIN_OP_ADD},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("y")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_BinaryExpr_4(t *testing.T) {

	// x := y + z
	in := tree.SingleAssign{
		Left: tree.Ident{Val: "x"},
		Right: tree.BinaryExpr{
			Left:  tree.Ident{Val: "y"},
			Op:    token.ADD,
			Right: tree.Ident{Val: "z"},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: inst.FETCH_PUSH, Data: value.Ident("y")},
		inst.Inst{Code: inst.FETCH_PUSH, Data: value.Ident("z")},
		inst.Inst{Code: inst.BIN_OP_ADD},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_BinaryExpr_5(t *testing.T) {

	// x := y + @Add(1, 1)
	in := tree.SingleAssign{
		Left: tree.Ident{Val: "x"},
		Right: tree.BinaryExpr{
			Left: tree.Ident{Val: "y"},
			Op:   token.ADD,
			Right: tree.SpellCall{
				Name: "Add",
				Args: []tree.Expr{
					tree.NumLit{Val: number.New("1")},
					tree.NumLit{Val: number.New("1")},
				},
			},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: inst.FETCH_PUSH, Data: value.Ident("y")},
		inst.Inst{Code: inst.STACK_PUSH},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
		inst.Inst{Code: inst.SPELL_CALL, Data: value.Ident("Add")},
		inst.Inst{Code: inst.BIN_OP_ADD},
		inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_SpellCall_1(t *testing.T) {

	// @Print()
	in := tree.SpellCall{
		Name: "Print",
		Args: []tree.Expr{},
	}

	exp := []inst.Inst{
		inst.Inst{Code: inst.STACK_PUSH},
		inst.Inst{Code: inst.SPELL_CALL, Data: value.Ident("Print")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}

func TestCompile_SpellCall_2(t *testing.T) {

	// @Print(x)
	in := tree.SpellCall{
		Name: "Print",
		Args: []tree.Expr{
			tree.Ident{Val: "x"},
		},
	}

	exp := []inst.Inst{
		inst.Inst{Code: inst.STACK_PUSH},
		inst.Inst{Code: inst.FETCH_PUSH, Data: value.Ident("x")},
		inst.Inst{Code: inst.SPELL_CALL, Data: value.Ident("Print")},
	}

	act, e := Compile(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireInsts(t, exp, act)
}
