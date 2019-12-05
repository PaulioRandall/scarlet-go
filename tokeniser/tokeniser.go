// tokeniser is responsible for recursively suppling tokens from the source
// code and ensuring it follows the production rules.
package tokeniser

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenEmitter is a recursive thunk prototype that, when called, emits a
// token. The next emitter is also returned to allow recursive tokenisation.
// The end of the token stream is reached once the TokenEmitter becomes nil.
//
// E.g:
// for emitter := tokeniser.New(src_code); emitter != nil; {
//   tok, emitter, perr := emitter()
//   // ...check error and do something with token...
// }
type TokenEmitter func() (token.Token, TokenEmitter, perror.Perror)

// New creates a TokenEmitter thunk for the first token in the supplied source
// code.
func New(src string) TokenEmitter {
	//source.New(src)
	return nil
}
