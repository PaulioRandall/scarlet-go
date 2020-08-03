package cmd

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/runtime"
	"github.com/PaulioRandall/scarlet-go/scarlet/docs"
	"github.com/PaulioRandall/scarlet-go/shared/inst"
)

const GENERAL_ERROR = 1

func Run(args Arguments) (int, error) {

	if args.empty() {
		return GENERAL_ERROR, fmt.Errorf("Missing command!")
	}

	command := args.shift()

	switch command {
	case "help":
		return help(args)

	case "docs":
		return docs.Docs(args.shiftDefault(""))

	case "build":
		_, code, e := buildFromArgs(args)
		return code, e

	case "run":
		ins, code, e := buildFromArgs(args)
		if e != nil {
			return code, e
		}
		return run(ins)
	}

	return GENERAL_ERROR, fmt.Errorf("Unknown command %q", command)
}

func run(ins []inst.Instruction) (int, error) {

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

func buildFromArgs(args Arguments) ([]inst.Instruction, int, error) {

	c := config{}
	e := captureConfig(&c, args)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	ins, e := build(c)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	return ins, 0, nil
}
