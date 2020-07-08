package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
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

	for i := run.env.counter(); i < size; i = run.env.tick() {

		run.exe(run.ins[i])

		if run.halt {
			if run.e != nil {
				return false, run.e
			}

			if run.exitCode >= 0 {
				return true, nil
			}

			return i+1 >= size, nil
		}
	}

	return true, nil
}

func (run Runtime) Stop() {
	run.halt = true
}

func (run Runtime) exe(in Instruction) {
}
