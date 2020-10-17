package processor

import (
	"errors"
	"testing"

	"github.com/PaulioRandall/scarlet-go/number"

	"github.com/PaulioRandall/scarlet-go/token2/code"
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/value"

	"github.com/stretchr/testify/require"
)

type testMem struct {
	ins []inst.Inst
	ids map[value.Ident]value.Value
}

func (m *testMem) Has(c Counter) bool {
	return uint(len(m.ins)) > uint(c)
}

func (m *testMem) Fetch(c Counter) (inst.Inst, error) {
	if !m.Has(c) {
		return inst.Inst{}, errors.New("Program counter out of bounds")
	}
	return m.ins[c], nil
}

func (m *testMem) Get(id value.Ident) (value.Value, error) {
	if v, ok := m.ids[id]; ok {
		return v, nil
	}
	return nil, errors.New("Identifier " + string(id) + " not found in scope")
}

func (m *testMem) Bind(id value.Ident, v value.Value) error {
	m.ids[id] = v
	return nil
}

func TestProcess_Assign(t *testing.T) {

	// x := 1
	m := &testMem{
		ins: []inst.Inst{
			inst.Inst{Code: code.STACK_PUSH, Data: value.Num{number.New("1")}},
			inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{
		value.Ident("x"): value.Num{number.New("1")},
	}

	expStk := value.Stack{}

	p := New(m)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, m.ids)
	require.Equal(t, expStk, p.Stack)
}

func TestProcess_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, text
	m := &testMem{
		ins: []inst.Inst{
			inst.Inst{Code: code.STACK_PUSH, Data: value.Bool(true)},
			inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("x")},
			inst.Inst{Code: code.STACK_PUSH, Data: value.Num{number.New("1")}},
			inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("y")},
			inst.Inst{Code: code.STACK_PUSH, Data: value.Str("text")},
			inst.Inst{Code: code.SCOPE_BIND, Data: value.Ident("z")},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{
		value.Ident("x"): value.Bool(true),
		value.Ident("y"): value.Num{number.New("1")},
		value.Ident("z"): value.Str("text"),
	}

	expStk := value.Stack{}

	p := New(m)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, m.ids)
	require.Equal(t, expStk, p.Stack)
}

func TestProcess_Add(t *testing.T) {

	// 1 + 2
	m := &testMem{
		ins: []inst.Inst{
			inst.Inst{Code: code.STACK_PUSH, Data: value.Num{number.New("1")}},
			inst.Inst{Code: code.STACK_PUSH, Data: value.Num{number.New("2")}},
			inst.Inst{Code: code.OP_ADD},
		},
		ids: map[value.Ident]value.Value{},
	}

	expIds := map[value.Ident]value.Value{}

	expStk := value.Stack{}
	expStk.Push(value.Num{number.New("3")})

	p := New(m)
	p.Run()

	require.False(t, p.Stopped)
	require.Nil(t, p.Err, "ERROR: %+v", p.Err)
	require.Equal(t, expIds, m.ids)
	require.Equal(t, expStk, p.Stack)
}
