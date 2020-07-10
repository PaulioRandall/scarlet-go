package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

type CompileFunc func() (inst.Instruction, CompileFunc, error)

type TokenStream interface {
	Next() token.Token
}

func New(ts TokenStream) CompileFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	com := &compiler{
		Queue: inst.Queue{},
		ts:    ts,
	}
	com.buff = com.ts.Next()

	if com.empty() {
		return nil
	}

	return com.compile
}

func StreamAll(ts TokenStream) ([]inst.Instruction, error) {

	var (
		e   error
		in  inst.Instruction
		ins = []inst.Instruction{}
	)

	for f := New(ts); f != nil; {
		if in, f, e = f(); e != nil {
			return nil, e
		}

		ins = append(ins, in)
	}

	return ins, nil
}

func CompileAll(tks []token.Token) ([]inst.Instruction, error) {
	return StreamAll(token.NewStream(tks))
}
