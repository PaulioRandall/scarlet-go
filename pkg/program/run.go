package program

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/af_runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
)

func printRunHelp() {

	s := `'run' compiles, validates, then runs a script.

Usage:

	scarlet run [options] <script file>

Options:

	-nofmt
		Don't format the script.
	-log 'output folder'
		Logs the output of each compilation stage as labelled files into the
		'output folder'.
`

	fmt.Println(s)
}

func run(ins *inst.Instruction) error {

	rt := runtime.New(ins)
	rt.Start()

	if rt.Env().Err != nil {
		return NewErr(rt.Env().ExitCode, rt.Env().Err)
	}

	if rt.Env().ExitCode != 0 {
		return NewErr(rt.Env().ExitCode, nil)
	}

	return nil
}
