package processor

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type scanSnipErr struct {
	error
	token.Snippet
	msg string
}

func errSnip(snip token.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		error:   fmt.Errorf(msg, args...),
		Snippet: snip,
	}
}
