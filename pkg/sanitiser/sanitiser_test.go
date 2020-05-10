package sanitiser

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Test_R1(t *testing.T) {
	checkRemoves(t, Token{NEWLINE, "", 0, 0})
}

func Test_R2(t *testing.T) {
	checkRemoves(t, Token{WHITESPACE, "", 0, 0})
}

func Test_R3(t *testing.T) {
	checkRemoves(t, Token{COMMENT, "", 0, 0})
}

func Test_R4(t *testing.T) {
	checkRemoves(t, Token{UNDEFINED, "", 0, 0})
}

func Test_R5(t *testing.T) {

	in := []Token{
		Token{ID, "", 0, 0},
		Token{TERMINATOR, "", 0, 0},
		Token{TERMINATOR, "", 0, 0},
	}

	exp := []Token{
		Token{ID, "", 0, 0},
		Token{TERMINATOR, "", 0, 0},
	}

	checkMany(t, exp, in)
}

func Test_R6(t *testing.T) {
	checkRemovesTerminators(t, Token{DELIM, "", 0, 0})
}

func Test_R7(t *testing.T) {
	checkRemovesTerminators(t, Token{BLOCK_OPEN, "", 0, 0})
}

func Test_R8(t *testing.T) {
	checkRemovesTerminators(t, Token{BLOCK_CLOSE, "", 0, 0})
}

func Test_R9(t *testing.T) {
	checkRemovesTerminators(t, Token{MATCH, "", 0, 0})
}

func Test_R10(t *testing.T) {
	checkRemovesTerminators(t, Token{LIST, "", 0, 0})
}
