package runtime

type Context struct {
	parent  *Context
	pure    bool
	defined *map[string]Result
	local   map[string]Result
}

func NewCtx(parent *Context, pure bool) *Context {
	return &Context{
		parent:  parent,
		pure:    pure,
		defined: &map[string]Result{},
		local:   map[string]Result{},
	}
}

func (ctx *Context) Definitions() map[string]Result {
	return *ctx.defined
}

func (ctx *Context) Locals() map[string]Result {
	return ctx.local
}

func (ctx *Context) GetDefined(id string) (Result, bool) {

	v, ok := (*ctx.defined)[id]
	if !ok {
		v = Result{}
	}

	return v, ok
}

func (ctx *Context) SetDefinition(id string, r Result) {
	(*ctx.defined)[id] = r
}

func (ctx *Context) GetLocal(id string) (Result, bool) {
	v, ok := ctx.local[id]
	return v, ok
}

func (ctx *Context) SetLocal(id string, r Result) {
	ctx.local[id] = r
}

func (ctx *Context) GetVar(id string) (Result, bool) {

	for c := ctx; c != nil; c = c.parent {
		if v, ok := c.GetLocal(id); ok {
			return v, true
		}

		if c.pure {
			break
		}
	}

	return Result{}, false
}

func (ctx *Context) SetVar(id string, r Result) {

	for c := ctx; c != nil; c = c.parent {
		if _, ok := c.GetLocal(id); ok {
			c.SetLocal(id, r)
			return
		}

		if c.pure {
			break
		}
	}

	ctx.SetLocal(id, r)
}

func (ctx *Context) Get(id string) (Result, bool) {

	v, ok := ctx.GetDefined(id)
	if ok {
		return v, true
	}

	return ctx.GetVar(id)
}

func (ctx *Context) Set(final bool, id string, r Result) {

	if final {
		ctx.SetDefinition(id, r)
		return
	}

	ctx.SetVar(id, r)
}

func (ctx Context) String() string {

	const NEWLINE = "\n"
	const TAB = "\t"

	s := "variables:" + NEWLINE

	if len(ctx.local) == 0 && len(*ctx.defined) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return s
	}

	for def, v := range *ctx.defined {
		s += TAB + "def " + def + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.local {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return s
}
