package sanitise

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

type SanitiseFunc func() (token.Token, SanitiseFunc, error)

type TokenStream interface {
	Next() token.Token
}

func New(ts TokenStream) SanitiseFunc {

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
