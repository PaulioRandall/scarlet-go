package cmd

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/formatter"

	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"

	"github.com/PaulioRandall/scarlet-go/parser/a_scanner"
	"github.com/PaulioRandall/scarlet-go/parser/b_sanitiser"
	"github.com/PaulioRandall/scarlet-go/parser/c_checker"
	"github.com/PaulioRandall/scarlet-go/parser/d_shunter"
	"github.com/PaulioRandall/scarlet-go/parser/e_compiler"
)

func build(c config) ([]inst.Instruction, error) {

	s, e := ioutil.ReadFile(c.script)
	if e != nil {
		return nil, e
	}

	con, e := scanAll(c, string(s))
	if e != nil {
		return nil, e
	}

	e = sanitiseAll(c, con)
	if e != nil {
		return nil, e
	}

	e = checkAll(c, con)
	if e != nil {
		return nil, e
	}

	e = formatAll(c)
	if e != nil {
		return nil, e
	}

	con, e = shuntAll(c, con)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(c, con)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

func scanAll(c config, s string) (*lexeme.Container, error) {

	con, e := scanner.ScanStr(s)
	if e != nil {
		return nil, e
	}

	return con, nil
}

func sanitiseAll(c config, con *lexeme.Container) error {
	sanitiser.SanitiseAll(con)
	return nil
}

func checkAll(c config, con *lexeme.Container) error {
	return checker.CheckAll(con)
}

func shuntAll(c config, con *lexeme.Container) (*lexeme.Container, error) {
	con = shunter.ShuntAll(con)
	return con, nil
}

func compileAll(c config, con *lexeme.Container) ([]inst.Instruction, error) {
	ins := compiler.CompileAll(con)
	return ins, nil
}

func formatAll(c config) error {

	if c.nofmt {
		return nil
	}

	return formatter.FormatFile(c.script)
}
