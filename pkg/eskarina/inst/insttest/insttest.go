package insttest

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/code"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"

	"github.com/stretchr/testify/require"
)

func Feign(ins ...*inst.Instruction) *inst.Instruction {

	var first *inst.Instruction
	var last *inst.Instruction

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

func newIn(code code.Code, data interface{}) *inst.Instruction {
	return &inst.Instruction{
		Code: code,
		Data: data,
	}
}

func Equal(t *testing.T, exp, act *inst.Instruction) {

	for exp != nil || act != nil {

		if exp == nil && act != nil {
			require.Nil(t, act, "Want: EOF\nHave: %s", act.String())
		}

		if exp != nil && act == nil {
			require.NotNil(t, act, "Want: %s\nHave: nil", exp.String())
		}

		equalContent(t, exp, act, fmt.Sprintf(
			"Unexepected instruction\nWant: %s\nHave: %s",
			exp.String(), act.String(),
		))

		exp, act = exp.Next, act.Next
	}
}

func equalContent(t *testing.T, exp, act *inst.Instruction, msg string) {

	if exp == nil {
		require.Nil(t, act, msg)
		return
	}

	require.NotNil(t, act, msg)
	require.Equal(t, exp.Code, act.Code, msg)
	require.Equal(t, exp.Data, act.Data, msg)
}
