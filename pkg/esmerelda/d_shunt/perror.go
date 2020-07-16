package shunt

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func errorUnexpectedEOF(after token.Snippet) error {
	return perror.NewAfterSnippet("Unexpected EOF", after)
}

func errorUnexpectedToken(have token.Token) error {
	return fail(have, "Token not expected here %q", have.String())
}

func errorWrongToken(want []Prop, have token.Token) error {
	w := JoinProps(" & ", want...)
	h := JoinProps(" & ", have.Props()...)
	return fail(have, "Want %q, have %q", w, h)
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
