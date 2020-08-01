package cmd

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/runtime"
	"github.com/PaulioRandall/scarlet-go/shared/inst"
)

const GENERAL_ERROR = 1

func Run(args Arguments) (int, error) {

	if args.Empty() {
		return GENERAL_ERROR, fmt.Errorf("Missing command!")
	}

	command := args.Shift()

	switch command {
	case "help":
		return Help(args)

	case "docs":
		return Docs(args)

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

	return GENERAL_ERROR, fmt.Errorf("Unknown command %q", command)
}

func run(ins *inst.Instruction) (int, error) {

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

func build(args Arguments) (*inst.Instruction, int, error) {

	c := Config{}
	e := CaptureConfig(&c, args)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	ins, e := Build(c)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	return ins, 0, nil
}
