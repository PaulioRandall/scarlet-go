package lexor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// DummyScanner is a ScanToken implementation that returns the token elements
// recursively.
func DummyScanner(tokens []token.Token) ScanToken {
	return func() (t token.Token, st ScanToken, _ ScanErr) {
		size := len(tokens)

		if size > 0 {
			t = tokens[0]
		}

		if size > 1 {
			st = DummyScanner(tokens[1:])
		}

		return
	}
}

// ScanTokenTest performs a test of a ScanToken implemention.
func ScanTokenTest(t *testing.T, f ScanToken, exps ...token.Token) {

	var tok token.Token
	var e error

	var size int = len(exps)
	var i int

	for i = 0; f != nil; i++ {

		tok, f, e = f()
		require.Nil(t, e)

		if size == 0 {
			break
		}

		if i < size {
			assert.Equal(t, exps[i], tok, "Token[%d]", i)
		}
	}

	assert.Equal(t, size, i, "Expected %d tokens but got %d", size, i)
}

// ScanTokenErrTest performs a test that the ScanToken implemention returns the
// expected error.
func ScanTokenErrTest(t *testing.T, f ScanToken, expAt int) error {

	var tok token.Token
	var e error

	for i := 0; i < expAt; i++ {
		tok, f, e = f()

		require.Nil(t, e,
			"Error occurred sooner than expected, Token[%d]", i)
		require.NotNil(t, f,
			"Token stream ended sooner than expected, Token[%d]", i)
	}

	tok, f, e = f()

	require.NotNil(t, e, "Error not returned when expected")
	require.Empty(t, tok, "Returned Token must be empty upon error")
	require.Nil(t, f, "Returned ScanToken function must be nil upon error")

	return e
}
