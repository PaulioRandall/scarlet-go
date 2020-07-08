package runtime

type environment struct {
	ctx  *context
	defs map[string]result
}

type context struct {
	counter int
	values  *stack
	vars    *map[string]result
}

func newEnv() *environment {

	ctx := &context{
		counter: -1,
		values:  &stack{},
		vars:    &map[string]result{},
	}

	return &environment{
		ctx:  ctx,
		defs: map[string]result{},
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
