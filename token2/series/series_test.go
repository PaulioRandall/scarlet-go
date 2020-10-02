package series

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	//"github.com/stretchr/testify/require"
)

func TestSeries_Prepend(t *testing.T) {

	a, b, c, d := dummyLexemes()
	s := New()

	doTest := func(lexs ...lexeme.Lexeme) {
		expHead, expTail, _ := chainLexemes(lexs...)
		requireChain(t, expHead, s.head)
		requireChain(t, expTail, s.tail)
	}

	s.Prepend(d)
	doTest(d)

	s.Prepend(c)
	doTest(c, d)

	s.Prepend(b)
	doTest(b, c, d)

	s.Prepend(a)
	doTest(a, b, c, d)
}
