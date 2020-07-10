package testutils

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/number"

	"github.com/stretchr/testify/require"
)

func RequireInstructionSlice(t *testing.T, exps, acts []inst.Instruction) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected: "+exps[i].String()+"\nBut no actual instructions remain")

		require.True(t, i < expSize,
			"Did not expect any more instructions\nBut got ("+acts[i].String()+")")

		requireInstruction(t, exps[i], acts[i])
	}
}

func requireInstruction(t *testing.T, exp, act inst.Instruction) {

	require.NotNil(t, act, "Expected instruction ("+exp.String()+")\nBut got nil")
	msg := "Expected (" + exp.String() + ")\nActual   (" + act.String() + ")"

	require.Equal(t, exp.Code(), act.Code(), msg)
	requireData(t, exp.Data(), act.Data(), msg)
	requireSnippet(t, exp, act, msg)
}

func requireData(t *testing.T, exp, act interface{}, msg string) {

	if v, ok := exp.([]interface{}); ok {
		require.Implements(t, ([]interface{})(nil), act)
		requireDataSlice(t, v, act.([]interface{}), msg)
		return
	}

	requireDataItem(t, exp, act, msg)
}

func requireDataSlice(t *testing.T, exps, acts []interface{}, msg string) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected %q\nBut no actual data items remain", exps[i])

		require.True(t, i < expSize,
			"Did not expect any more data items\nBut got %q", acts[i])

		requireDataItem(t, exps[i], acts[i], msg)
	}
}

func requireDataItem(t *testing.T, exp, act interface{}, msg string) {

	switch v := exp.(type) {
	case number.Number:
		require.Implements(t, (*number.Number)(nil), act, msg)
		n := act.(number.Number)
		require.True(t, v.Equal(n), msg)

	default:
		require.Equal(t, exp, act, msg)
	}
}
