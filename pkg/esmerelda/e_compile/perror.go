package compile

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func errorUnexpectedEOF(com *compiler) error {
	return fail(com.buff, "Want %q, have EOF")
}

func errorUnexpectedToken(have token.Token) error {
	return fail(have, "Token not expected here %q", have.String())
}

func errorWrongToken(want fmt.Stringer, have token.Token) error {
	return fail(have, "Want %q, have %q", want.String(), have.String())
}

func fail(snip perror.Snippet, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewBySnippet(msg, snip)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}