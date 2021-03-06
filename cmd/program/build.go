package program

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/scarlet/parser"
	"github.com/PaulioRandall/scarlet-go/scarlet/scanner"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
)

// Build performs a simple workflow that converts a scroll into a set of
// lower level instructions.
func Build(c BuildCmd) ([]tree.Stat, error) {

	b, e := ioutil.ReadFile(c.Scroll)
	if e != nil {
		return nil, e
	}

	src := []rune(string(b))
	tks, e := scan(src)
	if e != nil {
		return nil, e
	}

	nodes, e := parser.ParseAll(tks)
	if e != nil {
		return nil, e
	}
	return validate(nodes), nil
}

func scan(src []rune) ([]token.Lexeme, error) {

	var (
		r  []token.Lexeme
		l  token.Lexeme
		pt = scanner.New(src)
		e  error
	)

	for pt != nil {
		if l, pt, e = pt(); e != nil {
			return nil, e
		}
		if !l.IsRedundant() {
			r = append(r, l)
		}
	}

	return r, nil
}

// Temp until a validator pkg is created
func validate(nodes []tree.Node) []tree.Stat {
	stmts := make([]tree.Stat, len(nodes))
	for i, n := range nodes {
		s, ok := n.(tree.Stat)
		if !ok {
			panic("Result of expression ignored")
		}
		stmts[i] = s
	}
	return stmts
}
