package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/queue"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token/types"
)

type CompileFunc func() (inst.Instruction, CompileFunc, error)

type TokenStream interface {
	Next() token.Token
}

type compiler struct {
	Queue
	ts   TokenStream
	buff token.Token
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

func (com *compiler) compile() (inst.Instruction, CompileFunc, error) {

	if com.Queue.Empty() {
		if e := next(com); e != nil {
			return nil, nil, e
		}
	}

	in := com.Take().(inst.Instruction)
	if com.empty() {
		return in, nil, nil
	}

	return in, com.compile, nil
}

func (com *compiler) empty() bool {
	return com.buff == nil && com.Queue.Empty()
}

func (com *compiler) next() token.Token {

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
