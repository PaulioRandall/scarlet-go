package parser

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type (
	scanPosErr struct {
		token.UTF8Pos
		msg string
	}

	scanSnipErr struct {
		token.Snippet
		msg string
	}
)

func (e scanPosErr) Error() string {
	return e.msg
}

func (e scanSnipErr) Error() string {
	return e.msg
}

func errPos(pos token.UTF8Pos, msg string, args ...interface{}) error {
	return scanPosErr{
		UTF8Pos: pos,
		msg:     "Parser: " + fmt.Sprintf(msg, args...),
	}
}

func errSnip(snip token.Snippet, msg string, args ...interface{}) error {
	return scanSnipErr{
		Snippet: snip,
		msg:     "Parser: " + fmt.Sprintf(msg, args...),
	}
}
