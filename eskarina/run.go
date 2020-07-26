package eskarina

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/eskarina/cmd"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
)

const GENERAL_ERROR = 1

func Run(args cmd.Arguments) (int, error) {

	if args.Empty() {
		return GENERAL_ERROR, fmt.Errorf("Missing command!")
	}

	command := args.Shift()

	switch command {
	case "help":
		return cmd.Help(args)

	case "docs":
		return cmd.Docs(args)

	case "build":
		_, code, e := build(args)
		return code, e

	case "run":
		ins, code, e := build(args)
		if e != nil {
			return code, e
		}
		return cmd.Run(ins)
	}

	return GENERAL_ERROR, fmt.Errorf("Unknown command %q", command)
}

func build(args cmd.Arguments) (*inst.Instruction, int, error) {

	c := cmd.Config{}
	e := cmd.CaptureConfig(&c, args)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	ins, e := cmd.Build(c)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	return ins, 0, nil
}
