package enviro

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/spells"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coSpell(env *Environment, in inst.Instruction) {

	name := in.Data.(string)

	entry, ok := spells.LookUp(name)
	if !ok {
		env.Fail(perror.New("Unknown spell %q", name))
		return
	}

	args := popArgs(env)
	entry.Spell(entry, env, args)
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
