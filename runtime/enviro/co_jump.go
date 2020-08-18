package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coJump(env *Environment, in inst.Instruction, jumpIf bool) {

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
