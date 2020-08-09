package cmd

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/runtime"
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

	case "docs", "man":
		return docs(args)

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
	e := c.captureConfig(args)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	ins, e := build(c)
	if e != nil {
		return nil, GENERAL_ERROR, e
	}

	return ins, 0, nil
}

func docs(args Arguments) (int, error) {

	searchTerm := args.shiftDefault("")
	text, found := manual.Search(searchTerm)

	if !found {
		return 1, fmt.Errorf("No documentation for %q", searchTerm)
	}

	fmt.Println(text)
	return 0, nil
}
