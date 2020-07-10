package spells

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
)

type Enviro interface {
	Exit(code uint32)
	Fail(error)
	Bind(id string, value result.Result)
	Unbind(id string)
}

func Invoke(env Enviro, name string, args []result.Result) {

	switch strings.ToLower(name) {
	case "exit":
		spell_exit(env, args)

	case "print":
		spell_print(env, args)

	case "println":
		spell_println(env, args)

	case "set":
		spell_set(env, args)

	case "del":
		spell_del(env, args)

	default:
		perror.Panic("Unknown spell %q", name)
	}
}
