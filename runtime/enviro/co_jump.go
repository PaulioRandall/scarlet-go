package enviro

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coJumpBack(env *Environment, in inst.Instruction) {
	jumpSize := in.Data.(int)
	jumpSize = -jumpSize
	env.JumpBy(jumpSize)
}

func coJumpIf(env *Environment, in inst.Instruction, jumpIf bool) {

	condition, ok := env.PopVal().(types.Bool)
	if !ok {
		env.Fail(perror.New("Expected bool for jump condition"))
		return
	}

	if jumpIf == bool(condition) {
		jumpSize := in.Data.(int)
		env.JumpBy(jumpSize)
	}
}
