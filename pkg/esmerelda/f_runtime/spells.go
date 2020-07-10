package runtime

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

func invokeSpell(env *environment, in inst.Instruction) {

	data := in.Data().([]interface{})
	argCount, val := data[0].(int), data[1].(string)

	switch strings.ToLower(val) {
	case "exit":
		spell_exit(env, popArgs(env, argCount))

	case "print":
		spell_print(env, popArgs(env, argCount))

	case "println":
		spell_println(env, popArgs(env, argCount))

	case "set":
		spell_set(env, popArgs(env, argCount))

	case "del":
		spell_del(env, popArgs(env, argCount))

	default:
		perror.Panic("Unknown spell %q", val)
	}
}

func popArgs(env *environment, size int) []result {

	rs := make([]result, size)

	for i := size - 1; i >= 0; i-- {
		rs[i] = env.pop()
		size--
	}

	return rs
}

func spell_exit(env *environment, args []result) {

	if len(args) != 1 {
		env.err(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].Num(); ok {
		env.exitCode = int(c.Integer())
		env.halt = true
		return
	}

	env.err(errors.New("@Exit requires its argument be a number"))
}

func spell_print(env *environment, args []result) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func spell_println(env *environment, args []result) {
	spell_print(env, args)
	fmt.Println()
}

func spell_set(env *environment, args []result) {

	if len(args) != 2 {
		env.err(errors.New("@Set requires one argument"))
		return
	}

	id, ok := args[0].Str()
	if !ok {
		env.err(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.put(id, args[1])
}

func spell_del(env *environment, args []result) {

	if len(args) != 1 {
		env.err(errors.New("@Del requires one argument"))
		return
	}

	id, ok := args[0].Str()
	if !ok {
		env.err(errors.New("@Del requires its argument be an identifier string"))
		return
	}

	env.del(id)
}
