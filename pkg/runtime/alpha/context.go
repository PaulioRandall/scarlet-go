package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// alphaContext implements pkg/runtime/Context.
type alphaContext struct {
	pure   bool
	fixed  map[string]result
	local  map[string]result
	parent *alphaContext
}

func (ctx alphaContext) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variables:" + NEWLINE

	if len(ctx.local) == 0 && len(ctx.fixed) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for id, v := range ctx.fixed {
		s += TAB + "FIX " + id + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.local {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return
}

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

func (ctx *alphaContext) SetFixed(id token.Token, v result) {

	name := id.Value

	if _, ok := ctx.fixed[name]; ok {
		panic(err("SetFixed", id, "Cannot change a fixed variable"))
	}

	delete(ctx.local, name)
	ctx.fixed[name] = v
}

func (ctx *alphaContext) SetLocal(id token.Token, v result) {

	for c := ctx; c != nil; c = c.parent {
		if _, ok := c.fixed[id.Value]; ok {
			panic(err("SetLocal", id, "Cannot change a fixed variable"))
		}
	}

	ctx.local[id.Value] = v
}

func (ctx *alphaContext) Set(id token.Token, v result) {
	if !ctx.set(id, v) {
		ctx.local[id.Value] = v
	}
}

func (ctx *alphaContext) set(id token.Token, v result) bool {

	varName := id.Value

	if _, ok := ctx.fixed[varName]; ok {
		panic(err("Set", id, "Cannot change a fixed variable"))
	}

	if _, ok := ctx.local[varName]; ok {
		ctx.setOrDelLocal(id, v)
		return true
	}

	return !ctx.pure &&
		ctx.parent != nil &&
		ctx.parent.set(id, v)
}

func (ctx *alphaContext) setOrDelLocal(id token.Token, v result) {
	if _, ok := v.(voidLiteral); ok {
		delete(ctx.local, id.Value)
	} else {
		ctx.local[id.Value] = v
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
