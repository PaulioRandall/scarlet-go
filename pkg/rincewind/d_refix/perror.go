package refix

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/perror"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
)

func errorUnexpectedEOF(want token.Token) error {
	return fail(want, "Want %q, have EOF")
}

func errorUnexpectedToken(have token.Token) error {
	return fail(have, "Token not expected here %q", have.String())
}

func errorWrongToken(want fmt.Stringer, have token.Token) error {
	return fail(have, "Want %q, have %q", want.String(), have.String())
}

func errorMissingExpression(have token.Token) error {
	return fail(have, "Missing expression")
}

func fail(snip perror.Snippet, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewBySnippet(msg, snip)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}
