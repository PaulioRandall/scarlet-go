package processor2

import (
	//"testing"

	//"github.com/PaulioRandall/scarlet-go/scarlet/token"
	//	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"
	//"github.com/stretchr/testify/require"
)

func numValue(n string) value.Num {
	return value.Num{Number: number.New(n)}
}

/*

func TestLiteral_1(t *testing.T) {

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
*/
