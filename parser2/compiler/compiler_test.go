package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/number"

	"github.com/PaulioRandall/scarlet-go/token2/code"
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"

	"github.com/stretchr/testify/require"
)

func requireInsts(t *testing.T, exp, act []inst.Inst) {
	require.Equal(t, len(exp), len(act))
	for i, v := range act {
		require.Equal(t, exp[i], v)
	}
}

func TestCompile_SingleAssign(t *testing.T) {

	// x := 1
	in := tree.SingleAssign{
		Left:  tree.Ident{Val: "x"},
		Right: tree.NumLit{Val: number.New("1")},
	}

	exp := []inst.Inst{
		inst.Inst{Code: code.STACK_PUSH, Data: value.Num{number.New("1")}},
		inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
	}

	act := Compile(in)
	requireInsts(t, exp, act)
}
