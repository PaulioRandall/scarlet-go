package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type (
	scanPosErr struct {
		error
		token.UTF8Pos
		msg string
	}

	scanSnipErr struct {
		error
		token.Snippet
		msg string
	}
)

func errPos(pos token.UTF8Pos, msg string, args ...interface{}) error {
	return scanPosErr{
		UTF8Pos: pos,
		msg:     fmt.Sprintf(msg, args...),
	}
}

func errSnip(snip token.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     fmt.Sprintf(msg, args...),
	}
}
