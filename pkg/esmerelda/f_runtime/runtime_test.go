package runtime

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/f_runtime/enviro"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T,
	given []inst.Instruction,
	expStack []result.Result,
	expDefs map[string]result.Result,
	expBindings map[string]result.Result) {

	run := New(given)
	run.Start()

	requireDoneEnv(t, run.Env(), len(given)+1, expStack, expDefs, expBindings)
}

func requireStack(t *testing.T, exps, acts []result.Result) {

	expSize := len(exps)
	actSize := len(acts)

	for i := 0; i < expSize || i < actSize; i++ {

		require.True(t, i < actSize,
			"Expected: "+exps[i].String()+"\nBut no actual instructions remain")

		require.True(t, i < expSize,
			"Did not expect any more instructions\nBut got ("+acts[i].String()+")")

		require.Equal(t, exps[i], acts[i])
	}
}

func requireDoneEnv(
	t *testing.T,
	env *enviro.Environment,
	counter int,
	expStack []result.Result,
	expDefs map[string]result.Result,
	expBindings map[string]result.Result) {

	require.Nil(t, env.Err)
	require.True(t, env.Halted)
	require.True(t, env.Done)
	require.Equal(t, counter, env.Ctx.Counter)
	requireStack(t, expStack, env.Ctx.Stack.ToArray())
	require.Equal(t, expDefs, env.Defs)
	require.Equal(t, expBindings, *env.Ctx.Bindings)
}

func ins(code Code, data interface{}) inst.Instruction {
	return inst.Inst{
		InstCode: code,
		InstData: data,
		Opener:   token.Tok{},
		Closer:   token.Tok{},
	}
}

func Test1_1(t *testing.T) {

	// GIVEN a single value stack push instruction
	// THEN a single value is pushed to stack

	// "abc"
	given := []inst.Instruction{
		ins(IN_VAL_PUSH, "abc"),
	}

	expStack := []result.Result{
		result.Result{
			RType: result.RT_STRING,
			Value: "abc",
		},
	}

	expDefs := map[string]result.Result{}

	expBindings := map[string]result.Result{}

	doTest(t, given, expStack, expDefs, expBindings)
}

func Test1_2(t *testing.T) {

	// GIVEN a multiple value stack push instructions
	// THEN all values are pushed to stack

	// true, 1, "abc"
	given := []inst.Instruction{
		ins(IN_VAL_PUSH, true),
		ins(IN_VAL_PUSH, number.New("1")),
		ins(IN_VAL_PUSH, "abc"),
	}

	expStack := []result.Result{
		result.Result{
			RType: result.RT_STRING,
			Value: "abc",
		},
		result.Result{
			RType: result.RT_NUMBER,
			Value: number.New("1"),
		},
		result.Result{
			RType: result.RT_BOOL,
			Value: true,
		},
	}

	expDefs := map[string]result.Result{}

	expBindings := map[string]result.Result{}

	doTest(t, given, expStack, expDefs, expBindings)
}

func Test1_3(t *testing.T) {

	// GIVEN a set spell with parameters
	// THEN the spell is invoked with correct parameters
	// AND the expected variable binding is made

	// @Set("x", "abc")
	given := []inst.Instruction{
		ins(IN_VAL_PUSH, "x"),
		ins(IN_VAL_PUSH, "abc"),
		ins(IN_SPELL, []interface{}{2, "set"}),
	}

	expStack := []result.Result{}

	expDefs := map[string]result.Result{}

	expBindings := map[string]result.Result{
		"x": result.Result{
			RType: result.RT_STRING,
			Value: "abc",
		},
	}

	doTest(t, given, expStack, expDefs, expBindings)
}

// GIVEN a several spells
// THEN each spell is invoked
// AND the expected variable bindings are made

// @Set("x", "abc")
func Test1_4(t *testing.T) {

	given := []inst.Instruction{
		ins(IN_VAL_PUSH, "x"),
		ins(IN_VAL_PUSH, number.New("1")),
		ins(IN_SPELL, []interface{}{2, "set"}),
		ins(IN_VAL_PUSH, "y"),
		ins(IN_VAL_PUSH, number.New("2")),
		ins(IN_SPELL, []interface{}{2, "set"}),
		ins(IN_VAL_PUSH, "x"),
		ins(IN_VAL_PUSH, "y"),
		ins(IN_SPELL, []interface{}{2, "println"}),
	}

	expStack := []result.Result{}

	expDefs := map[string]result.Result{}

	expBindings := map[string]result.Result{
		"x": result.Result{
			RType: result.RT_NUMBER,
			Value: number.New("1"),
		},
		"y": result.Result{
			RType: result.RT_NUMBER,
			Value: number.New("2"),
		},
	}

	doTest(t, given, expStack, expDefs, expBindings)
}
