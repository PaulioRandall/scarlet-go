package processor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

func numValue(n string) value.Num {
	return value.Num{Number: number.New(n)}
}

func TestProcess_Assign_1(t *testing.T) {

	// x := 1
	env := &runtimeEnv{
		ins: []inst.Inst{
			inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
			inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{
		value.Ident("x"): numValue("1"),
	}

	expStk := value.Stack{}

	p := New(env, env)
	p.Run()

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.True(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.False(t, p.Halted)
	require.Equal(t, expIds, env.ids)
	require.Equal(t, expStk, env.Stack)
}

func TestProcess_Assign_2(t *testing.T) {

	// x := y
	env := &runtimeEnv{
		ins: []inst.Inst{
			inst.Inst{Code: inst.FETCH_PUSH, Data: value.Ident("y")},
			inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
		},
		ids: map[value.Ident]value.Value{
			value.Ident("y"): numValue("1"),
		},
	}

	expIds := map[value.Ident]value.Value{
		value.Ident("y"): numValue("1"),
		value.Ident("x"): numValue("1"),
	}

	expStk := value.Stack{}

	p := New(env, env)
	p.Run()

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.True(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.False(t, p.Halted)
	require.Equal(t, expIds, env.ids)
	require.Equal(t, expStk, env.Stack)
}

func TestProcess_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, text
	env := &runtimeEnv{
		ins: []inst.Inst{
			inst.Inst{Code: inst.STACK_PUSH, Data: value.Bool(true)},
			inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("x")},
			inst.Inst{Code: inst.STACK_PUSH, Data: numValue("1")},
			inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("y")},
			inst.Inst{Code: inst.STACK_PUSH, Data: value.Str("text")},
			inst.Inst{Code: inst.SCOPE_BIND, Data: value.Ident("z")},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{
		value.Ident("x"): value.Bool(true),
		value.Ident("y"): numValue("1"),
		value.Ident("z"): value.Str("text"),
	}

	expStk := value.Stack{}

	p := New(env, env)
	p.Run()

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.True(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.False(t, p.Halted)
	require.Equal(t, expIds, env.ids)
	require.Equal(t, expStk, env.Stack)
}

func processBinOpTest(t *testing.T,
	exp, left, right value.Value,
	opCode inst.Code) {

	env := &runtimeEnv{
		ins: []inst.Inst{
			inst.Inst{Code: inst.STACK_PUSH, Data: left},
			inst.Inst{Code: inst.STACK_PUSH, Data: right},
			inst.Inst{Code: opCode},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{}

	expStk := value.Stack{}
	expStk.Push(exp)

	p := New(env, env)
	p.Run()

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.True(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.False(t, p.Halted)
	require.Equal(t, expIds, env.ids)

	// Implementations of number.Number may not return the correct results when
	// using == or != so number.Equal should be used to check equality.
	require.Equal(t, expStk.Size(), env.Stack.Size())
	want := expStk.Top()
	have := env.Stack.Top()
	if !have.Equal(want) {
		require.Equal(t, want, have)
	}
}

func TestProcess_Add(t *testing.T) {
	// 1 + 2
	processBinOpTest(t,
		numValue("3"),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_ADD,
	)
}

func TestProcess_Sub(t *testing.T) {
	// 1 - 2
	processBinOpTest(t,
		numValue("-1"),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_SUB,
	)
}

func TestProcess_Mul(t *testing.T) {
	// 2 * 4
	processBinOpTest(t,
		numValue("8"),
		numValue("2"),
		numValue("4"),
		inst.BIN_OP_MUL,
	)
}

func TestProcess_Div(t *testing.T) {
	// 12 / 3
	processBinOpTest(t,
		numValue("4"),
		numValue("12"),
		numValue("3"),
		inst.BIN_OP_DIV,
	)
}

func TestProcess_Rem(t *testing.T) {
	// 5 % 3
	processBinOpTest(t,
		numValue("2"),
		numValue("5"),
		numValue("3"),
		inst.BIN_OP_REM,
	)
}

func TestProcess_And(t *testing.T) {
	// false && false
	processBinOpTest(t,
		value.Bool(false),
		value.Bool(false),
		value.Bool(false),
		inst.BIN_OP_AND,
	)
	// true && false
	processBinOpTest(t,
		value.Bool(false),
		value.Bool(true),
		value.Bool(false),
		inst.BIN_OP_AND,
	)
	// true && true
	processBinOpTest(t,
		value.Bool(true),
		value.Bool(true),
		value.Bool(true),
		inst.BIN_OP_AND,
	)
}

func TestProcess_Or(t *testing.T) {
	// false || false
	processBinOpTest(t,
		value.Bool(false),
		value.Bool(false),
		value.Bool(false),
		inst.BIN_OP_OR,
	)
	// true || false
	processBinOpTest(t,
		value.Bool(true),
		value.Bool(true),
		value.Bool(false),
		inst.BIN_OP_OR,
	)
	// true || true
	processBinOpTest(t,
		value.Bool(true),
		value.Bool(true),
		value.Bool(true),
		inst.BIN_OP_OR,
	)
}

func TestProcess_Less(t *testing.T) {
	// 1 < 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_LESS,
	)
	// 2 < 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_LESS,
	)
	// 3 < 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("3"),
		numValue("2"),
		inst.BIN_OP_LESS,
	)
}

func TestProcess_More(t *testing.T) {
	// 1 > 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_MORE,
	)
	// 2 > 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_MORE,
	)
	// 3 > 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("3"),
		numValue("2"),
		inst.BIN_OP_MORE,
	)
}

func TestProcess_LessOrEqual(t *testing.T) {
	// 1 <= 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_LEQU,
	)
	// 2 <= 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_LEQU,
	)
	// 3 <= 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("3"),
		numValue("2"),
		inst.BIN_OP_LEQU,
	)
}

