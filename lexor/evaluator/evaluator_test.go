package evaluator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

func TestWrap_1(t *testing.T) {
	//

	lexor.ScanTokenTest(t,
		New(lexor.DummyScanToken([]token.Token{
			token.NewToken(token.ID, "abc", 0, 0),
			token.NewToken(token.NEWLINE, "\n", 0, 3),
			token.NewToken(token.WHITESPACE, "   ", 1, 0),
			token.NewToken(token.STR_LITERAL, "`efg`", 1, 3),
			token.NewToken(token.STR_TEMPLATE, `"hij"`, 1, 6),
			token.NewToken(token.COMMENT, "// xyz", 1, 9),
		})),
		token.NewToken(token.ID, "abc", 0, 0),
		token.NewToken(token.NEWLINE, "\n", 0, 3),
		token.NewToken(token.STR_LITERAL, "efg", 1, 3),
		token.NewToken(token.STR_TEMPLATE, "hij", 1, 6),
	)
}
