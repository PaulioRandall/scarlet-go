package lexor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"
)

func TestEvaluator_Next_1(t *testing.T) {

	sc := NewScanner("abc" + "\n" +
		" \t\r\v\f" + "`efg`" + "// comment")

	ev := NewEvaluator(sc)

	doTest(t, ev,
		tok(token.ID, "abc", 0, 0),
		tok(token.TERMINATOR, "\n", 0, 3),
		tok(token.STR_LITERAL, "efg", 1, 5),
	)
}
