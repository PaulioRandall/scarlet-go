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
		tok(token.KIND_ID, "abc", 0, 0),
		tok(token.TERMINATOR, "\n", 0, 3),
		tok(token.STR, "efg", 1, 5),
	)
}
