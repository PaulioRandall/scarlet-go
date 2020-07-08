package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/queue"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type CompileFunc func() (instruction, CompileFunc, error)

type TokenStream interface {
	Next() Token
}

type compiler struct {
	Queue
	ts   TokenStream
	buff Token
}

func New(ts TokenStream) CompileFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	com := &compiler{
		Queue: Queue{},
		ts:    ts,
	}
	com.buff = com.ts.Next()

	if com.empty() {
		return nil
	}

	return com.compile
}

func (com *compiler) compile() (instruction, CompileFunc, error) {

	if com.Queue.Empty() {
		if e := next(com); e != nil {
			return instruction{}, nil, e
		}
	}

	in := com.Take().(instruction)
	if com.empty() {
		return in, nil, nil
	}

	return in, com.compile, nil
}

func (com *compiler) empty() bool {
	return com.buff == nil && com.Queue.Empty()
}

func (com *compiler) next() Token {

	if com.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	r := com.buff
	com.buff = com.ts.Next()

	return r
}

func (com *compiler) discard() {
	com.next()
}

func (com *compiler) match(ty interface{}) bool {

	switch v := ty.(type) {
	case GenType:
		return v == GE_ANY || v == com.buff.GenType()

	case SubType:
		return v == SU_ANY || v == com.buff.SubType()
	}

	failNow("com.Match requires a GenType or SubType as an argument")
	return false
}
