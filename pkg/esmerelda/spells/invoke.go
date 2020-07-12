package spells

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
)

type Spell func(env Enviro, args []result.Result)

type Enviro interface {
	Exit(code int)
	Fail(error)
	Bind(id string, value result.Result)
	Unbind(id string)
}

var register = map[string]Spell{}

func Invoke(env Enviro, name string, args []result.Result) {

	k := strings.ToLower(name)
	sp := register[k]

	if sp == nil {
		perror.Panic("Unknown spell %q", name)
	}

	sp(env, args)
}
