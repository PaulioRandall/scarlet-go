package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type (
	scanPosErr struct {
		error
		token.UTF8Pos
	}

	scanSnipErr struct {
		error
		token.Snippet
	}
)

func errPos(pos token.UTF8Pos, msg string, args ...interface{}) error {
	return scanPosErr{
		error:   fmt.Errorf(msg, args...),
		UTF8Pos: pos,
	}
}

func errSnip(snip token.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		error:   fmt.Errorf(msg, args...),
		Snippet: snip,
	}
}
