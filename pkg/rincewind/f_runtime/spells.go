package runtime

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

func invokeSpell(env *environment, in Instruction) {

	data := in.Data().([]interface{})
	argCount, val := data[0].(int), data[1].(string)

	switch strings.ToLower(val) {
	case "println":
		spell_println(popArgs(env, argCount))
	}

	perror.Panic("Unknown spell")
}

func popArgs(env *environment, size int) []result {

	rs := make([]result, size)

	for i := 0; i < size; i++ {
		rs[i] = env.pop()
		size--
	}

	return rs
}

func spell_println(args []result) {

	for _, v := range args {
		fmt.Print(v.String())
	}

	fmt.Println()
}
