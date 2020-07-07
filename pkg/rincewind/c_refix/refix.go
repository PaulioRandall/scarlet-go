package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type RefixFunc func() (Token, RefixFunc, error)

type TokenStream interface {
	Next() Token
}

func New(ts TokenStream) RefixFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	rfx := &refixer{ts: ts}
	rfx.bufferNext()

	if rfx.empty() {
		return nil
	}

	return rfx.refix
}
