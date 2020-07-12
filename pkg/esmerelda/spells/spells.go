package spells

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/registry"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/std"
)

func init() {

	ns := namespace("")
	std.RegisterAll(ns)
}

func namespace(prefix string) registry.RegFunc {
	return func(name string, sp registry.Spell) {
		e := registry.Register(prefix+name, sp)
		if e != nil {
			panic(e)
		}
	}
}

func LookUp(name string) registry.Spell {
	return registry.LookUp(name)
}

func Register(name string, sp registry.Spell) error {
	return registry.Register(name, sp)
}

func Unregister(name string) error {
	return registry.Unregister(name)
}
