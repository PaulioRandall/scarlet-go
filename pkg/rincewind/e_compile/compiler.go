package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

type compiler struct {
	inst.Queue
	ts   TokenStream
	buff token.Token
}

func (com *compiler) compile() (inst.Instruction, CompileFunc, error) {

	if com.Queue.Empty() {
		if e := next(com); e != nil {
			return inst.Inst{}, nil, e
		}
	}

	in := com.Take()
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
		return v == GEN_ANY || v == com.buff.GenType()

	case SubType:
		return v == SUB_ANY || v == com.buff.SubType()
	}

	failNow("com.Match requires a GenType or SubType as an argument")
	return false
}
