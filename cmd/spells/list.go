package spells

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

func List_New(env spell.Runtime, in []value.Value, out *spell.Output) {
	list := make([]value.Value, len(in))
	for i, v := range in {
		list[i] = v
	}
	out.Set(0, value.List(list))
}
