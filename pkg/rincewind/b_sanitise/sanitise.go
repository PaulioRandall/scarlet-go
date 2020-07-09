package sanitise

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type SanitiseFunc func() (token.Token, SanitiseFunc, error)

func New(ts token.Stream) SanitiseFunc {

	if ts == nil {
		perror.Panic("Non-nil TokenStream required")
	}

	san := &sanitiser{ts: ts}
	san.buff = san.ts.Next()

	if san.empty() {
		return nil
	}

	return san.next
}
