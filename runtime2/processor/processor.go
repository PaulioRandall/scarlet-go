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

	// Has returns true if the program 'counter' has not reached the end of
	// the instruction list.
	Has(Counter) bool

	// Fetch returns the instruction specified by the program 'counter'.
	Fetch(Counter) (inst.Inst, error)

	// Get returns the value associated with the specified identifier.
	Get(value.Ident) (value.Value, error)

	// Bind sets the value of a variable overwriting any existing value.
	Bind(value.Ident, value.Value) error
}

// Processor executes instructions in a similar fashion to a CPU but at a
// higher level.
type Processor struct {
	Memory  Memory  // Access to memory, i.e. instructions and variables
	Counter Counter // Program counter
	Stop    bool    // True to interupt execution after the next instruction
	Stopped bool    // True if execution was stopped by an interupt or error
	Halt    bool    // True to halt execution, invoked only by instructions
	Err     error
}

// PleaseStop tells the processor to stop execution after finishing the current
// instruction. 'Processor.Stopped' will be set to true upon stopping.
func (p *Processor) PleaseStop() {
	p.Stop = true
}

// Run begins or continues execution of instructions and returns true if the
// execution halted because the program counter reached the end of the
// instruction list or a halt instruction was encountered. False is returned if
// execution was requested to halt via a separate process, the 'Processor.Err'
// value contains an error from previous execution, or an error occurred
// resulting in the 'Processor.Err' value being set.
func (p *Processor) Run() {

	var in inst.Inst

	if p.Err != nil {
		p.Stopped = true
		return
	}

	p.Stop = false
	p.Stopped = false
	p.Halt = false

	for !p.Halt && p.Memory.Has(p.Counter) {

		if p.Stop {
			p.Stopped = true
			return
		}

		if in, p.Err = p.Memory.Fetch(p.Counter); p.Err != nil {
			p.Stopped = true
			return
		}

		if p.Halt, p.Err = Process(p.Memory, in); p.Err != nil {
			p.Stopped = true
			return
		}

		p.Counter++ // Increment when on 'Halt' but never on 'Stop'
	}
}

// Process the instruction 'in' using the memory 'm'. 'halt' should only be
// returned as true if an instruction specifically requests execution to halt.
func Process(m Memory, in inst.Inst) (halt bool, e error) {
	// TODO
	return
}
