package refix

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

type piper struct {
	ts token.Stream
}

func (p piper) Next() interface{} {
	return p.ts.Next()
}

func (p piper) Match(ifaceTk, ifacePat interface{}) bool {

	if ifaceTk == nil {
		return false
	}

	tk, ok := ifaceTk.(token.Token)
	if !ok {
		failNow("refixer pipestack contains something other than a Token")
	}

	if pat, ok := ifacePat.(GenType); ok {
		return pat == GE_ANY || pat == tk.GenType()
	}

	if pat, ok := ifacePat.(SubType); ok {
		return pat == SU_ANY || pat == tk.SubType()
	}

	failNow("refixer.Match requires a GenType or SubType as the second argument")
	return false
}

func (p piper) Expect(ifaceTk, ifacePat interface{}) error {

	if ifaceTk == nil {
		return errorUnexpectedEOF(ifaceTk.(token.Token))
	}

	if ifacePat == nil {
		failNow("GenType or SubType required")
	}

	if p.Match(ifaceTk, ifacePat) {
		return nil
	}

	return errorWrongToken(ifacePat.(fmt.Stringer), ifaceTk.(token.Token))
}
