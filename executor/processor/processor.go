package processor

import (
	"github.com/PaulioRandall/scarlet-go/token/code"
	"github.com/PaulioRandall/scarlet-go/token/inst"
	"github.com/PaulioRandall/scarlet-go/token/value"
)

// TODO: Needs testing!

// Runtime represents the source of instructions and handler for performing
// context dependent instructions such as access to variables.
type Runtime interface {

	// Next returns the next instruction inicated by the program counter. True is
	// returned if an instruction was returned otherwise the end of program has
	// been reached.
	Next() (inst.Inst, bool)

	// Push a value onto the top of the value stack.
	Push(value.Value)

	// Pop a value off the top of the value stack,
	Pop() value.Value

	// Fetch returns the value associated with the specified identifier.
	Fetch(value.Ident) (value.Value, error)

	// Bind sets the value of a variable overwriting any existing value.
	Bind(value.Ident, value.Value) error
}

// Processor executes instructions in a similar fashion to a CPU but at a
// higher level.
type Processor struct {
	Runtime Runtime // Access to memory etc, e.g. instructions and variables
	Stop    bool    // True to interupt execution after the next instruction
	Stopped bool    // True if execution was stopped by an interupt or error
	Halt    bool    // True to halt execution, invoked only by instructions
	Err     error
}

// New returns a new Processor with the specified memory installed.
func New(rt Runtime) *Processor {
	return &Processor{Runtime: rt}
}

// Run begins or continues execution of instructions and returns true if the
// execution halted because the program counter reached the end of the
// instruction list or a halt instruction was encountered. False is returned if
// execution was requested to halt via a separate process, the 'Processor.Err'
// value contains an error from previous execution, or an error occurred
// resulting in the 'Processor.Err' value being set.
func (p *Processor) Run() {

	var in inst.Inst
	var ok bool

	if p.Err != nil {
		p.Stopped = true
		return
	}

	p.Stop = false
	p.Stopped = false
	p.Halt = false

	for !p.Halt {

		if p.Stop {
			p.Stopped = true
			return
		}

		if in, ok = p.Runtime.Next(); !ok {
			return
		}

		if p.Halt, p.Err = p.Process(in); p.Err != nil {
			p.Stopped = true
			return
		}
	}
}

// Process the instruction 'in' using the memory 'm'. 'halt' should only be
// returned as true if an instruction specifically requests execution to halt.
func (p *Processor) Process(in inst.Inst) (halt bool, e error) {
	switch {
	case in.Code == code.STACK_PUSH:
		p.Runtime.Push(in.Data)
	case in.Code == code.SCOPE_BIND:
		e = p.Runtime.Bind(in.Data.(value.Ident), p.Runtime.Pop())
	case processNumOp(p, in):
	default:
		panic("Unhandled instruction code: " + in.Code.String())
	}
	return
}

func processNumOp(p *Processor, in inst.Inst) bool {

	binNumOp := func(f func(l, r *value.Num)) {
		r := p.Runtime.Pop().(value.Num)
		l := p.Runtime.Pop().(value.Num)
		l.Number = l.Number.Copy()
		f(&l, &r) // Answer is always held in the left value
		p.Runtime.Push(l)
	}

	binCmpOp := func(f func(l, r *value.Num) bool) {
		r := p.Runtime.Pop().(value.Num)
		l := p.Runtime.Pop().(value.Num)
		p.Runtime.Push(value.Bool(f(&l, &r)))
	}

	switch in.Code {
	case code.BIN_OP_ADD:
		binNumOp(func(l, r *value.Num) { l.Number.Add(r.Number) })
	case code.BIN_OP_SUB:
		binNumOp(func(l, r *value.Num) { l.Number.Sub(r.Number) })
	case code.BIN_OP_MUL:
		binNumOp(func(l, r *value.Num) { l.Number.Mul(r.Number) })
	case code.BIN_OP_DIV:
		binNumOp(func(l, r *value.Num) { l.Number.Div(r.Number) })
	case code.BIN_OP_REM:
		binNumOp(func(l, r *value.Num) { l.Number.Mod(r.Number) })

	case code.BIN_OP_AND:
		l, r := p.Runtime.Pop().(value.Bool), p.Runtime.Pop().(value.Bool)
		p.Runtime.Push(l && r)
	case code.BIN_OP_OR:
		l, r := p.Runtime.Pop().(value.Bool), p.Runtime.Pop().(value.Bool)
		p.Runtime.Push(l || r)

	case code.BIN_OP_LESS:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.Less(r.Number) })
	case code.BIN_OP_MORE:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.More(r.Number) })
	case code.BIN_OP_LEQU:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.LessOrEqual(r.Number) })
	case code.BIN_OP_MEQU:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.MoreOrEqual(r.Number) })

	case code.BIN_OP_EQU:
		r, l := p.Runtime.Pop(), p.Runtime.Pop()
		p.Runtime.Push(value.Bool(l.Equal(r)))
	case code.BIN_OP_NEQU:
		r, l := p.Runtime.Pop(), p.Runtime.Pop()
		p.Runtime.Push(value.Bool(!l.Equal(r)))

	default:
		return false
	}

	return true
}
