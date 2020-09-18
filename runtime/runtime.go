package runtime

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/runtime/enviro"
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

func (run *Runtime) Start() {

	if run.env.Err != nil {
		panic("Runtime previously encountered an error and cannot continue")
	}

	run.env.Halted = false
	size := len(run.ins)

	for i := run.env.Tick(); i < size; i = run.env.Tick() {

		run.env.Exe(run.ins[i])

		if run.env.Halted {
			run.halted(i+1 >= size)
			return
		}
	}

	run.halted(true)
}

func (run *Runtime) Stop() {
	run.env.Halted = true
}

func (run *Runtime) halted(done bool) {

	if run.env.Err != nil {
		if run.env.ExitCode < 1 {
			run.env.ExitCode = 1
		}
		return
	}

	if run.env.ExitCode > 0 {
		return
	}

	if run.env.Done || done {
		run.env.Exit(0)
	}
}
