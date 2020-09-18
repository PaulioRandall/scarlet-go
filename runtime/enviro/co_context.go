package enviro

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func coCtxGet(env *Environment, in inst.Instruction) {

	id := in.Data.(string)
	r, ok := env.Get(id)

	if !ok {
		e := newErr("Undeclared variable '%q'", id)
		env.Fail(e)
		return
	}

	env.PushVal(r)
}

func coCtxSet(env *Environment, in inst.Instruction) {

	id := in.Data.(string)
	v := env.PopVal()

	if v == nil {
		e := newErr("Assignment fail '%q', value stack is empty", id)
		env.Fail(e)
		return
	}

	if _, ok := v.(types.Nil); ok {
		env.Unbind(id)
		return
	}

	env.Bind(id, v)
}
