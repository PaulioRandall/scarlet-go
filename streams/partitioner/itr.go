package partitioner

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// tkItr represents an iterator of tokens.
type tkItr struct {
	tks   []lexeme.Token
	size  int
	index int
}

func (itr *tkItr) peek() lexeme.Token {

	if itr.index >= itr.size {
		return lexeme.Token{}
	}

	s := itr.tks[itr.index]
	return s
}

func (itr *tkItr) next() lexeme.Token {

	s := itr.peek()

	if s != (lexeme.Token{}) {
		itr.index++
	}

	return s
}
