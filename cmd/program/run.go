package program

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/cmd/runtime"
	"github.com/PaulioRandall/scarlet-go/scarlet/processor"
)

// ExeResult represents the result of executing a program within a Processor.
type ExeResult interface {
	error
	Ok() bool
	ExitCode() int
}

type exeResult struct {
	err      error
	exitCode int
}

// Error returns the error message, only call if a call to Ok returns false.
func (e exeResult) Error() string {
	return e.err.Error()
}

// Ok returns true if no error was encountered during execution.
func (e exeResult) Ok() bool {
	return e.err == nil
}

// ExitCode returns the program exit code or an error code if the processor
// failed.
func (e exeResult) ExitCode() int {
	return e.exitCode
}

// Run performs the 'Build' workflow then executes the resultant instruction
// list.
func Run(c RunCmd) (ExeResult, error) {

	program, e := Build(c.BuildCmd)
	if e != nil {
		return nil, e
	}

	env := runtime.New(program)
	p := processor.New(env, env)
	p.Run()

	printEnv(env) // Temp

	return exeResult{err: env.GetErr(), exitCode: env.GetExitCode()}, nil
}

// Temp
func printEnv(env *runtime.RuntimeEnv) {
	fmt.Println("\nVariables created:")
	for id, v := range env.Scope {
		fmt.Println("\t" + id.String() + " " + v.String())
	}
}
