package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

type Runtime struct {
	ins      []Instruction
	env      *environment
	halt     bool
	exitCode int
	e        error
}

func New(ins []Instruction) Runtime {
	return Runtime{
		ins:      ins,
		env:      newEnv(),
		exitCode: -1,
	}
}

func (run Runtime) ExitCode() int {
	return run.exitCode
}

func (run Runtime) Start() (bool, error) {

	size := len(run.ins)

	for i := run.env.tick(); i < size; i = run.env.tick() {

		run.exe(run.ins[i])

		if run.halt {
			return run.halted(i+1 >= size)
		}
	}

	return true, nil
}

func (run Runtime) Stop() {
	run.halt = true
}

func (run Runtime) halted(hasMore bool) (bool, error) {

	if run.e != nil {
		return false, run.e
	}

	if run.exitCode >= 0 {
		return true, nil
	}

	return hasMore, nil
}

func (run Runtime) exe(in Instruction) {

	switch in.Code() {
	case IN_VAL_PUSH:
		//run.env.

	case IN_CTX_GET:

	case IN_SPELL:

	default:
		run.e = perror.NewBySnippet("", "Unknown instruction code", in)
	}
}
