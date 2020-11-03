package processor2

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type scanSnipErr struct {
	token.Snippet
	msg string
}

func (e scanSnipErr) Error() string {
	return e.msg
}

func errSnip(snip token.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     fmt.Sprintf(msg, args...),
	}
}
