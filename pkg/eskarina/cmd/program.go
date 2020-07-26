package cmd

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
)

const GENERAL_ERROR = 1

func Execute(args Arguments) (int, error) {

	if args.empty() {
		return GENERAL_ERROR, fmt.Errorf("Missing command!")
	}

	cmd := args.take()

	switch cmd {
	case "help":
		return help(args)

	case "docs":
		return docs(args)

	case "build":
		_, code, e := build(args)
		return code, e

	case "run":
		ins, code, e := build(args)
		if e != nil {
			return code, e
		}
		return run(ins)
	}

	return GENERAL_ERROR, fmt.Errorf("Unknown command %q", cmd)
}

func build(args Arguments) (*inst.Instruction, int, error) {

	c := config{}
	e := captureConfig(&c, args)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	ins, e := buildFromConfig(c)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	return ins, 0, nil
}
