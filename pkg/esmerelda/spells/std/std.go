package std

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/registry"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/types"
)

func QuickRegister() {
	RegisterAll(func(name string, sp registry.Spell) {
		e := registry.Register(""+name, sp)
		if e != nil {
			panic(e)
		}
	})
}

func RegisterAll(reg registry.RegFunc) {
	reg("exit", Exit)
	reg("print", Print)
	reg("println", Println)
	reg("set", Set)
	reg("del", Del)
}

func Exit(env registry.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].(types.Num); ok {
		env.Exit(int(c.Integer()))
		return
	}

	env.Fail(errors.New("@Exit requires its argument be a number"))
}

func Print(_ registry.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func Println(_ registry.Enviro, args []types.Value) {
	Print(nil, args)
	fmt.Println()
}

func Set(env registry.Enviro, args []types.Value) {

	if len(args) != 2 {
		env.Fail(errors.New("@Set requires two arguments"))
		return
	}

	idStr, ok := args[0].(types.Str)
	id := string(idStr)

	if !ok || !isIdentifier(id) {
		env.Fail(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.Bind(id, args[1])
}

func Del(env registry.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Del requires one argument"))
		return
	}

	id, ok := args[0].(types.Str)
	if !ok {
		env.Fail(errors.New("@Del requires its argument be an identifier string"))
		return
	}

	env.Unbind(string(id))
}