func TestProcess_MoreOrEqual(t *testing.T) {
	// 1 >= 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_MEQU,
	)
	// 2 >= 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_MEQU,
	)
	// 3 >= 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("3"),
		numValue("2"),
		inst.BIN_OP_MEQU,
	)
}

func TestProcess_Equal(t *testing.T) {
	// 1 == 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_EQU,
	)
	// 2 == 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_EQU,
	)
	// 2 == "apple"
	processBinOpTest(t,
		value.Bool(false),
		numValue("2"),
		value.Str("apple"),
		inst.BIN_OP_EQU,
	)
}

func TestProcess_NotEqual(t *testing.T) {
	// 1 != 2
	processBinOpTest(t,
		value.Bool(true),
		numValue("1"),
		numValue("2"),
		inst.BIN_OP_NEQU,
	)
	// 2 != 2
	processBinOpTest(t,
		value.Bool(false),
		numValue("2"),
		numValue("2"),
		inst.BIN_OP_NEQU,
	)
	// 2 != "apple"
	processBinOpTest(t,
		value.Bool(true),
		numValue("2"),
		value.Str("apple"),
		inst.BIN_OP_NEQU,
	)
}

func TestProcess_SpellCall_1(t *testing.T) {

	testSpell := func(env spell.Runtime, args []value.Value) []value.Value {
		require.Equal(t, 1, len(args))
		require.Equal(t, value.Str("abc"), args[0])
		return []value.Value{}
	}

	// x := y
	env := &runtimeEnv{
		ins: []inst.Inst{
			inst.Inst{Code: inst.STACK_PUSH, Data: value.Str("abc")},
			inst.Inst{Code: inst.SPELL_CALL, Data: value.Ident("Print")},
		},
		ids: map[value.Ident]value.Value{},
		book: spell.Book{
			"print": spell.Inscription{
				Spell:     testSpell,
				Name:      "Print",
				ParamsIn:  1,
				ParamsOut: spell.NO_ARGS,
			},
		},
	}

	expIds := map[value.Ident]value.Value{}

	expStk := value.Stack{}

	p := New(env, env)
	p.Run()

	require.Nil(t, env.err, "ERROR: %+v", env.err)
	require.True(t, env.exitFlag)
	require.Equal(t, 0, env.exitCode)
	require.False(t, p.Halted)
	require.Equal(t, expIds, env.ids)
	require.Equal(t, expStk, env.Stack)
}
