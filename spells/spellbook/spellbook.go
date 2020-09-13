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

			sp, _ := LookUpDoc(v)

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

type Spell func(env Enviro, args []types.Value)
type SpellDoc struct {
	Name     string
	Sig      string
	Desc     string
	Examples []string
}
type Inscriber func(name string, spell Spell)

var spellBook = map[string]Spell{}
var spellDocs = map[string]SpellDoc{}

func Register(sp Spell, doc SpellDoc) error {

	if doc.Name == "" {
		panic(fmt.Errorf("Attempted to register a spell with no name"))
	}

	if !isSpellIdentifier(doc.Name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", doc.Name)
	}

	if sp == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", doc.Name)
	}

	k := strings.ToLower(doc.Name)
	curr := spellBook[k]

	if curr != nil {
		return fmt.Errorf("Spell with name %q already registered", doc.Name)
	}

	spellBook[k] = sp
	spellDocs[k] = doc
	return nil
}

func LookUp(name string) Spell {
	k := strings.ToLower(name)
	return spellBook[k]
}

func LookUpDoc(name string) (SpellDoc, bool) {
	name = strings.ToLower(name)
	s, ok := spellDocs[name]
	return s, ok
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
