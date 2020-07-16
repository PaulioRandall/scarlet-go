package check

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

type checker struct {
	token.Queue
	ts   TokenStream
	buff token.Token
}

func (chk *checker) check() (token.Token, CheckFunc, error) {

	if chk.Queue.Empty() {
		if e := next(chk); e != nil {
			return nil, nil, e
		}
	}

	tk := chk.Take()
	if chk.empty() {
		return tk, nil, nil
	}

	return tk, chk.check, nil
}

func (chk *checker) Put(tk token.Token) {
	chk.Queue.Put(tk)
}

func (chk *checker) Take() token.Token {
	return chk.Queue.Take().(token.Token)
}

func (chk *checker) bufferNext() {

	if chk.buff != nil {
		chk.Put(chk.buff)
	}

	chk.buff = chk.ts.Next()
}

func (chk *checker) empty() bool {
	return chk.buff == nil && chk.Queue.Empty()
}

func (chk *checker) match(props ...Prop) bool {

	if chk.buff == nil {
		return false
	}

	for _, p := range props {
		if chk.buff.IsNot(p) {
			return false
		}
	}

	return true
}

func (chk *checker) accept(props ...Prop) bool {

	if chk.match(props...) {
		chk.bufferNext()
		return true
	}

	return false
}

func (chk *checker) expect(props ...Prop) error {

	if chk.match(props...) {
		chk.bufferNext()
		return nil
	}

	if chk.buff == nil {
		return perror.New("Unexpected EOF")
	}

	msg := fmt.Sprintf(
		"Want %q, have %q",
		JoinProps(" & ", props...),
		JoinProps(" & ", chk.buff.Props()...),
	)
	return perror.NewBySnippet(msg, chk.buff)
}
