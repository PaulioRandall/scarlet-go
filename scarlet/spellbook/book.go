package spellbook

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type (
	// Book represents a collections of named spells.
	Book map[string]Spell

	// Spell represents a builtin function... because to the readers and writers
	// of scrolls, it's indistinguishable from magic.
	Spell func(book Book, env Runtime, args []value.Value)

	// Runtime is a handler for performing memory related and context dependent
	// instructions such as access to scope variables and storing exit and error
	// information for the processor. It's a subset of the Runtime used by the
	// Processor that only exposes appropriate functionality for spells.
	Runtime interface {

		// Bind sets the value of a variable overwriting any existing value.
		Bind(value.Ident, value.Value)

		// Fetch returns the value associated with the specified identifier.
		Fetch(value.Ident) value.Value

		// Fail sets the error and exit status a non-recoverable error occurs
		// during execution.
		Fail(int, error)

		// Exit causes the program to exit with the specified exit code.
		Exit(int)

		// GetErr returns the error if set else returns nil.
		GetErr() error

		// GetExitCode returns the currently set exit code. Only meaningful if the
		// exit flag has been set.
		GetExitCode() int

		// GetExitFlag returns true if the program should stop execution after
		// finishing any instruction currently being executed.
		GetExitFlag() bool
	}
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
