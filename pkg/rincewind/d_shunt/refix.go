package shunt

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type RefixFunc func() (token.Token, RefixFunc, error)

func New(ts token.Stream) RefixFunc {

	if ts == nil {
		failNow("Non-nil Token Stream required")
	}

	shy := &shuntingYard{
		Stack: &token.Stack{},
		ts:    ts,
	}
	shy.buff = shy.ts.Next()

	if shy.empty() {
		return nil
	}

	return shy.shunt
}
