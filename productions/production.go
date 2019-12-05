package productions

import (
	"github.com/PaulioRandall/scarlet-go/cookies"
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenEmitter is a thunk prototype that emits a token when called.
type TokenEmitter func() (token.Token, TokenEmitter, perror.Perror)

// newlineEmitter returns a TokenEmitter that returns a newline token along with
// the next emitter to use `after`.
func newlineEmitter(src *source.Source, after TokenEmitter) TokenEmitter {
	return func() (token.Token, TokenEmitter, perror.Perror) {

		n := cookies.NewlineRunes(src.Runes(), 0)
		if n == 0 {
			return token.Empty(), nil, perror.Newish(
				"Expected newline characters, i.e. LF or CRLF",
				src.Where(),
			)
		}

		return src.SliceNewline(n, token.NEWLINE), after, nil
	}
}
