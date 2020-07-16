package program

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/a_scan"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/b_sanitise"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/c_check"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/d_shunt"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/e_compile"
)

type buildConfig struct {
	script  string
	nofmt   bool
	log     bool
	logFile string
}

func (bc buildConfig) logFilename(ext string) string {
	f := filepath.Base(bc.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(bc.logFile, f+ext)
}

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

func build(args Arguments) ([]inst.Instruction, error) {

	bc := buildConfig{}
	e := parseBuildArgs(&bc, args)
	if e != nil {
		return nil, e
	}

	ins, e := buildFromConfig(bc)
	return ins, e
}

func parseBuildArgs(bc *buildConfig, args Arguments) error {

	for args.more() {

		if !strings.HasPrefix(args.peek(), "-") {
			break
		}

		e := buildOption(bc, args)
		if e != nil {
			return e
		}
	}

	if args.empty() {
		e := fmt.Errorf("Expected script filename")
		return NewGenErr(e)
	}

	bc.script = args.take()

	if args.more() {
		e := fmt.Errorf("Unexpected argument %q", args.peek())
		return NewGenErr(e)
	}

	return nil
}

func buildOption(bc *buildConfig, args Arguments) error {

	switch args.peek() {
	case "-nofmt":
		nofmtOption(bc, args)

	case "-log":
		return logOption(bc, args)

	default:
		e := fmt.Errorf("Unexpected option %q", args.peek())
		return NewGenErr(e)
	}

	return nil
}

func nofmtOption(bc *buildConfig, args Arguments) {
	bc.nofmt = true
	args.take()
}

func logOption(bc *buildConfig, args Arguments) error {

	if args.count() < 2 {
		e := fmt.Errorf("Missing %q folder name", args.peek())
		return NewGenErr(e)
	}

	bc.log = true
	args.take()
	bc.logFile = args.take()

	return nil
}

func buildFromConfig(bc buildConfig) ([]inst.Instruction, error) {

	s, e := ioutil.ReadFile(bc.script)
	if e != nil {
		return nil, NewGenErr(e)
	}

	tks, e := scanAll(bc, string(s))
	if e != nil {
		return nil, e
	}

	tks, e = sanitiseAll(bc, tks)
	if e != nil {
		return nil, e
	}

	tks, e = check.CheckAll(tks)
	if e != nil {
		return nil, e
	}

	tks, e = shuntAll(bc, tks)
	if e != nil {
		return nil, e
	}

	ins, e := compileAll(bc, tks)
	if e != nil {
		return nil, e
	}

	return ins, nil
}

func scanAll(bc buildConfig, s string) ([]token.Token, error) {

	tks, e := scan.ScanAll(s)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(bc, ".scanned", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func sanitiseAll(bc buildConfig, tks []token.Token) ([]token.Token, error) {

	var e error
	tks, e = sanitise.SanitiseAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(bc, ".sanitised", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func shuntAll(bc buildConfig, tks []token.Token) ([]token.Token, error) {

	var e error
	tks, e = shunt.ShuntAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	e = logPhase(bc, ".shunted", tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return tks, nil
}

func compileAll(bc buildConfig, tks []token.Token) ([]inst.Instruction, error) {

	ins, e := compile.CompileAll(tks)
	if e != nil {
		return nil, NewGenErr(e)
	}

	if !bc.log {
		return ins, nil
	}

	f := bc.logFilename(".compiled")
	e = writeInstPhaseFile(f, ins)
	if e != nil {
		return nil, NewGenErr(e)
	}

	return ins, nil
}

func logPhase(bc buildConfig, ext string, tks []token.Token) error {

	if !bc.log {
		return nil
	}

	f := bc.logFilename(ext)
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
