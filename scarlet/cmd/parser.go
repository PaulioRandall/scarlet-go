package cmd

import (
	//"io"
	"os"

	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"

	"github.com/PaulioRandall/scarlet-go/parser/a_scanner"
	"github.com/PaulioRandall/scarlet-go/parser/b_sanitiser"
	"github.com/PaulioRandall/scarlet-go/parser/c_checker"
	"github.com/PaulioRandall/scarlet-go/parser/d_shunter"
	"github.com/PaulioRandall/scarlet-go/parser/e_compiler"
)

func scanAll(c config, s string) (*lexeme.Container, error) {

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

func sanitiseAll(c config, con *lexeme.Container) error {

	sanitiser.SanitiseAll(con)

	e := logPhase(c, ".sanitised", con.Head())
	if e != nil {
		return e
	}

	return nil
}

func checkAll(c config, con *lexeme.Container) error {

	e := checker.CheckAll(con)
	if e != nil {
		return e
	}

	return nil
}

func shuntAll(c config, con *lexeme.Container) (*lexeme.Container, error) {

	con = shunter.ShuntAll(con)

	e := logPhase(c, ".shunted", con.Head())
	if e != nil {
		return nil, e
	}

	return con, nil
}

func compileAll(c config, con *lexeme.Container) ([]inst.Instruction, error) {

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
func formatAll(c config, con *lexeme.Container) (con *lexeme.Container, error) {

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
func logPhase(c config, ext string, head *lexeme.Lexeme) error {

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

func writeInstPhaseFile(filename string, ins []inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.PrintAll(f, ins)
}
