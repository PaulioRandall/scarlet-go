package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/runtime/result"
	//"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	//. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type Context struct {
	parent  *Context
	pure    bool
	defined map[string]Result
	local   map[string]Result
}

func NewCtx(parent *Context, pure bool) Context {
	return Context{
		parent:  parent,
		pure:    true,
		defined: map[string]Result{},
		local:   map[string]Result{},
	}
}

func (ctx Context) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variables:" + NEWLINE

	if len(ctx.local) == 0 && len(ctx.defined) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for def, v := range ctx.defined {
		s += TAB + "def " + def + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.local {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return
}

/*
func (ctx *alphaContext) Get(id string) result {

	if v := ctx.getFixed(id); v != nil {
		return v
	}

	if v := ctx.getVar(id); v != nil {
		return v
	}

	return nil
}

func (ctx *alphaContext) GetLocal(id string) result {

	if v, ok := ctx.local[id]; ok {
		return v
	}

	return nil
}

func (ctx *alphaContext) getFixed(id string) result {

	for c := ctx; c != nil; c = c.parent {
		if v, ok := c.fixed[id]; ok {
			return v
		}
	}

	return nil
}

func (ctx *alphaContext) getVar(id string) result {

	for c := ctx; c != nil; c = c.parent {

		if v, ok := c.local[id]; ok {
			return v
		}

		if c.pure {
			return nil
		}
	}

	return nil
}

func (ctx *alphaContext) SetFixed(id Token, v result) {

	name := id.Value()

	if _, ok := ctx.fixed[name]; ok {
		err.Panic("Cannot change a fixed variable", err.At(id))
	}

	delete(ctx.local, name)
	ctx.fixed[name] = v
}

func (ctx *alphaContext) SetLocal(id Token, v result) {

	for c := ctx; c != nil; c = c.parent {
		if _, ok := c.fixed[id.Value()]; ok {
			err.Panic("Cannot change a fixed variable", err.At(id))
		}
	}

	ctx.local[id.Value()] = v
}

func (ctx *alphaContext) Set(id Token, v result) {
	if !ctx.set(id, v) {
		ctx.setOrDelLocal(id, v)
	}
}

func (ctx *alphaContext) set(id Token, v result) bool {

	varName := id.Value()

	if _, ok := ctx.fixed[varName]; ok {
		err.Panic("Cannot change a fixed variable", err.At(id))
	}

	if _, ok := ctx.local[varName]; ok {
		ctx.setOrDelLocal(id, v)
		return true
	}

	return !ctx.pure &&
		ctx.parent != nil &&
		ctx.parent.set(id, v)
}

func (ctx *alphaContext) setOrDelLocal(id Token, v result) {
	if _, ok := v.(voidLiteral); ok {
		delete(ctx.local, id.Value())
	} else {
		ctx.local[id.Value()] = v
	}
}

func (ctx *alphaContext) Spawn(pure bool) *alphaContext {
	return &alphaContext{
		pure:   pure,
		fixed:  map[string]result{},
		local:  map[string]result{},
		parent: ctx,
	}
}

func (ctx *alphaContext) Parent() *alphaContext {
	return ctx.parent
}
*/
