package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coSpell(env *Environment, in inst.Instruction) {

	name := in.Data.(string)

	sp := spells.LookUp(name)
	if sp == nil {
		env.Fail(perror.New("Unknown spell %q", name))
		return
	}

	args := popArgs(env)
	sp.Invoke(env, args)
}

func popArgs(env *Environment) []types.Value {

	isNotDelim := func(v types.Value) bool {
		_, is := v.(types.Delim)
		return !is
	}

	vs := []types.Value{}

	for v := env.PopVal(); isNotDelim(v); v = env.PopVal() {
		vs = append([]types.Value{v}, vs...)
	}

	return vs
}
