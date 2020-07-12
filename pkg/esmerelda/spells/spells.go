package spells

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/registry"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/std"
)

func init() {
	reg("exit", std.Exit)
	reg("print", std.Print)
	reg("println", std.Println)
	reg("set", std.Set)
	reg("del", std.Del)
}

func reg(name string, sp registry.Spell) {
	e := registry.Register(name, sp)
	if e != nil {
		panic(e)
	}
}

func LookUp(name string) registry.Spell {
	return registry.LookUp(name)
}

func Register(name string, sp registry.Spell) error {
	return registry.Register(name, sp)
}
