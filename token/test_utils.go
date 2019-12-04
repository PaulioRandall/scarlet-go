package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ScanTokenTest performs a test of a ScanToken implemention.
func ScanTokenTest(t *testing.T, f ScanToken, exps ...Token) {

	var tok Token
	var e error

	var size int = len(exps)
	var n int

	for i := 0; f != nil; i++ {
		n = i

		tok, f, e = f()
		require.Nil(t, e)

		if i < size {
			assert.Equal(t, exps[i], tok, "Token[%d]", i)
		}
	}

	assert.Equal(t, n, size, "Expected %d tokens but got %d", size, n)
}
