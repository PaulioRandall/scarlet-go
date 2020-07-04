package group

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
	ERR_MISSING_STATEMENT  string = "GROUP_ERR_MISSING_STATEMENT"
)

func errorUnexpectedEOF(clt *collector) error {
	return fail(clt.buff, ERR_UNEXPECTED_EOF, "Want %q, have EOF")
}

func errorUnexpectedToken(clt *collector) error {
	return fail(clt.buff, ERR_UNEXPECTED_TOKEN,
		"Token not expected here %q", clt.buff.String())
}

func errorWrongToken(clt *collector, want Token) error {
	return fail(clt.buff, ERR_WRONG_TOKEN,
		"Want %q, have %q", want.String(), clt.buff.String())
}

func errorMissingExpression(clt *collector) error {
	return fail(clt.buff, ERR_MISSING_EXPRESSION, "Missing expression")
}

func errorMissingStatement(clt *collector) error {
	return fail(clt.buff, ERR_MISSING_STATEMENT, "Missing statement")
}

func fail(snip perror.Snippet, code, msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return perror.NewBySnippet(code, msg, snip)
}

func failNow(msg string, args ...interface{}) {
	perror.Panic(msg, args...)
}
