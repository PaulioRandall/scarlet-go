package beta

import (
	"github.com/PaulioRandall/scarlet-go/streams/parser/alpha"
)

// statItr represents an iterator of Statements.
type statItr struct {
	stats []alpha.AlphaStatement
	size  int
	index int
}

func (itr *statItr) next() (alpha.AlphaStatement, bool) {

	if itr.index >= itr.size {
		return alpha.AlphaStatement{}, false
	}

	s := itr.stats[itr.index]
	itr.index++
	return s, true
}
