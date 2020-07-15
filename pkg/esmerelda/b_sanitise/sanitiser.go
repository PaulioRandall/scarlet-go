package sanitise

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

type sanitiser struct {
	ts   TokenStream
	buff token.Token
}

func (san *sanitiser) next() (token.Token, SanitiseFunc, error) {

	if san.empty() {
		perror.Panic("No tokens remaining, call `match`, `hasNext`, or `empty` first")
	}

	tk := san.bufferNext()

	if san.empty() {
		return tk, nil, nil
	}

	return tk, san.next, nil
}

func (san *sanitiser) empty() bool {
	return san.buff == nil
}

func (san *sanitiser) bufferNext() token.Token {

	prev := san.buff

BUFFER:
	san.buff = san.ts.Next()

	switch {
	case san.buff == nil:
		return prev

	case san.buff.Is(PR_REDUNDANT):
		goto BUFFER

	case prev == nil && san.buff.Is(PR_TERMINATOR):
		goto BUFFER

	case prev == nil:
		return prev // First buffer only

	case prev.Is(PR_TERMINATOR) && san.buff.Is(PR_TERMINATOR):
		goto BUFFER

	case prev.Is(PR_PARENTHESIS) && prev.Is(PR_OPENER) && san.buff.Is(PR_NEWLINE):
		goto BUFFER

	case prev.Is(PR_SEPARATOR) && san.buff.Is(PR_NEWLINE):
		goto BUFFER

	case prev.Is(PR_SEPARATOR) && san.buff.Is(PR_PARENTHESIS) && san.buff.Is(PR_CLOSER):
		prev = san.buff
		goto BUFFER
	}

	return prev
}
