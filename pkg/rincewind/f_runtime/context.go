package runtime

type environment struct {
	ctx      *context
	defs     map[string]result
	halt     bool
	exitCode int
	e        error
}

type context struct {
	*stack
	counter int
	vars    *map[string]result
}

func newEnv() *environment {

	ctx := &context{
		stack:   &stack{},
		counter: -1,
		vars:    &map[string]result{},
	}

	return &environment{
		ctx:      ctx,
		defs:     map[string]result{},
		exitCode: -1,
	}
}

func (env *environment) tick() int {
	c := env.ctx.counter
	env.ctx.counter++
	return c
}

func (env *environment) jump(idx int) {
	env.ctx.counter = idx - 1
}

func (env *environment) push(r result) {
	env.ctx.push(r)
}

func (env *environment) pop() result {
	return env.ctx.pop()
}

func (env *environment) get(id string) (result, bool) {

	v, ok := (*env.ctx.vars)[id]

	if !ok {
		v, ok = env.defs[id]
	}

	return v, ok
}

func (env *environment) put(id string, v result) {
	(*env.ctx.vars)[id] = v
}

func (env *environment) del(id string) {
	delete((*env.ctx.vars), id)
}

func (env *environment) def(id string, v result) bool {

	if _, ok := env.defs[id]; ok {
		return false
	}

	env.defs[id] = v
	return true
}
