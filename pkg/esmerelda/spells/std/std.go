package std

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/spells/registry"
)

func Exit(env registry.Enviro, args []result.Result) {

	if len(args) != 1 {
		env.Fail(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].Num(); ok {
		env.Exit(int(c.Integer()))
		return
	}

	env.Fail(errors.New("@Exit requires its argument be a number"))
}

func Print(_ registry.Enviro, args []result.Result) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func Println(_ registry.Enviro, args []result.Result) {
	Print(nil, args)
	fmt.Println()
}

func Set(env registry.Enviro, args []result.Result) {

	if len(args) != 2 {
		env.Fail(errors.New("@Set requires two arguments"))
		return
	}

	id, ok := args[0].Str()
	if !ok || !isIdentifier(id) {
		env.Fail(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.Bind(id, args[1])
}

func Del(env registry.Enviro, args []result.Result) {

	if len(args) != 1 {
		env.Fail(errors.New("@Del requires one argument"))
		return
	}

	id, ok := args[0].Str()
	if !ok {
		env.Fail(errors.New("@Del requires its argument be an identifier string"))
		return
	}

	env.Unbind(id)
}
