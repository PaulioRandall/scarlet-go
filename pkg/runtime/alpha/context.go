package alpha

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

// alphaContext implements pkg/runtime/Context.
type alphaContext struct {
	fixed  map[string]result
	vars   map[string]result
	parent *alphaContext
}

func (ctx alphaContext) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variables:" + NEWLINE

	if len(ctx.vars) == 0 && len(ctx.fixed) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for id, v := range ctx.fixed {
		s += TAB + "FIX " + id + " " + v.String() + NEWLINE
	}

	for id, v := range ctx.vars {
		s += TAB + id + " " + v.String() + NEWLINE
	}

	return
}

// Get returns an empty result if the ID does not exist.
func (ctx *alphaContext) Get(id string) result {
	v, ok := ctx.fixed[id]

	if !ok {
		v, ok = ctx.vars[id]

		if !ok {
			v = voidLiteral{}
		}
	}

	return v
}

func (ctx *alphaContext) SetFixed(id token.Token, v result) {

	name := id.Value

	if _, ok := ctx.fixed[name]; ok {
		panic(err("SetFixed", id, "Cannot change a fixed variable"))
	}

	delete(ctx.vars, name)
	ctx.fixed[name] = v
}

func (ctx *alphaContext) Set(id token.Token, v result) {

	name := id.Value

	if _, ok := ctx.fixed[name]; ok {
		panic(err("Set", id, "Cannot change a fixed variable"))
	}

	if _, ok := v.(voidLiteral); ok {
		delete(ctx.vars, name)
		return
	}

	ctx.vars[name] = v
}

func (ctx *alphaContext) Spawn() *alphaContext {

	fixed := make(map[string]result, len(ctx.fixed))

	for k, v := range ctx.fixed {
		fixed[k] = v
	}

	return &alphaContext{
		fixed:  fixed,
		vars:   make(map[string]result),
		parent: ctx,
	}
}

func (ctx *alphaContext) Parent() *alphaContext {
	return ctx.parent
}
