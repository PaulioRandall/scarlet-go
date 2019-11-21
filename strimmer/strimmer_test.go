package strimmer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/scarlet-go/token"
)

// TODO: Potential for abstraction with scanner_2_test.go.
func doTestWrap(t *testing.T, in string, exps ...token.Token) {

	var tok token.Token
	var st token.ScanToken = New(in)
	var e error

	var size int = len(exps)
	var n int

	for i := 0; st != nil; i++ {
		n = i

		tok, st, e = st()
		require.Nil(t, e)

		if i < size {
			assert.Equal(t, exps[i], tok, "Token[%d]", i)
		}
	}

	assert.Equal(t, n, size, "Expected %d tokens but got %d", size, n)
}

func TestWrap_1(t *testing.T) {
	doTestWrap(t,
		"PROCEDURE",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
	)
}

func TestWrap_2(t *testing.T) {
	doTestWrap(t,
		"PROCEDURE\nEND",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\n", token.NEWLINE, 0, 9, 10),
		token.NewFlat("END", token.END, 1, 0, 3),
	)
}

func TestWrap_3(t *testing.T) {
	doTestWrap(t,
		"\t\t\t",
	)
}

func TestWrap_4(t *testing.T) {
	doTestWrap(t,
		"PROCEDURE\t\tEND",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("END", token.END, 0, 11, 14),
	)
}
