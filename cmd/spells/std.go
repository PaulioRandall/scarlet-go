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

func Len(env spell.Runtime, in []value.Value, out *spell.Output) {

	type lengthy interface {
		Len() int
	}

	if len(in) != 1 {
		setError(env, "@Len requires one argument")
		return
	}

	v, ok := in[0].(lengthy)
	if !ok {
		setError(env, "@Len argument has no length")
		return
	}

	out.Set(0, value.NewInt(v.Len()))
}

func Exit(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 1 {
		setError(env, "@Exit requires one argument")
		return
	}

	c, ok := in[0].(value.Num)
	if !ok {
		setError(env, "@Exit requires its argument be a number")
		return
	}

	env.Exit(int(c.Integer()))
}

func Print(env spell.Runtime, in []value.Value, _ *spell.Output) {
	for _, v := range in {
		fmt.Print(v.String())
	}
}

func Println(env spell.Runtime, in []value.Value, out *spell.Output) {
	Print(env, in, out)
	fmt.Println()
}

func ParseNum(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 1 {
		setError(env, "@ParseNum requires one argument")
		return
	}

	v, ok := in[0].(value.Str)
	if !ok {
		setError(env, "@ParseNum argument must be a string")
	}

	n, e := strconv.ParseFloat(string(v), 64)
	if e != nil {
		out.Set(1, value.Str("Unable to parse `"+string(v)+"`"))
		return
	}

	out.Set(0, value.NewFloat(n))
}
