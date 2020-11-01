package spells

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

const err_code = 1 // General program error

func setError(env spell.Runtime, m string, args ...interface{}) {
	s := fmt.Sprintf(m, args...)
	env.Fail(err_code, errors.New(s))
}

func Exit(env spell.Runtime, args []value.Value) []value.Value {

	if len(args) != 1 {
		setError(env, "@Exit requires one argument")
		return nil
	}

	c, ok := args[0].(value.Num)
	if !ok {
		setError(env, "@Exit requires its argument be a number")
		return nil
	}

	env.Exit(int(c.Integer()))
	return nil
}

func Print(env spell.Runtime, args []value.Value) []value.Value {
	for _, v := range args {
		fmt.Print(v.String())
	}
	return nil
}

func Println(env spell.Runtime, args []value.Value) []value.Value {
	Print(env, args)
	fmt.Println()
	return nil
}
