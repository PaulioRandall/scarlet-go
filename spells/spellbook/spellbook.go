package spellbook

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func init() {
	manual.Register("spells", func() string {

		names := SpellNames()
		sort.Strings(names)

		sb := strings.Builder{}

		for i, v := range names {

			sp, _ := LookUp(v)

			if i != 0 {
				sb.WriteString("\n\n")
			}

			sb.WriteString(sp.Sig)
			sb.WriteString("\n\t")
			sb.WriteString(sp.Desc)
		}

		return sb.String()
	})
}

type Enviro interface {
	Exit(code int)
	Fail(error)
	Bind(id string, value types.Value)
	Unbind(id string)
}

type Spell func(spell Entry, env Enviro, args []types.Value)
type Entry struct {
	Name     string
	Sig      string
	Desc     string
	Examples []string
	Spell    Spell
}

var spellBook = map[string]Entry{}

func Register(sp Entry) error {

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

	if _, exists := spellBook[k]; exists {
		return fmt.Errorf("Spell with name %q already registered", sp.Name)
	}

	spellBook[k] = sp
	return nil
}

func LookUp(name string) (Entry, bool) {
	k := strings.ToLower(name)
	sp, ok := spellBook[k]
	return sp, ok
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
