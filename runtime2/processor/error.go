package processor

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/position"
)

type scanSnipErr struct {
	position.Snippet
	msg string
}

func (e scanSnipErr) Error() string {
	return e.msg
}

func errSnip(snip position.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     fmt.Sprintf(msg, args...),
	}
}
