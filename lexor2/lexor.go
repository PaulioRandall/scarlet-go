package lexor

import (
	"github.com/PaulioRandall/scarlet-go/token2"
)

// ScanToken returns the next token from a stream of code followed by the next
// ScanToken function that will return the next token. If the function is
// null then the end of the token stream has been reached.
type ScanToken func() (token.Token, ScanToken, ScanErr)
