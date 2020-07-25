package program

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/aa_scanner"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/ab_sanitiser"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/ac_checker"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/ad_shunter"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/ae_compiler"
	//	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/z_format"
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

	first, e := scanAll(c, string(s))
	if e != nil {
		return nil, e
	}
	/*
		e = formatAll(c, tks)
		if e != nil {
			return nil, e
		}
	*/
	first, e = sanitiseAll(c, first)
	if e != nil {
		return nil, e
	}

	e = checker.CheckAll(first)
	if e != nil {
		return nil, e
	}

	first, e = shuntAll(c, first)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(c, first)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

func scanAll(c config, s string) (*lexeme.Lexeme, error) {

	first, e := scanner.ScanStr(s)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(c, ".scanned", first)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return first, nil
}

/*
func formatAll(c config, first *lexeme.Lexeme) error {

	if c.nofmt {
		return nil
	}

	first = format.FormatAll(first, c.lineEndings)

	f, e := os.Create(c.script)
	if e != nil {
		return e
	}
	defer f.Close()

	return writeTokens(f, tks)
}
*/
func writeTokens(w io.Writer, first *lexeme.Lexeme) error {

	for lex := first; lex != nil; lex = lex.Next {

		b := []byte(lex.Raw)
		_, e := w.Write(b)

		if e != nil {
			return e
		}
	}

	return nil
}

func sanitiseAll(c config, first *lexeme.Lexeme) (*lexeme.Lexeme, error) {

	first = sanitiser.SanitiseAll(first)

	e := logPhase(c, ".sanitised", first)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return first, nil
}

func shuntAll(c config, first *lexeme.Lexeme) (*lexeme.Lexeme, error) {

	first = shunter.ShuntAll(first)

	e := logPhase(c, ".shunted", first)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return first, nil
}

func compileAll(c config, first *lexeme.Lexeme) (*inst.Instruction, error) {

	ins := compiler.CompileAll(first)

	if !c.log {
		return ins, nil
	}

	/*
		f := c.logFilename(".compiled")
		e := writeInstPhaseFile(f, ins)
		if e != nil {
			return nil, NewGenErr(e)
		}
	*/
	return ins, nil
}

func logPhase(c config, ext string, first *lexeme.Lexeme) error {

	if !c.log {
		return nil
	}

	f := c.logFilename(ext)
	return writeLexemeFile(f, first)
}

func writeLexemeFile(filename string, first *lexeme.Lexeme) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return lexeme.PrintAll(f, first)
}

/*
func writeInstPhaseFile(filename string, ins *inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.PrintAll(f, ins)
}
*/
