package enviro

import (
	//"github.com/PaulioRandall/scarlet-go/shared/inst"
	//"github.com/PaulioRandall/scarlet-go/shared/perror"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Context struct {
	*types.Stack
	Counter int
	Sub     *SubContext
}

type SubContext struct {
	Parent   *SubContext
	Bindings *map[string]types.Value
}

func (ctx *Context) Get(id string) (types.Value, bool) {

	for s := ctx.Sub; s != nil; s = s.Parent {
		if v, ok := (*s.Bindings)[id]; ok {
			return v, true
		}
	}

	return nil, false
}

func (ctx *Context) Bind(id string, v types.Value) {
	sub := ctx.subOf(id)
	(*sub.Bindings)[id] = v
}

func (ctx *Context) Unbind(id string) {
	sub := ctx.subOf(id)
	delete((*sub.Bindings), id)
}

func (ctx *Context) subOf(id string) *SubContext {

	for sub := ctx.Sub; sub != nil; sub = sub.Parent {
		if _, ok := (*sub.Bindings)[id]; ok {
			return sub
		}
	}

	return ctx.Sub
}
