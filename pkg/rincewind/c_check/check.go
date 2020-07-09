package check

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/queue"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type CheckFunc func() (token.Token, CheckFunc, error)

type TokenStream interface {
	Next() token.Token
}

type checker struct {
	Queue
	ts   TokenStream
	buff token.Token
}

func New(ts TokenStream) CheckFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	chk := &checker{
		Queue: Queue{},
		ts:    ts,
	}
	chk.bufferNext()

	if chk.empty() {
		return nil
	}

	return chk.check
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
