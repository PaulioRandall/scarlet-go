package enviro

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coSpell(env *Environment, in inst.Instruction) {

	name := in.Data.(string)

	entry, ok := env.Spells.LookUp(name)
	if !ok {
		e := newErr("Unknown spell '%q'", name)
		env.Fail(e)
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
