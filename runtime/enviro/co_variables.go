package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
)

func coCtxGet(env *Environment, in inst.Instruction) {

	id := in.Data.(string)
	r, ok := env.Get(id)

	if !ok {
		env.Fail(perror.New("Undeclared variable %q", id))
		return
	}

	env.Push(r)
}

func coCtxSet(env *Environment, in inst.Instruction) {

	id := in.Data.(string)
	v := env.Pop()

	if v == nil {
		env.Fail(perror.New("Assignment fail %q, value stack is empty", id))
		return
	}

	env.Bind(id, v)
}
