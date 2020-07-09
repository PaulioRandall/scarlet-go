package sanitise

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type SanitiseFunc func() (token.Token, SanitiseFunc, error)

type TokenStream interface {
	Next() token.Token
}

func New(ts TokenStream) SanitiseFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	san := &sanitiser{ts: ts}
	san.buff = san.ts.Next()

	if san.empty() {
		return nil
	}

	return san.next
}

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

	case san.buff.GenType() == GE_WHITESPACE:
		goto BUFFER

	case prev.GenType() == GE_TERMINATOR && san.buff.GenType() == GE_TERMINATOR:
		goto BUFFER

	case prev.SubType() == SU_PAREN_OPEN && san.buff.SubType() == SU_NEWLINE:
		goto BUFFER

	case prev.SubType() == SU_VALUE_DELIM && san.buff.SubType() == SU_NEWLINE:
		goto BUFFER

	case prev.SubType() == SU_VALUE_DELIM && san.buff.SubType() == SU_PAREN_CLOSE:
		prev = san.buff
		goto BUFFER
	}

	return prev
}
