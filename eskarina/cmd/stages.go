package cmd

import (
	//"io"
	"os"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"

	"github.com/PaulioRandall/scarlet-go/eskarina/stages/a_scanner"
	//"github.com/PaulioRandall/scarlet-go/eskarina/stages/b_format"
	"github.com/PaulioRandall/scarlet-go/eskarina/stages/b_sanitiser"
	"github.com/PaulioRandall/scarlet-go/eskarina/stages/c_checker"
	"github.com/PaulioRandall/scarlet-go/eskarina/stages/d_shunter"
	"github.com/PaulioRandall/scarlet-go/eskarina/stages/e_compiler"
)

func scanAll(c Config, s string) (*lexeme.Container, error) {

	con, e := scanner.ScanStr(s)
	if e != nil {
		return nil, e
	}

	e = logPhase(c, ".scanned", con.Head())
	if e != nil {
		return nil, e
	}

	return con, nil
}

func sanitiseAll(c Config, con *lexeme.Container) (*lexeme.Container, error) {

	con = sanitiser.SanitiseAll(con)

	e := logPhase(c, ".sanitised", con.Head())
	if e != nil {
		return nil, e
	}

	return con, nil
}

func checkAll(c Config, con *lexeme.Container) (*lexeme.Container, error) {

	var e error
	con, e = checker.CheckAll(con)
	if e != nil {
		return nil, e
	}

	return con, nil
}

func shuntAll(c Config, con *lexeme.Container) (*lexeme.Container, error) {

	con = shunter.ShuntAll(con)

	e := logPhase(c, ".shunted", con.Head())
	if e != nil {
		return nil, e
	}

	return con, nil
}

func compileAll(c Config, con *lexeme.Container) (*inst.Instruction, error) {

	ins := compiler.CompileAll(con)

	if !c.log {
		return ins, nil
	}

	f := c.logFilename(".compiled")
	e := writeInstPhaseFile(f, ins)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

/*
func formatAll(c Config, con *lexeme.Container) (con *lexeme.Container, error) {

	if c.nofmt {
		return nil
	}

	head = lexeme.CopyAll(head)
	head = format.FormatAll(head, c.lineEndings)

	f, e := os.Create(c.script)
	if e != nil {
		return e
	}
	defer f.Close()

	return writeLexemes(f, head)
}

func writeLexemes(w io.Writer, head *lexeme.Lexeme) error {

	for lex := head; lex != nil; lex = lex.Next {

		b := []byte(lex.Raw)
		_, e := w.Write(b)

		if e != nil {
			return e
		}
	}

	return nil
}
*/
func logPhase(c Config, ext string, head *lexeme.Lexeme) error {

	if !c.log {
		return nil
	}

	f := c.logFilename(ext)
	return writeLexemeFile(f, head)
}

func writeLexemeFile(filename string, head *lexeme.Lexeme) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return lexeme.PrintAll(f, head)
}

func writeInstPhaseFile(filename string, head *inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.PrintAll(f, head)
}
