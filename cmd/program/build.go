package program

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/todo/series"

	"github.com/PaulioRandall/scarlet-go/scarlet/compiler"
	"github.com/PaulioRandall/scarlet-go/scarlet/inst"
	"github.com/PaulioRandall/scarlet-go/scarlet/parser"
	"github.com/PaulioRandall/scarlet-go/scarlet/sanitiser"
	"github.com/PaulioRandall/scarlet-go/scarlet/scanner"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

// Build performs a simple workflow that converts a scroll into a set of
// lower level instructions.
func Build(c BuildCmd) ([]inst.Inst, error) {

	b, e := ioutil.ReadFile(c.Scroll)
	if e != nil {
		return nil, e
	}

	src := []rune(string(b))
	tks, e := scanner.ScanAll(src)
	if e != nil {
		return nil, e
	}

	tks = sanitiser.Sanitise(tks)
	s := sliceToSeries(tks)

	trees, e := parser.ParseAll(s)
	if e != nil {
		return nil, e
	}

	insSlice, e := compiler.CompileAll(trees)
	if e != nil {
		return nil, e
	}

	return joinInst(insSlice), nil
}

// joinInst will need to be replaced with more a sohpisticated process once
// functions arrive.
func joinInst(insSlice [][]inst.Inst) []inst.Inst {
	r := []inst.Inst{}
	for _, in := range insSlice {
		r = append(r, in...)
	}
	return r
}

func sliceToSeries(tks []token.Lexeme) *series.Series {
	s := series.New()
	for _, v := range tks {
		s.Append(v)
	}
	s.JumpToStart()
	return s
}
