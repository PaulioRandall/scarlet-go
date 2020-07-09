package refix

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/pipestack"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type RefixFunc func() (token.Token, RefixFunc, error)

func New(ts token.Stream) RefixFunc {

	if ts == nil {
		failNow("Non-nil Token Stream required")
	}

	p := piper{ts}
	rfx := &refixer{
		NewPipeStack(p, p),
	}

	if rfx.Empty() {
		return nil
	}

	return rfx.refix
}
