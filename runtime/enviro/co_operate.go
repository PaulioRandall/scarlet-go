package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func popOperands(env *Environment) (left, right types.Num, ok bool) {

	right, ok = env.Pop().(types.Num)
	if !ok {
		env.Fail(perror.New("Expected number on right side of operation"))
		return
	}

	left, ok = env.Pop().(types.Num)
	if !ok {
		env.Fail(perror.New("Expected number on left side of operation"))
	}

	return
}

func coAdd(env *Environment, in inst.Instruction) {

	left, right, ok := popOperands(env)
	if !ok {
		return
	}

	left.Add(right.Number)
	env.Push(left)
}
