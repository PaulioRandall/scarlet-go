package processor

import (
	"errors"
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

type testRuntime struct {
	value.Stack
	started bool
	counter int
	ins     []inst.Inst
	ids     map[value.Ident]value.Value
}

func (rt *testRuntime) Next() (inst.Inst, bool) {

	if rt.started {
		rt.counter++
	} else {
		rt.started = true
	}

	if rt.counter >= len(rt.ins) {
		return inst.Inst{}, false
	}

	return rt.ins[rt.counter], true
}

func (rt *testRuntime) Fetch(id value.Ident) (value.Value, error) {
	if v, ok := rt.ids[id]; ok {
		return v, nil
	}
	return nil, errors.New("Identifier " + string(id) + " not found in scope")
}

func (rt *testRuntime) FetchPush(id value.Ident) error {
	if v, ok := rt.ids[id]; ok {
		rt.Push(v)
		return nil
	}
	return errors.New("Identifier " + string(id) + " not found in scope")
}

func (rt *testRuntime) Bind(id value.Ident, v value.Value) error {
	rt.ids[id] = v
	return nil
}

func numValue(n string) value.Num {
	return value.Num{Number: number.New(n)}
}

func TestProcess_Assign_1(t *testing.T) {

	// x := 1
	rt := &testRuntime{
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

	p := New(rt)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, rt.ids)
	require.Equal(t, expStk, rt.Stack)
}

func TestProcess_Assign_2(t *testing.T) {

	// x := y
	rt := &testRuntime{
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

	p := New(rt)
	p.Run()

	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.False(t, p.Stopped)
	require.Equal(t, expIds, rt.ids)
	require.Equal(t, expStk, rt.Stack)
}

func TestProcess_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, text
	rt := &testRuntime{
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

	p := New(rt)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, rt.ids)
	require.Equal(t, expStk, rt.Stack)
}

func processBinOpTest(t *testing.T,
	exp, left, right value.Value,
	opCode inst.Code) {

	rt := &testRuntime{
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

	p := New(rt)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, rt.ids)

	// Implementations of number.Number may not return the correct results when
	// using == or != so number.Equal should be used to check equality.
	require.Equal(t, expStk.Size(), rt.Stack.Size())
	want := expStk.Top()
	have := rt.Stack.Top()
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
