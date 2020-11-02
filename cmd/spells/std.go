package spells

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

const err_code = 1 // General program error

func setError(env spell.Runtime, m string, args ...interface{}) {
	s := fmt.Sprintf(m, args...)
	env.Fail(err_code, errors.New(s))
}

func ParseNum(env spell.Runtime, args []value.Value) []value.Value {

	if len(args) != 1 {
		setError(env, "@ParseNum requires one argument")
		return nil
	}

	if v, ok := args[0].(value.Str); ok {
		n, e := strconv.ParseFloat(string(v), 64)
		if e != nil {
			return []value.Value{
				value.NewFloat(0),
				value.Str("Unable to parse `" + string(v) + "`"),
			}
		}
		return []value.Value{value.NewFloat(n), value.Str("")}
	}

	setError(env, "@ParseNum argument must be a string")
	return nil
}

func Len(env spell.Runtime, args []value.Value) []value.Value {

	type lengthy interface {
		Len() int
	}

	if len(args) != 1 {
		setError(env, "@Len requires one argument")
		return nil
	}

	if v, ok := args[0].(lengthy); ok {
		return []value.Value{
			value.NewInt(v.Len()),
		}
	}

	setError(env, "@Len argument has no length")
	return nil
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
