package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/runtime/result"
)

type Context struct {
	parent  *Context
	pure    bool
	defined map[string]Result
	local   map[string]Result
}

func NewCtx(parent *Context, pure bool) *Context {
	return &Context{
		parent:  parent,
		pure:    pure,
		defined: map[string]Result{},
		local:   map[string]Result{},
	}
}

func (ctx *Context) GetDefined(id string) (Result, bool) {

	for c := ctx; c != nil; c = ctx.parent {
		if def, ok := c.defined[id]; ok {
			return def, ok
		}
	}

	return Result{}, false
}

func (ctx *Context) SetDefined(id string, r Result) {
	ctx.defined[id] = r
}

func (ctx *Context) GetLocal(id string) (Result, bool) {
	v, ok := ctx.local[id]
	return v, ok
}

func (ctx *Context) SetLocal(id string, r Result) {
	ctx.local[id] = r
}

func (ctx *Context) GetVar(id string) (Result, bool) {

	for c := ctx; c != nil; c = ctx.parent {
		if v, ok := c.GetLocal(id); ok {
			return v, ok
		}
	}

	return Result{}, false
}

func (ctx *Context) SetVar(id string, r Result) {

	for c := ctx; c != nil; c = ctx.parent {
		if _, ok := c.GetLocal(id); ok {
			c.SetLocal(id, r)
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
		ctx.SetDefined(id, r)
		return
	}

	ctx.SetVar(id, r)
}

func (ctx Context) String() string {

	const NEWLINE = "\n"
	const TAB = "\t"

	s := "variables:" + NEWLINE

	if len(ctx.local) == 0 && len(ctx.defined) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return s
	}

	for def, v := range ctx.defined {
		s += TAB + "def " + def + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.local {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return s
}
