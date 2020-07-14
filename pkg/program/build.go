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

func build(args []string) ([]inst.Instruction, error) {

	bc := buildConfig{}

	for ; len(args) > 0; args = args[1:] {

		if !strings.HasPrefix(args[0], "-") {
			break
		}

		buildOption(&bc, args[0])
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("Expected script filename")
	}

	if len(args) > 1 {
		return nil, fmt.Errorf("Unexpected argument %q", args[1])
	}

	bc.script = args[0]
	return buildScript(bc)
}

func buildOption(bc *buildConfig, arg string) error {

	switch arg {
	case "-nofmt":
		bc.nofmt = true

	case "-log":
		bc.log = true

	default:
		return fmt.Errorf("Unexpected option %q", arg)
	}

	return nil
}

func buildScript(bc buildConfig) ([]inst.Instruction, error) {

	s, e := ioutil.ReadFile(bc.script)
	if e != nil {
		return nil, e
	}

	tks, e := scan.ScanAll(string(s))
	if e != nil {
		return nil, e
	}

	tks, e = sanitise.SanitiseAll(tks)
	if e != nil {
		return nil, e
	}

	tks, e = check.CheckAll(tks)
	if e != nil {
		return nil, e
	}

	tks, e = shunt.ShuntAll(tks)
	if e != nil {
		return nil, e
	}

	ins, e := compile.CompileAll(tks)
	return ins, e
}
