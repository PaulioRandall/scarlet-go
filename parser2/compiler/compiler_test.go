package compiler

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/number"

	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/tree"
	"github.com/PaulioRandall/scarlet-go/token2/value"

	"github.com/stretchr/testify/require"
)

func requireRiscs(t *testing.T, exp, act []inst.RiscInst) {
	require.Equal(t, len(exp), len(act))
	for i, v := range act {
		require.Equal(t, exp[i], v)
	}
}

func requireData(t *testing.T, exp, act inst.InstData) {
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

	expRisc := []inst.RiscInst{
		inst.RiscInst{Inst: inst.STK_PUSH, Data: 1},
		inst.RiscInst{Inst: inst.SCP_BIND, Data: 2},
	}

	expData := inst.InstData{
		1: value.Num{number.New("1")},
		2: value.Ident("x"),
	}

	ds := inst.NewDataSet()
	actRisc := Compile(in, ds)
	requireRiscs(t, expRisc, actRisc)
	requireData(t, expData, ds.Compile())
}
