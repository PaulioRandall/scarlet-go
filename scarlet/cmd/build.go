package cmd

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/temp"

	"github.com/PaulioRandall/scarlet-go/token/container"

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

	con_old := temp.ConvertContainer(con)

	e = checkAll(c, con_old)
	if e != nil {
		return nil, e
	}

	con_old, e = shuntAll(c, con_old)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(c, con_old)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

func scanAll(c config, s string) (*container.Container, error) {

	con, e := scanner.ScanString(s)
	if e != nil {
		return nil, e
	}

	if c.logDir != "" {
		con_old := temp.ConvertContainer(con)
		e := logContainer(c, con_old, "scanned")
		return con, e
	}

	return con, nil
}

func sanitiseAll(c config, con *container.Container) error {

	sanitiser.SanitiseAll(con.Iterator())
	if c.logDir != "" {
		con_old := temp.ConvertContainer(con)
		return logContainer(c, con_old, "sanitised")
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
