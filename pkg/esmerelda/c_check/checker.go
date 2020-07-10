package check

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
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

func (chk *checker) match(ty interface{}) bool {

	if chk.buff == nil {
		return false
	}

	switch v := ty.(type) {
	case GenType:
		return v == chk.buff.GenType()

	case SubType:
		return v == chk.buff.SubType()
	}

	perror.Panic("Invalid kind of token type")
	return false
}

func (chk *checker) accept(ty interface{}) bool {

	if chk.match(ty) {
		chk.bufferNext()
		return true
	}

	return false
}

func (chk *checker) expect(ty interface{}) error {

	if chk.match(ty) {
		chk.bufferNext()
		return nil
	}

	if chk.buff == nil {
		return perror.New("Unexpected EOF")
	}

	var msg string

	switch v := ty.(type) {
	case GenType:
		msg = fmt.Sprintf("Want %q, have %q",
			v.String(), chk.buff.GenType().String())

	case SubType:
		msg = fmt.Sprintf("Want %q, have %q",
			v.String(), chk.buff.SubType().String())
	}

	return perror.NewBySnippet(msg, chk.buff)
}
