package program

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/f_runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
)

func run(args []string) error {

	ins, e := build(args)
	if e != nil {
		const BUILD_ERROR_CODE = 1
		err(BUILD_ERROR_CODE, e)
	}

	exitCode, e := exeInstructions(ins)
	if e != nil {
		err(exitCode, e)
	}

	os.Exit(exitCode)
	return nil
}

func exeInstructions(ins []inst.Instruction) (int, error) {
	rt := runtime.New(ins)
	rt.Start()
	return rt.Env().ExitCode, rt.Env().Err
}

func err(exitCode int, e error) {
	fmt.Println()
	fmt.Printf("[ERROR] %d\n", exitCode)
	fmt.Printf("%+v\n", e)
	os.Exit(exitCode)
}
