package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token/position"
)

type (
	scanPosErr struct {
		error
		position.UTF8Pos
		msg string
	}

	scanSnipErr struct {
		error
		position.Snippet
		msg string
	}
)

func errPos(pos position.UTF8Pos, msg string, args ...interface{}) error {
	return scanPosErr{
		UTF8Pos: pos,
		msg:     fmt.Sprintf(msg, args...),
	}
}

func errSnip(snip position.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     fmt.Sprintf(msg, args...),
	}
}
