package evaluator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New(lexor.DummyScanToken([]token.Token{
			token.OfValue(token.ID, "abc"),
			token.OfValue(token.NEWLINE, "\n"),
			token.OfValue(token.WHITESPACE, "   "),
			token.OfValue(token.STR_LITERAL, "`efg`"),
			token.OfValue(token.STR_TEMPLATE, `"hij"`),
			token.OfValue(token.COMMENT, "// xyz"),
		})),
		token.OfValue(token.ID, "abc"),
		token.OfValue(token.NEWLINE, "\n"),
		token.OfValue(token.STR_LITERAL, "efg"),
		token.OfValue(token.STR_TEMPLATE, "hij"),
	)
}
