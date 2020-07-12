package runtime

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/f_runtime/enviro"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"

	"github.com/stretchr/testify/require"
)

func ins(code Code, data interface{}) inst.Instruction {
	return inst.Inst{
		InstCode: code,
		InstData: data,
	}
}

func requireStack(t *testing.T, stk *enviro.Stack, exps ...result.Result) {

	i := 0

	stk.Descend(func(act result.Result) {
		require.True(t, i < len(exps), "Environment stack contains too many bindings")
		require.Equal(t, exps[i], act)
		i++
	})
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

	require.True(t, run.Env().Halted)
	require.True(t, run.Env().Done)
	require.Equal(t, 0, len(run.Env().Defs))
	require.Nil(t, run.Env().Err)

	require.Equal(t, 2, run.Env().Ctx.Counter)
	require.Equal(t, 0, len(*run.Env().Ctx.Bindings))

	requireStack(t, run.Env().Ctx.Stack,
		result.Result{
			RType: result.RT_BOOL,
			Value: data,
		},
	)
}

func Test1_2(t *testing.T) {

	given := []inst.Instruction{
		ins(IN_VAL_PUSH, true),
		ins(IN_VAL_PUSH, "abc"),
	}

	run := New(given)
	finished, e := run.Start()

	require.True(t, finished)
	require.Nil(t, e)

	require.True(t, run.Env().Halted)
	require.True(t, run.Env().Done)
	require.Equal(t, 0, len(run.Env().Defs))
	require.Nil(t, run.Env().Err)

	require.Equal(t, 3, run.Env().Ctx.Counter)
	require.Equal(t, 0, len(*run.Env().Ctx.Bindings))

	requireStack(t, run.Env().Ctx.Stack,
		result.Result{
			RType: result.RT_STRING,
			Value: "abc",
		},
		result.Result{
			RType: result.RT_BOOL,
			Value: true,
		},
	)
}
