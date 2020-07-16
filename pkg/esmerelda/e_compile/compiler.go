package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
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

func (com *compiler) peek() token.Token {

	if com.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	return com.buff
}

func (com *compiler) next() token.Token {

	if com.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	r := com.buff
	com.buff = com.ts.Next()

	return r
}

func (com *compiler) match(props ...Prop) bool {
	return com.buff.Is(props...)
}
