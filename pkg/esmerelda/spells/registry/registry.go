package registry

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/result"
)

type Enviro interface {
	Exit(code int)
	Fail(error)
	Bind(id string, value result.Result)
	Unbind(id string)
}

type Spell func(env Enviro, args []result.Result)

var registry = map[string]Spell{}

func LookUp(name string) Spell {
	k := strings.ToLower(name)
	return registry[k]
}

func Register(name string, sp Spell) error {

	if !isSpellIdentifier(name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", name)
	}

	if sp == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", name)
	}

	k := strings.ToLower(name)
	curr := registry[k]

	if curr != nil {
		return fmt.Errorf("Spell with name %q already registered", name)
	}

	registry[k] = sp
	return nil
}

func reg(name string, sp Spell) {
	e := Register(name, sp)
	if e != nil {
		panic(e)
	}
}

func isSpellIdentifier(id string) bool {

	// E.g.
	// abc
	// abc.xyz
	// a.b.c.d

	newPart := true

	for _, ru := range id {

		switch {
		case newPart && unicode.IsLetter(ru):
			newPart = false

		case newPart:
			return false

		case ru == '.':
			newPart = true
		}
	}

	return !newPart
}
