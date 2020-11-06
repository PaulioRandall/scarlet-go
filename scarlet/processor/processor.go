package processor

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

// Runtime is a handler for performing memory related and context dependent
// instructions such as access to the value stack and acces to scope variables.
type Runtime interface {

	// Push a value onto the top of the value stack.
	Push(value.Value)

	// Pop a value off the top of the value stack,
	Pop() value.Value

	// Spellbook returns the book containing the spells available during runtime.
	Spellbook() spell.Book

	// Bind sets the value of a variable overwriting any existing value.
	Bind(value.Ident, value.Value)

	// Unbind.
	Unbind(value.Ident)

	// Fetch returns the value associated with the specified identifier.
	Fetch(value.Ident) value.Value

	// Fetch the value associated with the specified identifier and pushes onto
	// the value stack.
	FetchPush(value.Ident)

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

// Program provides access to the program instructions.
type Program interface {

	// Next returns the next instruction inicated by the program counter. True is
	// returned if an instruction was returned otherwise the end of program has
	// been reached.
	Next() (inst.Inst, bool)
}

// Processor executes instructions in a similar fashion to a CPU but at a
// higher level.
type Processor struct {
	Program Program // Access to instructions and the value stack
	Env     Runtime // Access to memory and context dependent behaviour
	Halt    bool    // True to interupt execution after the next instruction
	Halted  bool    // True if execution was stopped by an interupt
}

// New returns a new Processor with the specified Program and Runtime.
func New(p Program, env Runtime) *Processor {
	return &Processor{
		Program: p,
		Env:     env,
	}
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

	if p.Halt {
		p.Halted = true
		return
	}

	p.Halted = false

	for !p.Env.GetExitFlag() {

		if p.Halt {
			p.Halt = false
			p.Halted = true
			return // Processor was directly interupted
		}

		if in, ok = p.Program.Next(); !ok {
			p.Env.Exit(0)
			return // No more instructions to execute
		}

		p.Process(in)
	}
}

// Process the instruction 'in' using the memory 'm'. 'halt' should only be
// returned as true if an instruction specifically requests execution to halt.
func (p *Processor) Process(in inst.Inst) {
	switch {
	case in.Code == inst.FETCH_PUSH:
		p.Env.FetchPush(in.Data.(value.Ident))

	case in.Code == inst.STACK_PUSH:
		p.Env.Push(in.Data)

	case in.Code == inst.SCOPE_BIND:
		v := p.Env.Pop()
		if v == nil {
			return // Temp
		}
		p.Env.Bind(in.Data.(value.Ident), v)

	case in.Code == inst.SPELL_CALL:
		spellCall(p, in)

	case processNumOp(p, in):

	default:
		panic("Unhandled instruction code: " + in.Code.String())
	}
}

func processNumOp(p *Processor, in inst.Inst) bool {

	binNumOp := func(f func(l, r *value.Num)) {
		r := p.Env.Pop().(value.Num)
		l := p.Env.Pop().(value.Num)
		l.Number = l.Number.Copy()
		f(&l, &r) // Answer is always held in the left value
		p.Env.Push(l)
	}

	binCmpOp := func(f func(l, r *value.Num) bool) {
		r := p.Env.Pop().(value.Num)
		l := p.Env.Pop().(value.Num)
		p.Env.Push(value.Bool(f(&l, &r)))
	}

	switch in.Code {
	case inst.BIN_OP_ADD:
		binNumOp(func(l, r *value.Num) { l.Number.Add(r.Number) })
	case inst.BIN_OP_SUB:
		binNumOp(func(l, r *value.Num) { l.Number.Sub(r.Number) })
	case inst.BIN_OP_MUL:
		binNumOp(func(l, r *value.Num) { l.Number.Mul(r.Number) })
	case inst.BIN_OP_DIV:
		binNumOp(func(l, r *value.Num) { l.Number.Div(r.Number) })
	case inst.BIN_OP_REM:
		binNumOp(func(l, r *value.Num) { l.Number.Mod(r.Number) })

	case inst.BIN_OP_AND:
		l, r := p.Env.Pop().(value.Bool), p.Env.Pop().(value.Bool)
		p.Env.Push(l && r)
	case inst.BIN_OP_OR:
		l, r := p.Env.Pop().(value.Bool), p.Env.Pop().(value.Bool)
		p.Env.Push(l || r)

	case inst.BIN_OP_LESS:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.Less(r.Number) })
	case inst.BIN_OP_MORE:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.More(r.Number) })
	case inst.BIN_OP_LEQU:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.LessOrEqual(r.Number) })
	case inst.BIN_OP_MEQU:
		binCmpOp(func(l, r *value.Num) bool { return l.Number.MoreOrEqual(r.Number) })

	case inst.BIN_OP_EQU:
		r, l := p.Env.Pop(), p.Env.Pop()
		p.Env.Push(value.Bool(l.Equal(r)))
	case inst.BIN_OP_NEQU:
		r, l := p.Env.Pop(), p.Env.Pop()
		p.Env.Push(value.Bool(!l.Equal(r)))

	default:
		return false
	}

	return true
}

func spellCall(p *Processor, in inst.Inst) {

	name := in.Data.(value.Ident)
	sp, ok := p.Env.Spellbook().Lookup(string(name))

	if !ok {
		panic("Unknown spell: " + name)
	}

	args := []value.Value{}
	for a := p.Env.Pop(); a != nil; a = p.Env.Pop() {
		args = append(args, a)
	}

	out := spell.NewOutput(sp.Outputs)
	sp.Spell(p.Env, args, out)

	for _, v := range out.Slice() {
		p.Env.Push(v)
	}
}
