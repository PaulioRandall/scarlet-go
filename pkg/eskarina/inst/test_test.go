package inst

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/code"

	"github.com/stretchr/testify/require"
)

func newIn(code code.Code, data interface{}) *Instruction {
	return &Instruction{
		Code: code,
		Data: data,
	}
}

func feign(ins ...*Instruction) *Instruction {

	var first *Instruction
	var last *Instruction

	for _, in := range ins {

		if first == nil {
			first = in
			last = in
			continue
		}

		last.Next = in
		last = in
	}

	return first
}

func setup() (a, b, c, d *Instruction) {
	a = newIn(code.CO_VAL_PUSH, "a")
	b = newIn(code.CO_CTX_GET, "b")
	c = newIn(code.CO_SPELL, "c")
	d = newIn(code.CO_VAL_PUSH, "d")
	return
}

func setupList() (a, b, c, d *Instruction) {
	a, b, c, d = setup()
	_ = feign(a, b, c)
	return
}

func halfEqual(t *testing.T, exp, act *Instruction) {

	if exp == nil {
		require.Nil(t, act)
		return
	}

	require.NotNil(t, act)
	require.Equal(t, exp.Code, act.Code)
	require.Equal(t, exp.Data, act.Data)
}

func fullEqual(t *testing.T, exp, next, act *Instruction) {
	halfEqual(t, exp, act)
	halfEqual(t, next, act.Next)
}
