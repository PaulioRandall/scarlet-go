package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/f_runtime/enviro"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

type Runtime struct {
	ins []inst.Instruction
	env *enviro.Environment
}

func New(ins []inst.Instruction) *Runtime {
	return &Runtime{
		ins: ins,
		env: enviro.New(),
	}
}

func (run *Runtime) Env() *enviro.Environment {
	return run.env
}

func (run *Runtime) Start() (bool, error) {

	if run.env.Err != nil {
		perror.Panic("Runtime previously encountered an error and cannot continue")
	}

	run.env.Halted = false
	size := len(run.ins)

	for i := run.env.Tick(); i < size; i = run.env.Tick() {

		run.env.Exe(run.ins[i])

		if run.env.Halted {
			return run.halted(i+1 >= size)
		}
	}

	run.env.Exit(0)
	return true, nil
}

func (run *Runtime) Stop() {
	run.env.Halted = true
}

func (run *Runtime) halted(done bool) (bool, error) {

	if run.env.Err != nil {
		return false, run.env.Err
	}

	if run.env.Done || done {
		run.env.Exit(0)
		return true, nil
	}

	return false, nil
}
