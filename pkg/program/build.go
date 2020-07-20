package program

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/a_scan"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/b_sanitise"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/c_check"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/d_shunt"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/e_compile"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/z_format"
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

func buildFromConfig(c config) ([]inst.Instruction, error) {

	s, e := ioutil.ReadFile(c.script)
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e := scanAll(c, string(s))
	if e != nil {
		return nil, e
	}

	e = formatAll(c, tks)
	if e != nil {
		return nil, e
	}

	tks, e = sanitiseAll(c, tks)
	if e != nil {
		return nil, e
	}

	tks, e = check.CheckAll(tks)
	if e != nil {
		return nil, e
	}

	tks, e = shuntAll(c, tks)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(c, tks)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

func scanAll(c config, s string) ([]token.Token, error) {

	tks, e := scan.ScanAll(s)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(c, ".scanned", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func formatAll(c config, tks []token.Token) error {

	if c.nofmt {
		return nil
	}

	tks = format.FormatAll(tks, c.lineEndings)

	f, e := os.Create(c.script)
	if e != nil {
		return e
	}
	defer f.Close()

	return writeTokens(f, tks)
}

func writeTokens(w io.Writer, tks []token.Token) error {

	for _, tk := range tks {

		b := []byte(tk.Raw())
		_, e := w.Write(b)

		if e != nil {
			return e
		}
	}

	return nil
}

func sanitiseAll(c config, tks []token.Token) ([]token.Token, error) {

	var e error
	tks, e = sanitise.SanitiseAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(c, ".sanitised", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func shuntAll(c config, tks []token.Token) ([]token.Token, error) {

	var e error
	tks, e = shunt.ShuntAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(c, ".shunted", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func compileAll(c config, tks []token.Token) ([]inst.Instruction, error) {

	ins, e := compile.CompileAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	if !c.log {
		return ins, nil
	}

	f := c.logFilename(".compiled")
	e = writeInstPhaseFile(f, ins)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return ins, nil
}

func logPhase(c config, ext string, tks []token.Token) error {

	if !c.log {
		return nil
	}

	f := c.logFilename(ext)
	return writeTokenPhaseFile(f, tks)
}

func writeTokenPhaseFile(filename string, tks []token.Token) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return token.PrintAll(f, tks)
}

func writeInstPhaseFile(filename string, ins []inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.PrintAll(f, ins)
}
