package spellbook

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/cmd/runtime"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type (
	// Book represents a spellbook that stores collections of named spells.
	Book map[string]Spell

	// Spell represents a builtin function.
	Spell func(book Book, env runtime.Environment, args []value.Value)
)

// Register stores a named spell within the Book returning an erro if the name
// or spell are invalid.
func (b Book) Register(name string, spell Spell) error {

	if name == "" {
		panic(fmt.Errorf("Attempted to register a spell with no name"))
	}

	if !isSpellIdent(name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", name)
	}

	if spell == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", name)
	}

	k := strings.ToLower(name)
	b[k] = spell
	return nil
}

// Names returns the unsorted names of all spells in the Book.
func (b Book) Names() []string {
	keys := make([]string, len(b))
	var i int
	for k := range b {
		keys[i] = k
		i++
	}
	return keys
}

// Lookup returns the spell given its name. If the spell is nil then no such
// spell exists.
func (b Book) LookUp(name string) Spell {
	k := strings.ToLower(name)
	s, _ := b[k]
	return s
}

func isSpellIdent(id string) bool {

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
