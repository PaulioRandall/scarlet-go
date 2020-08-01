package cmd

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"

	"github.com/PaulioRandall/scarlet-go/eskarina/parser/f_runtime"
)

func Run(ins *inst.Instruction) (int, error) {

	rt := runtime.New(ins)
	rt.Start()

	if rt.Env().Err != nil {
		return rt.Env().ExitCode, rt.Env().Err
	}

	if rt.Env().ExitCode != 0 {
		return rt.Env().ExitCode, nil
	}

	return 0, nil
}
