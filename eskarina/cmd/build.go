package cmd

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
)

func Build(c Config) (*inst.Instruction, error) {

	s, e := ioutil.ReadFile(c.script)
	if e != nil {
		return nil, e
	}

	head, e := scanAll(c, string(s))
	if e != nil {
		return nil, e
	}
	/*
		e = formatAll(c, head)
		if e != nil {
			return nil, e
		}
	*/
	e = sanitiseAll(c, head)
	if e != nil {
		return nil, e
	}

	e = checkAll(c, head)
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
