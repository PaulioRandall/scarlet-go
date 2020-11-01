package spells

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/spellbook"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

const err_code = 1 // General program error

func setError(env spellbook.Runtime, m string, args ...interface{}) {
	s := fmt.Sprintf(m, args...)
	env.Fail(err_code, errors.New(s))
}

func Exit(_ spellbook.Book, env spellbook.Runtime, args []value.Value) []value.Value {

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

func Print(_ spellbook.Book, env spellbook.Runtime, args []value.Value) []value.Value {
	for _, v := range args {
		fmt.Print(v.String())
	}
	return nil
}

func Println(b spellbook.Book, env spellbook.Runtime, args []value.Value) []value.Value {
	Print(b, env, args)
	fmt.Println()
	return nil
}
