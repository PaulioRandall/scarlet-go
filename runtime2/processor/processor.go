package processor

import (
	"github.com/PaulioRandall/scarlet-go/token2/code"
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
	Memory  Memory      // Access to memory, i.e. instructions and variables
	Counter Counter     // Program counter
	Stack   value.Stack // Value stack
	Stop    bool        // True to interupt execution after the next instruction
	Stopped bool        // True if execution was stopped by an interupt or error
	Halt    bool        // True to halt execution, invoked only by instructions
	Err     error
}

// New returns a new Processor with the specified memory installed.
func New(m Memory) *Processor {
	return &Processor{Memory: m}
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

		if p.Halt, p.Err = Process(p, in); p.Err != nil {
			p.Stopped = true
			return
		}

		p.Counter++ // Increment when on 'Halt' but never on 'Stop'
	}
}

// Process the instruction 'in' using the memory 'm'. 'halt' should only be
// returned as true if an instruction specifically requests execution to halt.
func Process(p *Processor, in inst.Inst) (halt bool, e error) {
	switch {
	case in.Code == code.STACK_PUSH:
		p.Stack.Push(in.Data)
	case in.Code == code.SCOPE_BIND:
		e = p.Memory.Bind(in.Data.(value.Ident), p.Stack.Pop())
	case processNumOp(p, in):
	default:
		panic("Unhandled instruction code: " + in.Code.String())
	}
	return
}

func processNumOp(p *Processor, in inst.Inst) bool {

	binNumOp := func(f func(l, r *value.Num)) {
		r := p.Stack.Pop().(value.Num)
		l := p.Stack.Pop().(value.Num)
		l.Number = l.Number.Copy()
		f(&l, &r) // Answer is always held in the left value
		p.Stack.Push(l)
	}

	switch in.Code {
	case code.OP_ADD:
		binNumOp(func(l, r *value.Num) { l.Number.Add(r.Number) })
	case code.OP_SUB:
		binNumOp(func(l, r *value.Num) { l.Number.Sub(r.Number) })
	case code.OP_MUL:
		binNumOp(func(l, r *value.Num) { l.Number.Mul(r.Number) })
	case code.OP_DIV:
		binNumOp(func(l, r *value.Num) { l.Number.Div(r.Number) })
	case code.OP_REM:
		binNumOp(func(l, r *value.Num) { l.Number.Mod(r.Number) })

	case code.OP_AND:
		answer := p.Stack.Pop().(value.Bool) && p.Stack.Pop().(value.Bool)
		p.Stack.Push(answer)
	case code.OP_OR:
		answer := p.Stack.Pop().(value.Bool) || p.Stack.Pop().(value.Bool)
		p.Stack.Push(answer)

	case code.OP_LESS:
		binNumOp(func(l, r *value.Num) { l.Number.Less(r.Number) })
	case code.OP_MORE:
		binNumOp(func(l, r *value.Num) { l.Number.More(r.Number) })
	case code.OP_LEQU:
		binNumOp(func(l, r *value.Num) { l.Number.LessOrEqual(r.Number) })
	case code.OP_MEQU:
		binNumOp(func(l, r *value.Num) { l.Number.MoreOrEqual(r.Number) })

	case code.OP_EQU:
		r, l := p.Stack.Pop(), p.Stack.Pop()
		p.Stack.Push(value.Bool(l.Equal(r)))
	case code.OP_NEQU:
		r, l := p.Stack.Pop(), p.Stack.Pop()
		p.Stack.Push(value.Bool(!l.Equal(r)))

	default:
		return false
	}

	return true
}
