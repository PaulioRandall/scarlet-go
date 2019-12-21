package lexor

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// DummyScanToken creates a ScanToken function that returns the supplied token
// elements recursively.
func DummyScanToken(tokens []token.Token) ScanToken {

	size := len(tokens)
	i := 0

	return func() (t token.Token, st ScanToken, _ ScanErr) {

		if i >= size {
			return
		}

		t = tokens[i]
		i++

		if size-i > 0 {
			st = DummyScanToken(tokens)
		}

		return
	}
}

// AssertScanErr assert a ScanErr matches another except for the error message.
func AssertScanErr(t *testing.T, exp ScanErr, act ScanErr) {
	if exp == nil {
		assert.Nil(t, act, "Expected a nil ScanErr")
		return
	}

	require.NotNil(t, act, "Did not expect a nil ScanErr")
	assert.Equal(t, exp.Line(), act.Line(), "Wrong line index")
	assert.Equal(t, exp.Col(), act.Col(), "Wrong column index")

	if exp.Unwrap() == nil {
		assert.Nil(t, act.Unwrap(), "Did not expected a cause, a wrapped error")
	} else {
		assert.NotNil(t, act.Unwrap(), "Expected a cause, a wrapped error")
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
