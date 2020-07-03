package group

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type GroupFunc func() (grp, GroupFunc, error)

type TokenStream interface {
	Next() Token
}

func New(ts TokenStream) GroupFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	clt := &collector{ts: ts}
	clt.bufferNext()

	if clt.empty() {
		return nil
	}

	return clt.group
}

type collector struct {
	ts   TokenStream
	buff Token
}

func (clt *collector) group() (grp, GroupFunc, error) {

	gp := grp{}

	e := group(clt, &gp)
	if e != nil {
		return grp{}, nil, e
	}

	if clt.empty() {
		return gp, nil, nil
	}

	return gp, clt.group, nil
}

func (clt *collector) bufferNext() {

	for {
		clt.buff = clt.ts.Next()

		if clt.buff == nil {
			return
		}

		if clt.notMatchGen(GE_WHITESPACE) {
			break
		}
	}
}

func (clt *collector) hasNext() bool {
	return clt.buff != nil
}

func (clt *collector) empty() bool {
	return clt.buff == nil
}

func (clt *collector) next() Token {

	if clt.empty() {
		failNow("No tokens remaining, call `match`, `hasNext`, or `empty` first")
	}

	r := clt.buff
	clt.bufferNext()

	return r
}

func (clt *collector) matchGen(ge GenType) bool {
	return clt.hasNext() && clt.buff.GenType() == ge
}

func (clt *collector) notMatchGen(ge GenType) bool {
	return !clt.matchGen(ge)
}

func (clt *collector) expectGen(ge GenType) (Token, error) {

	if clt.matchGen(ge) {
		return clt.next(), nil
	}

	if clt.buff == nil {
		return nil, errorUnexpectedEOF(clt)
	}

	return nil, errorWrongToken(clt, clt.buff)
}

func (clt *collector) matchSub(su SubType) bool {
	return clt.hasNext() && clt.buff.SubType() == su
}

func (clt *collector) notMatchSub(su SubType) bool {
	return !clt.matchSub(su)
}

func (clt *collector) expectSub(su SubType) (Token, error) {

	if clt.matchSub(su) {
		return clt.next(), nil
	}

	if clt.buff == nil {
		return nil, errorUnexpectedEOF(clt)
	}

	return nil, errorWrongToken(clt, clt.buff)
}

func (clt *collector) acceptAppendGen(gp *grp, ge GenType) bool {

	if clt.matchGen(ge) {
		gp.tks = append(gp.tks, clt.next())
		return true
	}

	return false
}

func (clt *collector) expectAppendGen(gp *grp, ge GenType) error {

	tk, e := clt.expectGen(ge)
	if e != nil {
		return e
	}

	gp.tks = append(gp.tks, tk)
	return nil
}

func (clt *collector) acceptAppendSub(gp *grp, su SubType) bool {

	if clt.matchSub(su) {
		gp.tks = append(gp.tks, clt.next())
		return true
	}

	return false
}

func (clt *collector) expectAppendSub(gp *grp, su SubType) error {

	tk, e := clt.expectSub(su)
	if e != nil {
		return e
	}

	gp.tks = append(gp.tks, tk)
	return nil
}
