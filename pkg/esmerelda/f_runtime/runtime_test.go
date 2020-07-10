package runtime

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"

	"github.com/stretchr/testify/require"
)

func ins(code Code, data interface{}) inst.Instruction {
	return inst.Inst{
		InstCode: code,
		InstData: data,
	}
}

func Test1_1(t *testing.T) {

	data := true

	given := []inst.Instruction{
		ins(IN_VAL_PUSH, data),
	}

	run := New(given)
	finished, e := run.Start()

	require.True(t, finished)
	require.Nil(t, e)
}
