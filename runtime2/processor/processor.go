package processor

import (
	"github.com/PaulioRandall/scarlet-go/token2/inst"
	"github.com/PaulioRandall/scarlet-go/token2/value"
)

// Counter represents a program counter wihtin a processor.
type Counter uint

// Memory represents the source of instructions and handler for performing
// context dependent instructions such as access to variables.
type Memory interface {

	// More returns true if the program 'counter' has not reached the end of
	// the instruction list.
	More(Counter) bool

	// Next returns the instruction specified by the program 'counter'.
	Next(Counter) (inst.Inst, error)

	// Fetch returns the value associated with the specified identifier.
	Fetch(value.Ident) (value.Value, error)

	// Bind sets the value of a variable overwriting any existing value.
	Bind(value.Ident, value.Value) error
}

// Processor executes instructions in a similar fashion to a CPU but at a
// higher level.
type Processor struct {
	Memory  Memory
	Counter Counter
	Halted  bool
	Err     error
}

