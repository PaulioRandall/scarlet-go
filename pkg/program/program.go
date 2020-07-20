package program

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
)

type ScarletError struct {
	ExitCode int
	e        error
}

func (se ScarletError) Error() string {
	return fmt.Sprintf("[ERROR] %d\n%s\n", se.ExitCode, se.e.Error())
}

func NewErr(exitCode int, e error) error {
	return ScarletError{
		ExitCode: exitCode,
		e:        e,
	}
}

func NewGenErr(e error) error {
	return ScarletError{
		ExitCode: 1,
		e:        e,
	}
}

func Execute(args Arguments) error {

	if args.empty() {
		e := fmt.Errorf("Missing command!")
		return NewGenErr(e)
	}

	switch cmd := args.take(); cmd {
	case "help":
		return help(args)

	case "docs":
		return docs(args)

	case "build":
		_, e := build(args)
		return e

	case "run":
		ins, e := build(args)
		if e != nil {
			return e
		}
		return run(ins)

	default:
		e := fmt.Errorf("Unknown command %q", cmd)
		return NewGenErr(e)
	}

	return nil
}

func build(args Arguments) ([]inst.Instruction, error) {

	c := config{}
	e := captureConfig(&c, args)
	if e != nil {
		return nil, e
	}

	ins, e := buildFromConfig(c)
	if e != nil {
		return nil, e
	}

	return ins, nil
}
