package enviro

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/perror"
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

func New() *Environment {

	ctx := &Context{
		Stack: &types.Stack{},
	}

	ctx.PushSub()

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

func (env *Environment) JumpBy(amount int) {
	env.Ctx.Counter += amount
}

func (env *Environment) PushVal(v types.Value) {
	env.Ctx.Push(v)
}

func (env *Environment) PopVal() types.Value {
	return env.Ctx.Pop()
}

func (env *Environment) Get(id string) (types.Value, bool) {

	v, ok := env.Ctx.Get(id)

	if !ok {
		v, ok = env.Defs[id]
	}

	return v, ok
}

func (env *Environment) Bind(id string, v types.Value) {
	env.Ctx.Bind(id, v)
}

func (env *Environment) Unbind(id string) {
	env.Ctx.Unbind(id)
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
		env.PushVal(types.Delim{})

	case inst.CO_VAL_PUSH:
		v := types.BuiltinValueOf(in.Data)
		env.PushVal(v)

	case inst.CO_VAL_GET: // co_context.go
		coCtxGet(env, in)

	case inst.CO_VAL_BIND:
		coCtxSet(env, in)

	case inst.CO_ADD: // co_operation.go
		coAdd(env, in)

	case inst.CO_SUB:
		coSub(env, in)

	case inst.CO_MUL:
		coMul(env, in)

	case inst.CO_DIV:
		coDiv(env, in)

	case inst.CO_REM:
		coRem(env, in)

	case inst.CO_AND:
		coAnd(env, in)

	case inst.CO_OR:
		coOr(env, in)

	case inst.CO_LESS:
		coLess(env, in)

	case inst.CO_MORE:
		coMore(env, in)

	case inst.CO_LESS_EQU:
		coLessOrEqual(env, in)

	case inst.CO_MORE_EQU:
		coMoreOrEqual(env, in)

	case inst.CO_EQU:
		coEqual(env, in)

	case inst.CO_NOT_EQU:
		coNotEqual(env, in)

	case inst.CO_SPELL: // co_call.go
		coSpell(env, in)

	case inst.CO_JUMP_FALSE: // co_jump.go
		coJump(env, in, false)

	case inst.CO_SUB_CTX_PUSH:
		env.Ctx.PushSub()

	case inst.CO_SUB_CTX_POP:
		env.Ctx.PopSub()

	default:
		env.Fail(perror.New("Unknown instruction code: %q", in.Code))
	}
}
