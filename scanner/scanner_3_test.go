package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/scarlet-go/cookies/perror"
	"github.com/PaulioRandall/scarlet-go/token"
)

// TODO: Potential for abstraction with strimmer_test.go.
func doTestScanErr(t *testing.T, in string, expAt int, expPerr perror.Perror) {

	var tok token.Token
	var st token.ScanToken = New(in)
	var e perror.Perror

	for i := 0; i < expAt; i++ {
		tok, st, e = st()

		require.Nil(t, e,
			"Error occurred sooner than expected, Token[%d]", i)
		require.NotNil(t, st,
			"Token stream ended sooner than expected, Token[%d]", i)
	}

	tok, st, e = st()

	require.Empty(t, tok, "Expected an empty token")
	require.Nil(t, st, "Expected a nil scan function")
	require.NotNil(t, e, "Expected an error")

	assert.Equal(t, expPerr.Where(), e.Where())
}

func TestScanner_Scan_Err_1(t *testing.T) {
	doTestScanErr(t,
		"~~~",
		0,
		perror.New("", 0, 0, 0),
	)
}

func TestScanner_Scan_Err_2(t *testing.T) {
	doTestScanErr(t,
		"PROCEDURE\n  ~~~\nEND",
		3,
		perror.New("", 1, 2, 2),
	)
}
