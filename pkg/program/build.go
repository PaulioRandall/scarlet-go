package program

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/a_scan"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/b_sanitise"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/c_check"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/d_shunt"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/e_compile"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
)

type buildConfig struct {
	script string
	nofmt  bool
	log    bool
}

func printBuildHelp() {

	s := `'build' compiles and validates a script.

Usage:

	scarlet build [options] <script file>

Options:

	-nofmt
		Don't format the script.
	-log <output folder>
		Logs the output of each compilation stage as labelled files into the
		output folder.
`

	fmt.Println(s)
}

func parseBuildArgs(args []string) (buildConfig, error) {

	bc := buildConfig{}

	for ; len(args) > 0; args = args[1:] {

		if !strings.HasPrefix(args[0], "-") {
			break
		}

		e := buildOption(&bc, args[0])
		if e != nil {
			return buildConfig{}, e
		}
	}

	if len(args) == 0 {
		e := fmt.Errorf("Expected script filename")
		return buildConfig{}, NewGenErr(e)
	}

	if len(args) > 1 {
		e := fmt.Errorf("Unexpected argument %q", args[1])
		return buildConfig{}, NewGenErr(e)
	}

	bc.script = args[0]
	return bc, nil
}

func buildOption(bc *buildConfig, arg string) error {

	switch arg {
	case "-nofmt":
		bc.nofmt = true

	case "-log":
		bc.log = true

	default:
		e := fmt.Errorf("Unexpected option %q", arg)
		return NewGenErr(e)
	}

	return nil
}

func buildFromConfig(bc buildConfig) ([]inst.Instruction, error) {

	s, e := ioutil.ReadFile(bc.script)
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e := scan.ScanAll(string(s))
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e = sanitise.SanitiseAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e = check.CheckAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e = shunt.ShuntAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	ins, e := compile.CompileAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return ins, nil
}

func build(args []string) ([]inst.Instruction, error) {

	bc, e := parseBuildArgs(args)
	if e != nil {
		return nil, e
	}

	ins, e := buildFromConfig(bc)
	return ins, e
}
