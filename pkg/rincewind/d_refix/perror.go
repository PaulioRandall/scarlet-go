package refix

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

const (
	ERR_UNEXPECTED_EOF     string = "REFIX_ERR_UNEXPECTED_EOF"
	ERR_UNEXPECTED_TOKEN   string = "REFIX_ERR_UNEXPECTED_TOKEN"
	ERR_WRONG_TOKEN        string = "REFIX_ERR_WRONG_TOKEN"
	ERR_MISSING_EXPRESSION string = "REFIX_ERR_MISSING_EXPRESSION"
)

func errorUnexpectedEOF(want Token) error {
	return fail(want, ERR_UNEXPECTED_EOF, "Want %q, have EOF")
}

func errorUnexpectedToken(have Token) error {
	return fail(have, ERR_UNEXPECTED_TOKEN,
		"Token not expected here %q", have.String())
}

func errorWrongToken(want fmt.Stringer, have Token) error {
	return fail(have, ERR_WRONG_TOKEN,
		"Want %q, have %q", want.String(), have.String())
}

func errorMissingExpression(have Token) error {
	return fail(have, ERR_MISSING_EXPRESSION, "Missing expression")
}

func fail(snip perror.Snippet, code, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewBySnippet(code, msg, snip)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}
