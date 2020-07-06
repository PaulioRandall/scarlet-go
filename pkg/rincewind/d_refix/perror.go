package refix

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

const (
	ERR_UNEXPECTED_EOF     string = "GROUP_ERR_UNEXPECTED_EOF"
	ERR_UNEXPECTED_TOKEN   string = "GROUP_ERR_UNEXPECTED_TOKEN"
	ERR_WRONG_TOKEN        string = "GROUP_ERR_WRONG_TOKEN"
	ERR_MISSING_EXPRESSION string = "GROUP_ERR_MISSING_EXPRESSION"
)

func errorUnexpectedEOF(rfx *refixer) error {
	return fail(rfx.buff, ERR_UNEXPECTED_EOF, "Want %q, have EOF")
}

func errorUnexpectedToken(rfx *refixer) error {
	return fail(rfx.buff, ERR_UNEXPECTED_TOKEN,
		"Token not expected here %q", rfx.buff.String())
}

func errorWrongToken(rfx *refixer, want Token) error {
	return fail(rfx.buff, ERR_WRONG_TOKEN,
		"Want %q, have %q", want.String(), rfx.buff.String())
}

func errorMissingExpression(rfx *refixer) error {
	return fail(rfx.buff, ERR_MISSING_EXPRESSION, "Missing expression")
}

func fail(snip perror.Snippet, code, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewBySnippet(code, msg, snip)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}
