package insttest

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/shared/inst"

	"github.com/stretchr/testify/require"
)

func NewIn(code inst.Code, data interface{}) inst.Instruction {
	return inst.Instruction{
		Code: code,
		Data: data,
	}
}

func Equal(t *testing.T, exps, acts []inst.Instruction) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		if i >= expSize {
			require.Fail(t, "Want: EOF\nHave: %s", acts[i].String())
		}

		if i >= actSize {
			require.Fail(t, "Want: %s\nHave: nil", exps[i].String())
		}

		exp := exps[i]
		act := acts[i]

		msg := fmt.Sprintf(
			"Unexepected instruction\nWant: %s\nHave: %s",
			exp.String(), act.String(),
		)
		require.Equal(t, exp.Code, act.Code, msg)
		require.Equal(t, exp.Data, act.Data, msg)
	}
}
