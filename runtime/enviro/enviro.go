package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Environment struct {
	Ctx      *Context
	Defs     map[string]types.Value
	Halted   bool
	Err      error
	Done     bool
	ExitCode int
}

type Context struct {
	*types.Stack
	Counter  int
	Bindings *map[string]types.Value
}

func New() *Environment {

	ctx := &Context{
		Stack:    &types.Stack{},
		Bindings: &map[string]types.Value{},
	}

	return &Environment{
		Ctx:      ctx,
		Defs:     map[string]types.Value{},
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

func (env *Environment) Push(v types.Value) {
	env.Ctx.Push(v)
}

func (env *Environment) Pop() types.Value {
	return env.Ctx.Pop()
}

func (env *Environment) Get(id string) (types.Value, bool) {

	v, ok := (*env.Ctx.Bindings)[id]

	if !ok {
		v, ok = env.Defs[id]
	}

	return v, ok
}

func (env *Environment) Bind(id string, v types.Value) {
	(*env.Ctx.Bindings)[id] = v
}

func (env *Environment) Unbind(id string) {
	delete((*env.Ctx.Bindings), id)
}

func (env *Environment) Def(id string, v types.Value) bool {

	if _, ok := env.Defs[id]; ok {
		return false
	}

	env.Defs[id] = v
	return true
}

func (env *Environment) Exe(in inst.Instruction) {

	switch in.Code {
	case inst.CO_DELIM_PUSH:
		env.Push(types.Delim{})

	case inst.CO_VAL_PUSH:
		v := types.BuiltinValueOf(in.Data)
		env.Push(v)

	case inst.CO_CTX_GET:
		coCtxGet(env, in)

	case inst.CO_CTX_SET:
		coCtxSet(env, in)

	case inst.CO_ADD:
		coAdd(env, in)

	case inst.CO_SPELL:
		coSpell(env, in)

	default:
		env.Fail(perror.New("Unknown instruction code: %q", in.Code))
	}
}

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

	for v := env.Pop(); isNotDelim(v); v = env.Pop() {
		vs = append(vs, v)
	}

	return vs
}
