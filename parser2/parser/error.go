package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/position"
)

type scanSnipErr struct {
	error
	position.Snippet
	msg string
}

func errSnip(snip position.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     fmt.Sprintf(msg, args...),
	}
}
