package runtime

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

type Runtime struct {
	ins []Instruction
	env *environment
}

func New(ins []Instruction) Runtime {
	return Runtime{
		ins: ins,
		env: newEnv(),
	}
}

func (run Runtime) ExitCode() int {
	return run.env.exitCode
}

func (run Runtime) Start() (bool, error) {

	if run.env.e != nil {
		perror.Panic("Runtime previously encountered an error and cannot continue")
	}

	run.env.halt = false
	size := len(run.ins)

	for i := run.env.tick(); i < size; i = run.env.tick() {

		run.exe(run.ins[i])

		if run.env.halt {
			return run.halted(i+1 >= size)
		}
	}

	run.env.exitCode = 0
	return true, nil
}

func (run Runtime) Stop() {
	run.env.halt = true
}

func (run Runtime) halted(hasMore bool) (bool, error) {

	if run.env.e != nil {
		return false, run.env.e
	}

	if run.env.exitCode >= 0 {
		return true, nil
	}

	return hasMore, nil
}

func (run Runtime) err(e error) {
	run.env.e, run.env.halt = e, true
}

func (run Runtime) exe(in Instruction) {

	switch in.Code() {
	case IN_VAL_PUSH:
		run.env.push(result{
			ty:  resultTypeOf(in.Data()),
			val: in.Data(),
		})

	case IN_CTX_GET:
		id := in.Data().(string)
		r, ok := run.env.get(id)

		if ok {
			run.env.push(r)
		} else {
			msg := fmt.Sprintf("Undeclared variable %q", id)
			run.err(perror.NewBySnippet("", msg, in))
		}

	case IN_SPELL:
		invokeSpell(run.env, in)

	default:
		run.err(perror.NewBySnippet("", "Unknown instruction code", in))
	}
}
