package refix

import (
	"fmt"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/pipestack"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type RefixFunc func() (Token, RefixFunc, error)

type TokenStream interface {
	Next() Token
}

type pipeAndMatch struct {
	ts TokenStream
}

type refixer struct {
	*PipeStack
}

func New(ts TokenStream) RefixFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	pam := pipeAndMatch{ts}
	rfx := &refixer{
		NewPipeStack(pam, pam),
	}

	if rfx.Empty() {
		return nil
	}

	return rfx.refix
}

func (rfx *refixer) refix() (Token, RefixFunc, error) {

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

func (rfx *refixer) Next() Token {
	return rfx.PipeStack.Next().(Token)
}

func (rfx *refixer) PeekNext() Token {
	return rfx.PipeStack.PeekNext().(Token)
}

func (rfx *refixer) PeekStack() Token {
	return rfx.PipeStack.PeekTop().(Token)
}

func (rfx *refixer) Pop() Token {
	return rfx.PipeStack.Pop().(Token)
}

func (rfx *refixer) AcceptPop(other interface{}) Token {
	return rfx.PipeStack.AcceptPop(other).(Token)
}

func (rfx *refixer) ExpectPop(other interface{}) (Token, error) {

	tk, e := rfx.PipeStack.ExpectPop(other)
	if e != nil {
		return nil, e
	}

	return tk.(Token), nil
}

func (m pipeAndMatch) Next() interface{} {
	return m.ts.Next()
}

func (pipeAndMatch) Match(ifaceTk, ifacePat interface{}) bool {

	if ifaceTk == nil {
		return false
	}

	tk, ok := ifaceTk.(Token)
	if !ok {
		failNow("refixer pipestack contains something other than a Token")
	}

	if pat, ok := ifacePat.(GenType); ok {
		return pat == GE_ANY || pat == tk.GenType()
	}

	if pat, ok := ifacePat.(SubType); ok {
		return pat == SU_ANY || pat == tk.SubType()
	}

	failNow("refixer.Match requires a GenType or SubType as the second argument")
	return false
}

func (m pipeAndMatch) Expect(ifaceTk, ifacePat interface{}) error {

	if ifaceTk == nil {
		return errorUnexpectedEOF(ifaceTk.(Token))
	}

	if ifacePat == nil {
		failNow("GenType or SubType required")
	}

	if m.Match(ifaceTk, ifacePat) {
		return nil
	}

	return errorWrongToken(ifacePat.(fmt.Stringer), ifaceTk.(Token))
}
