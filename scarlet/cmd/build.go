package cmd

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/temp"

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

	con2, e := scanner.ScanString(s)
	if e != nil {
		return nil, e
	}

	con := temp.ConvertContainer(con2)

	if c.logDir != "" {
		return con, logContainer(c, con, "scanned")
	}

	return con, nil
}

func sanitiseAll(c config, con *lexeme.Container) error {

	sanitiser.SanitiseAll(con)
	if c.logDir != "" {
		return logContainer(c, con, "sanitised")
	}

	return nil
}

func checkAll(c config, con *lexeme.Container) error {
	return checker.CheckAll(con)
}

func shuntAll(c config, con *lexeme.Container) (*lexeme.Container, error) {

	con = shunter.ShuntAll(con)
	if c.logDir != "" {
		return con, logContainer(c, con, "shunted")
	}

	return con, nil
}

func compileAll(c config, con *lexeme.Container) ([]inst.Instruction, error) {

	ins := compiler.CompileAll(con)
	if c.logDir != "" {
		return ins, logInstructions(c, ins, "compiled")
	}

	return ins, nil
}
