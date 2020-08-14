package spellbook

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Enviro interface {
	Exit(code int)
	Fail(error)
	Bind(id string, value types.Value)
	Unbind(id string)
}

type Spell interface {
	Summary() string
	Invoke(env Enviro, args []types.Value)
}

type Inscriber func(name string, spell Spell)

var spellBook = map[string]Spell{}

func LookUp(name string) Spell {
	k := strings.ToLower(name)
	return spellBook[k]
}

func SpellNames() []string {

	keys := make([]string, len(spellBook))
	var i int

	for k := range spellBook {
		keys[i] = k
		i++
	}

	return keys
}

func Inscribe(name string, sp Spell) error {

	if !isSpellIdentifier(name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", name)
	}

	if sp == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", name)
	}

	k := strings.ToLower(name)
	curr := spellBook[k]

	if curr != nil {
		return fmt.Errorf("Spell with name %q already registered", name)
	}

	spellBook[k] = sp
	return nil
}

func ScribbleOut(name string) error {

	if !isSpellIdentifier(name) {
		return fmt.Errorf("Attempted to unregister spell with bad name %q", name)
	}

	k := strings.ToLower(name)
	delete(spellBook, k)
	return nil
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
