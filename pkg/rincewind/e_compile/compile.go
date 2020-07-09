package compile

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/queue"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type CompileFunc func() (inst.Instruction, CompileFunc, error)

func New(ts token.Stream) CompileFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	com := &compiler{
		Queue: queue.Queue{},
		ts:    ts,
	}
	com.buff = com.ts.Next()

	if com.empty() {
		return nil
	}

	return com.compile
}
