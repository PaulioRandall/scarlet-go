package beta

import (
	"github.com/PaulioRandall/scarlet-go/streams/parser/alpha"
)

// statItr represents an iterator of Statements.
type statItr struct {
	stats []alpha.Statement
	size  int
	index int
}

func (itr *statItr) next() (alpha.Statement, bool) {

	if itr.index >= itr.size {
		return alpha.Statement{}, false
	}

	s := itr.stats[itr.index]
	itr.index++
	return s, true
}
