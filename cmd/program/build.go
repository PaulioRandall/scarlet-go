package program

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/token/inst"

	"github.com/PaulioRandall/scarlet-go/parser/compiler"
	"github.com/PaulioRandall/scarlet-go/parser/parser"
	"github.com/PaulioRandall/scarlet-go/parser/sanitiser"
	"github.com/PaulioRandall/scarlet-go/parser/scanner"
)

// Build performs a simple workflow that converts a scroll into a set of
// lower level instructions.
func Build(c BuildCmd) ([]inst.Inst, error) {

	src, e := ioutil.ReadFile(c.Scroll)
	if e != nil {
		return nil, e
	}

	runes := []rune(string(src))
	s, e := scanner.ScanAll(runes)
	if e != nil {
		return nil, e
	}

	sanitiser.SanitiseAll(s)
	s.JumpToStart()

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
