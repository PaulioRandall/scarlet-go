package shunt

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

type RefixFunc func() (token.Token, RefixFunc, error)

type TokenStream interface {
	Next() token.Token
}

func New(ts TokenStream) RefixFunc {

	if ts == nil {
		failNow("Non-nil Token Stream required")
	}

	shy := &shuntingYard{
		Stack: token.Stack{},
		ts:    ts,
	}
	shy.buff = shy.ts.Next()

	if shy.empty() {
		return nil
	}

	return shy.shunt
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

func ShuntAll(tks []token.Token) ([]token.Token, error) {
	return StreamAll(token.NewStream(tks))
}
