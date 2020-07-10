package check

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type CheckFunc func() (token.Token, CheckFunc, error)

type TokenStream interface {
	Next() token.Token
}

func New(ts TokenStream) CheckFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	chk := &checker{
		Queue: token.Queue{},
		ts:    ts,
	}
	chk.bufferNext()

	if chk.empty() {
		return nil
	}

	return chk.check
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

func CheckAll(tks []token.Token) ([]token.Token, error) {
	return StreamAll(token.NewStream(tks))
}
