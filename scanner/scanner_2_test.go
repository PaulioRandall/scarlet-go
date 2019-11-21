package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/scarlet-go/token"
)

// TODO: Potential for abstraction with strimmer_test.go.
func doTestScanOK(t *testing.T, in string, exps ...token.Token) {

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

func TestScanner_Scan_OK_1(t *testing.T) {
	doTestScanOK(t,
		"PROCEDURE",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
	)
}

func TestScanner_Scan_OK_2(t *testing.T) {
	doTestScanOK(t,
		"PROCEDURE\nEND",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\n", token.NEWLINE, 0, 9, 10),
		token.NewFlat("END", token.END, 1, 0, 3),
	)
}

func TestScanner_Scan_OK_3(t *testing.T) {
	doTestScanOK(t,
		"\t\t\t",
		token.NewFlat("\t\t\t", token.WHITESPACE, 0, 0, 3),
	)
}

func TestScanner_Scan_OK_4(t *testing.T) {
	doTestScanOK(t,
		"PROCEDURE\t\tEND",
		token.NewFlat("PROCEDURE", token.PROCEDURE, 0, 0, 9),
		token.NewFlat("\t\t", token.WHITESPACE, 0, 9, 11),
		token.NewFlat("END", token.END, 0, 11, 14),
	)
}
