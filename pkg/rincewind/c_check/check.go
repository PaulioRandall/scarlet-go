package check

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/queue"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type CheckFunc func() (token.Token, CheckFunc, error)

func New(ts token.Stream) CheckFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	chk := &checker{
		Queue: queue.Queue{},
		ts:    ts,
	}
	chk.bufferNext()

	if chk.empty() {
		return nil
	}

	return chk.check
}
