package aggregator

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/stat"
)

// GroupTokens returns the next grouping of tokens (statement) from a stream
// of tokens followed by the callable (tail) function to get the group of tokens
// after next. If the function is null then the end of the token stream has
// been reached.
type GroupTokens func() (stat.Statement, GroupTokens, perror.Perror)
