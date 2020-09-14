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

type Spellbook map[string]Entry
type Spell func(spell Entry, env Enviro, args []types.Value)
type Entry struct {
	Spell    Spell
	Name     string
	Sig      string
	Desc     string
	Examples []string
}

func (spbk Spellbook) Register(
	spell Spell,
	name string,
	sig string,
	desc string,
	exam ...string,
) error {

	sp := Entry{
		Spell:    spell,
		Name:     name,
		Sig:      sig,
		Desc:     desc,
		Examples: exam,
	}

	return spbk.RegisterEntry(sp)
}

func (spbk Spellbook) RegisterEntry(sp Entry) error {

	if sp.Name == "" {
		panic(fmt.Errorf("Attempted to register a spell with no name"))
	}

	if !isSpellIdentifier(sp.Name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", sp.Name)
	}

	if sp.Spell == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", sp.Name)
	}

	k := strings.ToLower(sp.Name)

	if _, exists := spbk[k]; exists {
		return fmt.Errorf("Spell with name %q already registered", sp.Name)
	}

	spbk[k] = sp
	return nil
}

func (spbk Spellbook) Names() []string {

	keys := make([]string, len(spbk))

	var i int
	for k := range spbk {
		keys[i] = k
		i++
	}

	return keys
}

func (spbk Spellbook) LookUp(name string) (Entry, bool) {
	k := strings.ToLower(name)
	sp, ok := spbk[k]
	return sp, ok
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
