package program

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/scarlet/parser"
	"github.com/PaulioRandall/scarlet-go/scarlet/sanitiser"
	"github.com/PaulioRandall/scarlet-go/scarlet/scanner"
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
	tks, e := scanner.ScanAll(src)
	if e != nil {
		return nil, e
	}

	tks = sanitiser.Sanitise(tks)
	nodes, e := parser.ParseAll(tks)
	if e != nil {
		return nil, e
	}
	return validate(nodes), nil
}

// Temp until a validator pkg is created
func validate(nodes []tree.Node) []tree.Stat {
	stmts := make([]tree.Stat, len(nodes))
	for i, n := range nodes {
		s, ok := n.(tree.Stat)
		if !ok {
			panic("Result of expression ignored: " + n.Pos().String())
		}
		stmts[i] = s
	}
	return stmts
}
