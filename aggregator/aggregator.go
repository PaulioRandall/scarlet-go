package aggregator

import (
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/lexor/evaluator"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/PaulioRandall/scarlet-go/stat"
)

// GroupTokens returns the next grouping of tokens (statement) from a stream
// of tokens followed by the callable (tail) function to get the group of tokens
// after next. If the function is null then the end of the token stream has
// been reached.
type GroupTokens func() (stat.Statement, GroupTokens, token.Perror)

// New returns a function to group tokens
func New(src string) GroupTokens {
	st := evaluator.New(src)
	return group(st)
}

// group returns a GroupTokens function that will return the next statement in
// the token stream.
func group(st lexor.ScanToken) GroupTokens {

	if st == nil {
		return nil
	}

	return func() (stat.Statement, GroupTokens, token.Perror) {
		return nil, nil, nil
	}
}
