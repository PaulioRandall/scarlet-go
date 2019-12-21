package evaluator

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/token"
)

func TestWrap_1(t *testing.T) {
	lexor.ScanTokenTest(t,
		New(lexor.DummyScanToken([]token.Token{
			token.NewToken(token.STR_LITERAL, "`abc`", 0, 0),
			token.NewToken(token.NEWLINE, "\n", 0, 5),
			token.NewToken(token.FUNC, "F", 1, 0),
		})),
		token.NewToken(token.STR_LITERAL, "abc", 0, 0),
		token.NewToken(token.NEWLINE, "\n", 0, 5),
		token.NewToken(token.FUNC, "F", 1, 0),
	)
}
