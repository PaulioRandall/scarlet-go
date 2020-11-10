package spell

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type (
	// Spell represents a builtin function.
	Spell func(env Runtime, in []value.Value, out *Output)

	// Output is a container for spell return arguments.
	Output struct {
		size int
		out  []value.Value
	}

	// Book represents a collections of named spells.
	Book map[string]Inscription

	// Inscription represents a spell inscribed within a spell book.
	Inscription struct {
		Spell
		Name    string
		Outputs int
	}

	// Scope represents a mapping of declared identifiers, with their current
	// values, available within the current scope.
	Scope map[value.Ident]value.Value

	// Runtime is a handler for performing memory related and context dependent
	// instructions such as access to scope variables and storing exit and error
	// information for the processor. It's a subset of the Runtime used by the
	// Processor that only exposes appropriate functionality for spells.
	Runtime interface {

		// Spellbook returns the book containing all spells available. Changes made
		// will not be reflected within the current environment.
		Spellbook() Book

		// Scope returns a copy of the current scope. Changes made will not be
		// reflected within the current environment.
		Scope() Scope

		// Exists returns true if the specified identifier exists within the current
		//scope.
		Exists(value.Ident) value.Bool

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

// Inscribe stores a named spell within the Book returning an error if any of
// the arguments are invalid.
func (b Book) Inscribe(name string, outputs int, spell Spell) error {

	if name == "" {
		panic(fmt.Errorf("Attempted to register a spell with no name"))
	}

	if !isSpellIdent(name) {
		return fmt.Errorf("Attempted to register spell with bad name %q", name)
	}

	if spell == nil {
		return fmt.Errorf("Attempted to register nil spell with name %q", name)
	}

	if outputs < 0 {
		return fmt.Errorf("Attempted to register spell"+
			" with variable or negative output parameters %q", name)
	}

	k := strings.ToLower(name)
	b[k] = Inscription{
		Spell:   spell,
		Name:    name,
		Outputs: outputs,
	}

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
func (b Book) Lookup(name string) (Inscription, bool) {
	k := strings.ToLower(name)
	s, ok := b[k]
	return s, ok
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

// NewOutput returns a new initialised output.
func NewOutput(size int) *Output {
	return &Output{
		size: size,
		out:  make([]value.Value, size),
	}
}

// Get returns the return value of the index 'i' or nil if it has not been set
// yet.
func (o *Output) Get(i int) value.Value {
	if i >= o.size || i < 0 {
		panic("Out of range: invalid spell output index")
	}
	return o.out[i]
}

// Set sets the value of a spell return value.
func (o *Output) Set(i int, v value.Value) {
	if i >= o.size || i < 0 {
		panic("Out of range: invalid spell output index")
	}
	o.out[i] = v
}

// Slice returns a slice of all return values. Note that unset slots will be
// nil.
func (o *Output) Slice() []value.Value {
	return o.out
}
