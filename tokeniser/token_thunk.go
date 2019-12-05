// tokeniser is responsible for recursively suppling tokens from the source
// code and ensuring they follow the production rules.
package tokeniser

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TokenThunk is a recursive thunk prototype that, when called, emits a
// token. The next thunk is also returned to allow recursive tokenisation.
// The end of the token stream is reached once the TokenThunk becomes nil.
//
// E.g:
// for thunk := tokeniser.New(src_code); thunk != nil; {
//   tok, thunk, perr := thunk()
//   // ...check error and do something with token...
// }
type TokenThunk func() (token.Token, TokenThunk, perror.Perror)

// New creates a TokenThunk thunk for the first token in the supplied source
// code.
func New(src string) TokenThunk {
	//source.New(src)
	return nil
}
