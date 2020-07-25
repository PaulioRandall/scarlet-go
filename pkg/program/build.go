package program

import (
	"fmt"
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
)

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

func buildFromConfig(c config) (*inst.Instruction, error) {

	s, e := ioutil.ReadFile(c.script)
	if e != nil {
		return nil, NewGenErr(e)
	}

	head, e := scanAll(c, string(s))
	if e != nil {
		return nil, e
	}

	e = formatAll(c, head)
	if e != nil {
		return nil, e
	}

	head, e = sanitiseAll(c, head)
	if e != nil {
		return nil, e
	}

	head, e = checkAll(c, head)
	if e != nil {
		return nil, e
	}

	head, e = shuntAll(c, head)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(c, head)
	if e != nil {
		return nil, e
	}

	return ins, nil
}
