package enviro

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst/codes"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells"
)

type Environment struct {
	Ctx      *Context
	Defs     map[string]result.Result
	Halted   bool
	Err      error
	Done     bool
	ExitCode int
}

type Context struct {
	*Stack
	Counter  int
	Bindings *map[string]result.Result
}

func New() *Environment {

	ctx := &Context{
		Stack:    &Stack{},
		Bindings: &map[string]result.Result{},
	}

	return &Environment{
		Ctx:      ctx,
		Defs:     map[string]result.Result{},
		Halted:   true,
		ExitCode: 0,
	}
}

func (env *Environment) Exit(code int) {
	env.Done, env.ExitCode, env.Halted = true, code, true
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

func (env *Environment) Push(r result.Result) {
	env.Ctx.Push(r)
}

func (env *Environment) Pop() result.Result {
	return env.Ctx.Pop()
}

func (env *Environment) Get(id string) (result.Result, bool) {

	v, ok := (*env.Ctx.Bindings)[id]

	if !ok {
		v, ok = env.Defs[id]
	}

	return v, ok
}

func (env *Environment) Bind(id string, v result.Result) {
	(*env.Ctx.Bindings)[id] = v
}

func (env *Environment) Unbind(id string) {
	delete((*env.Ctx.Bindings), id)
}

func (env *Environment) Def(id string, v result.Result) bool {

	if _, ok := env.Defs[id]; ok {
		return false
	}

	env.Defs[id] = v
	return true
}

func (env *Environment) Exe(in inst.Instruction) {

	switch in.Code() {
	case IN_VAL_PUSH:
		env.Push(result.Result{
			RType: result.ResultTypeOf(in.Data()),
			Value: in.Data(),
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

func popArgs(env *Environment, size int) []result.Result {

	rs := make([]result.Result, size)

	for i := size - 1; i >= 0; i-- {
		rs[i] = env.Pop()
		size--
	}

	return rs
}

func invokeSpell(env *Environment, in inst.Instruction) {

	data := in.Data().([]interface{})
	argCount, name := data[0].(int), data[1].(string)

	sp := spells.LookUp(name)
	if sp == nil {
		msg := fmt.Sprintf("Unknown spell %q", name)
		env.Fail(perror.NewBySnippet(msg, in))
		return
	}

	args := popArgs(env, argCount)
	sp(env, args)
}
