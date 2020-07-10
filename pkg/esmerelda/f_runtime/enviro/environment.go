package enviro

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

type Environment struct {
	Ctx      *Context
	Defs     map[string]Result
	Halted   bool
	Err      error
	ExitCode int
}

type Context struct {
	*Stack
	Counter  int
	Bindings *map[string]Result
}

func New() *Environment {

	ctx := &Context{
		Stack:    &Stack{},
		Bindings: &map[string]Result{},
	}

	return &Environment{
		Ctx:      ctx,
		Defs:     map[string]Result{},
		Halted:   true,
		ExitCode: -1,
	}
}

func (env *Environment) Fail(e error) {
	env.Err, env.Halted = e, true
}

func (env *Environment) Tick() int {
	c := env.Ctx.Counter
	env.Ctx.Counter++
	return c
}

func (env *Environment) Jump(idx int) {
	env.Ctx.Counter = idx
}

func (env *Environment) Push(r Result) {
	env.Ctx.Push(r)
}

func (env *Environment) Pop() Result {
	return env.Ctx.Pop()
}

func (env *Environment) Get(id string) (Result, bool) {

	v, ok := (*env.Ctx.Bindings)[id]

	if !ok {
		v, ok = env.Defs[id]
	}

	return v, ok
}

func (env *Environment) Bind(id string, v Result) {
	(*env.Ctx.Bindings)[id] = v
}

func (env *Environment) Del(id string) {
	delete((*env.Ctx.Bindings), id)
}

func (env *Environment) Def(id string, v Result) bool {

	if _, ok := env.Defs[id]; ok {
		return false
	}

	env.Defs[id] = v
	return true
}

func (env *Environment) Exe(in inst.Instruction) {

	switch in.Code() {
	case IN_VAL_PUSH:
		env.Push(Result{
			ty:  resultTypeOf(in.Data()),
			val: in.Data(),
		})

	case IN_CTX_GET:
		id := in.Data().(string)
		r, ok := env.Get(id)

		if ok {
			env.Push(r)
		} else {
			msg := fmt.Sprintf("Undeclared variable %q", id)
			env.Fail(perror.NewBySnippet(msg, in))
		}

	case IN_SPELL:
		invokeSpell(env, in)

	default:
		env.Fail(perror.NewBySnippet("Unknown instruction code", in))
	}
}
