package environment

type Environment struct {
	ctx      *context
	defs     map[string]Result
	halt     bool
	exitCode int
	e        error
}

type context struct {
	*stack
	counter int
	vars    *map[string]Result
}

func newEnv() *Environment {

	ctx := &context{
		stack: &stack{},
		vars:  &map[string]Result{},
	}

	return &Environment{
		ctx:      ctx,
		defs:     map[string]Result{},
		exitCode: -1,
	}
}

func (env *Environment) Err(e error) {
	env.e, env.halt = e, true
}

func (env *Environment) Tick() int {
	c := env.ctx.counter
	env.ctx.counter++
	return c
}

func (env *Environment) Jump(idx int) {
	env.ctx.counter = idx
}

func (env *Environment) Push(r Result) {
	env.ctx.push(r)
}

func (env *Environment) Pop() Result {
	return env.ctx.pop()
}

func (env *Environment) Get(id string) (Result, bool) {

	v, ok := (*env.ctx.vars)[id]

	if !ok {
		v, ok = env.defs[id]
	}

	return v, ok
}

func (env *Environment) Bind(id string, v Result) {
	(*env.ctx.vars)[id] = v
}

func (env *Environment) Del(id string) {
	delete((*env.ctx.vars), id)
}

func (env *Environment) Def(id string, v Result) bool {

	if _, ok := env.defs[id]; ok {
		return false
	}

	env.defs[id] = v
	return true
}
