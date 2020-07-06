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

type refixer struct {
	ts   TokenStream
	stk  stack
	buff Token
}

func (rfx *refixer) refix() (Token, RefixFunc, error) {

	if rfx.empty() {
		failNow("No tokens remain, should call 'empty' or 'match' first")
	}

	tk, e := next(rfx)
	if e != nil {
		return nil, nil, e
	}

	if rfx.empty() {
		return tk, nil, nil
	}

	return tk, rfx.refix, nil
}

func (rfx *refixer) bufferNext() {
	rfx.buff = rfx.ts.Next()
}

func (rfx *refixer) empty() bool {
	return rfx.buff == nil && rfx.stk.empty()
}

func (rfx *refixer) match(gen GenType, sub SubType) bool {

	if rfx.empty() {
		return false
	}

	g := gen == GE_ANY || gen == rfx.buff.GenType()
	s := sub == SU_ANY || sub == rfx.buff.SubType()

	return g && s
}

func (rfx *refixer) notMatch(gen GenType, sub SubType) bool {
	return !rfx.match(gen, sub)
}

func (rfx *refixer) expect(gen GenType, sub SubType) (Token, error) {

	if rfx.empty() {
		return nil, errorUnexpectedEOF(rfx)
	}

	if rfx.match(gen, sub) {
		return rfx.next(), nil
	}

	return nil, errorWrongToken(rfx, rfx.buff)
}

func (rfx *refixer) next() Token {

	if rfx.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	tk := rfx.buff
	rfx.bufferNext()

	return tk
}
