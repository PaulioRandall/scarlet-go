package group

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/stat"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type GroupFunc func() (grp, GroupFunc, error)

type grp struct {
	st  StatType
	tks []Token
}

type TokenStream interface {
	Next() Token
}

func New(ts TokenStream) GroupFunc {

	if ts == nil {
		perror.ProgPanic("Non-nil TokenStream required")
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

	e := nextGroup(clt, &gp)
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

		if clt.notMatch(GE_WHITESPACE) {
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
		perror.ProgPanic(
			"No tokens remain, should call `match`, `hasNext`, or `empty` first")
	}

	r := clt.buff
	clt.bufferNext()

	return r
}

func (clt *collector) match(ge GenType) bool {
	return clt.hasNext() && clt.buff.GenType() == ge
}

func (clt *collector) notMatch(ge GenType) bool {
	return !clt.match(ge)
}

func (clt *collector) accept(ge GenType) bool {

	if clt.match(ge) {
		clt.next()
		return true
	}

	return false
}

func (clt *collector) expect(ge GenType) (Token, error) {

	if clt.match(ge) {
		return clt.next(), nil
	}

	var exp string
	if clt.buff == nil {
		exp = "EOF"
	} else {
		exp = clt.buff.GenType().String()
	}

	msg := fmt.Sprintf("Expected %s, got %s", ge.String(), exp)
	return nil, perror.NewBySnippet(msg, clt.buff)
}
