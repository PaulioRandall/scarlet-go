package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/pipestack"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type refixer struct {
	*PipeStack
}

func (rfx *refixer) refix() (token.Token, RefixFunc, error) {

	if rfx.Empty() {
		failNow("No tokens remain, should call 'empty', 'match', etc first")
	}

	tk, e := next(rfx)
	if e != nil {
		return nil, nil, e
	}

	if rfx.Empty() {
		return tk, nil, nil
	}

	return tk, rfx.refix, nil
}

func (rfx *refixer) Next() token.Token {
	return rfx.PipeStack.Next().(token.Token)
}

func (rfx *refixer) PeekBuff() token.Token {
	return rfx.PipeStack.PeekBuff().(token.Token)
}

func (rfx *refixer) PeekTop() token.Token {
	return rfx.PipeStack.PeekTop().(token.Token)
}

func (rfx *refixer) Pop() token.Token {
	return rfx.PipeStack.Pop().(token.Token)
}

func (rfx *refixer) AcceptPop(other interface{}) token.Token {
	return rfx.PipeStack.AcceptPop(other).(token.Token)
}

func (rfx *refixer) ExpectPop(other interface{}) (token.Token, error) {

	tk, e := rfx.PipeStack.ExpectPop(other)
	if e != nil {
		return nil, e
	}

	return tk.(token.Token), nil
}
