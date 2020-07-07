package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type stack struct {
	top  *node
	size int
}

type node struct {
	tk   Token
	next *node
}

type refixer struct {
	ts   TokenStream
	stk  stack
	buff Token
}

func (rfx *refixer) push(tk Token) {

	rfx.stk.top = &node{
		tk:   tk,
		next: rfx.stk.top,
	}

	rfx.stk.size++
}

func (rfx *refixer) peek() Token {

	if rfx.stk.size == 0 {
		return nil
	}

	return rfx.stk.top.tk
}

func (rfx *refixer) pop() Token {

	if rfx.stk.size == 0 {
		failNow("Nothing to pop")
	}

	tk := rfx.stk.top.tk
	rfx.stk.top = rfx.stk.top.next
	rfx.stk.size--

	return tk
}

func (rfx *refixer) bufferNext() {
	rfx.buff = rfx.ts.Next()
}

func (rfx *refixer) emptyStk() bool {
	return rfx.stk.size == 0
}

func (rfx *refixer) empty() bool {
	return rfx.buff == nil && rfx.stk.size == 0
}

func match(tk Token, gen GenType, sub SubType) bool {

	if tk == nil {
		return false
	}

	g := gen == GE_ANY || gen == tk.GenType()
	s := sub == SU_ANY || sub == tk.SubType()

	return g && s
}

func (rfx *refixer) match(gen GenType, sub SubType) bool {
	return match(rfx.buff, gen, sub)
}

func (rfx *refixer) notMatch(gen GenType, sub SubType) bool {
	return !match(rfx.buff, gen, sub)
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

func (rfx *refixer) pushNext() {

	if rfx.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	rfx.push(rfx.next())
}

func (rfx *refixer) matchPush(gen GenType, sub SubType) bool {

	if rfx.match(gen, sub) {
		rfx.push(rfx.next())
		return true
	}

	return false
}

func (rfx *refixer) expectPush(gen GenType, sub SubType) error {

	tk, e := rfx.expect(gen, sub)
	if e != nil {
		return e
	}

	rfx.push(tk)
	return nil
}

func (rfx *refixer) matchStk(gen GenType, sub SubType) bool {
	return match(rfx.peek(), gen, sub)
}

func (rfx *refixer) matchPop(gen GenType, sub SubType) Token {

	if match(rfx.peek(), gen, sub) {
		return rfx.pop()
	}

	return nil
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
