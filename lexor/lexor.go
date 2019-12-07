package lexor

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ScanToken returns the next token from a stream of source code followed by the
// callable (tail) function to get the token after next. If the function is
// null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken, perror.Perror)
