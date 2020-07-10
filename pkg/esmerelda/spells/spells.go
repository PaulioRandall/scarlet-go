package spells

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
)

func spell_exit(env Enviro, args []result.Result) {

	if len(args) != 1 {
		env.Fail(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].Num(); ok {
		env.Exit(uint32(c.Integer()))
		return
	}

	env.Fail(errors.New("@Exit requires its argument be a number"))
}

func spell_print(_ Enviro, args []result.Result) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func spell_println(_ Enviro, args []result.Result) {
	spell_print(nil, args)
	fmt.Println()
}

func spell_set(env Enviro, args []result.Result) {

	if len(args) != 2 {
		env.Fail(errors.New("@Set requires one argument"))
		return
	}

	id, ok := args[0].Str()
	if !ok {
		env.Fail(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.Bind(id, args[1])
}

func spell_del(env Enviro, args []result.Result) {

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
