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
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/ba_format"
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

	e = formatAll(c, head) // Must be done last
	if e != nil {
		return nil, e
	}

	head, e = sanitiseAll(c, head)
	if e != nil {
		return nil, e
	}

	e = checker.CheckAll(head)
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

func scanAll(c config, s string) (*lexeme.Lexeme, error) {

	head, e := scanner.ScanStr(s)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(c, ".scanned", head)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return head, nil
}

func formatAll(c config, head *lexeme.Lexeme) error {

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

func sanitiseAll(c config, head *lexeme.Lexeme) (*lexeme.Lexeme, error) {

	head = sanitiser.SanitiseAll(head)

	e := logPhase(c, ".sanitised", head)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return head, nil
}

func shuntAll(c config, head *lexeme.Lexeme) (*lexeme.Lexeme, error) {

	head = shunter.ShuntAll(head)

	e := logPhase(c, ".shunted", head)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return head, nil
}

func compileAll(c config, head *lexeme.Lexeme) (*inst.Instruction, error) {

	ins := compiler.CompileAll(head)

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
