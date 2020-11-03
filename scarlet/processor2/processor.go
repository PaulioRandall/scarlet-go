package processor2

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// Runtime is a handler for performing memory related and context dependent
// instructions such as access to the value stack and acces to scope variables.
type Runtime interface {

	// Spellbook returns the book containing the spells available during runtime.
	Spellbook() spell.Book

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

	// GetExitFlag returns true if the program should stop execution after
	// finishing any instruction currently being executed.
	GetExitFlag() bool
}

const (
	GENERAL_ERROR int = 1
)

func Execute(env Runtime, s tree.Stat) {
	switch v := s.(type) {
	case tree.SingleAssign:
		singleAssign(env, v)

	case tree.MultiAssign:
		multiAssign(env, v)

	case tree.SpellCall:
		spellCall(env, v)

	default:
		env.Fail(GENERAL_ERROR, errSnip(s.Pos(), "Unknown statement type"))
	}
}

func singleAssign(env Runtime, n tree.SingleAssign) {

}

func multiAssign(env Runtime, n tree.MultiAssign) {}

func spellCall(env Runtime, n tree.SpellCall) {}
