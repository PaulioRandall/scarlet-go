package check

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type CheckFunc func() (Token, CheckFunc, error)

type TokenStream interface {
	Next() Token
}

type checker struct {
	queue
	ts   TokenStream
	buff Token
}

func New(ts TokenStream) CheckFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	chk := &checker{
		queue: queue{},
		ts:    ts,
	}
	chk.bufferNext()

	if chk.empty() {
		return nil
	}

	return chk.check
}

func (chk *checker) check() (Token, CheckFunc, error) {

	if chk.queue.empty() {
		if e := next(chk); e != nil {
			return nil, nil, e
		}
	}

	tk := chk.queue.take()
	if chk.empty() {
		return tk, nil, nil
	}

	return tk, chk.check, nil
}

func (chk *checker) bufferNext() {

	if chk.buff != nil {
		chk.queue.put(chk.buff)
	}

	chk.buff = chk.ts.Next()
}

func (chk *checker) empty() bool {
	return chk.buff == nil && chk.queue.empty()
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

const ERR_WRONG_TOKEN string = "CHECK_ERR_WRONG_TOKEN"

func (chk *checker) expect(ty interface{}) error {

	if chk.match(ty) {
		chk.bufferNext()
		return nil
	}

	if chk.buff == nil {
		return perror.New(ERR_WRONG_TOKEN, "Unexpected EOF")
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

	return perror.NewBySnippet(ERR_WRONG_TOKEN, msg, chk.buff)
}
