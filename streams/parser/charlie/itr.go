package charlie

import (
	"github.com/PaulioRandall/scarlet-go/streams/parser/beta"
)

// statItr represents an iterator of Statements.
type statItr struct {
	stats []beta.BetaStatement
	size  int
	index int
}

func (itr *statItr) next() (beta.BetaStatement, bool) {

	if itr.index >= itr.size {
		return beta.BetaStatement{}, false
	}

	s := itr.stats[itr.index]
	itr.index++
	return s, true
}
