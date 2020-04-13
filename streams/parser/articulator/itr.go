package articulator

import (
	"github.com/PaulioRandall/scarlet-go/streams/parser/partitioner"
)

// statItr represents an iterator of Statements.
type statItr struct {
	stats []partitioner.Statement
	size  int
	index int
}

func (itr *statItr) next() (partitioner.Statement, bool) {

	if itr.index >= itr.size {
		return partitioner.Statement{}, false
	}

	s := itr.stats[itr.index]
	itr.index++
	return s, true
}
