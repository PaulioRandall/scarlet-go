package sanitise

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
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
	san.bufferNext()

	if san.empty() {
		return nil
	}

	return san.next
}

func StreamAll(ts TokenStream) ([]token.Token, error) {

	var (
		e   error
		tk  token.Token
		tks = []token.Token{}
	)

	for f := New(ts); f != nil; {
		if tk, f, e = f(); e != nil {
			return nil, e
		}

		tks = append(tks, tk)
	}

	return tks, nil
}

func SanitiseAll(tks []token.Token) ([]token.Token, error) {
	return StreamAll(token.NewStream(tks))
}
