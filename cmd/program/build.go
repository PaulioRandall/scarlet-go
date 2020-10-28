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

	src, e := ioutil.ReadFile(c.Scroll)
	if e != nil {
		return nil, e
	}

	s, e := scanAll(src)
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

func scanAll(src []byte) (*series.Series, error) {

	var (
		in = []rune(string(src))
		s  = series.New()
		l  token.Lexeme
		pt = scanner.New(in)
		e  error
	)

	for pt != nil {
		if l, pt, e = pt(); e != nil {
			return nil, e
		}
		s.Append(l)
	}

	return s, nil
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
