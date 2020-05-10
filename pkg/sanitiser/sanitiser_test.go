package sanitiser

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Test_I1(t *testing.T) {
	checkIgnores(t, Token{FUNC, "", 0, 0})
	checkIgnores(t, Token{FIX, "", 0, 0})
	checkIgnores(t, Token{ID, "", 0, 0})
	checkIgnores(t, Token{DELIM, "", 0, 0})
	checkIgnores(t, Token{ASSIGN, "", 0, 0})
	checkIgnores(t, Token{RETURNS, "", 0, 0})
	checkIgnores(t, Token{BLOCK_OPEN, "", 0, 0})
	checkIgnores(t, Token{BLOCK_CLOSE, "", 0, 0})
	checkIgnores(t, Token{PAREN_OPEN, "", 0, 0})
	checkIgnores(t, Token{PAREN_CLOSE, "", 0, 0})
	checkIgnores(t, Token{LIST, "", 0, 0})
	checkIgnores(t, Token{MATCH, "", 0, 0})
	checkIgnores(t, Token{GUARD_OPEN, "", 0, 0})
	checkIgnores(t, Token{GUARD_CLOSE, "", 0, 0})
	checkIgnores(t, Token{SPELL, "", 0, 0})
	checkIgnores(t, Token{NUMBER, "", 0, 0})
	checkIgnores(t, Token{BOOL, "", 0, 0})
	checkIgnores(t, Token{ADD, "", 0, 0})
	checkIgnores(t, Token{SUBTRACT, "", 0, 0})
	checkIgnores(t, Token{MULTIPLY, "", 0, 0})
	checkIgnores(t, Token{DIVIDE, "", 0, 0})
	checkIgnores(t, Token{REMAINDER, "", 0, 0})
	checkIgnores(t, Token{AND, "", 0, 0})
	checkIgnores(t, Token{OR, "", 0, 0})
	checkIgnores(t, Token{EQUAL, "", 0, 0})
	checkIgnores(t, Token{NOT_EQUAL, "", 0, 0})
	checkIgnores(t, Token{LESS_THAN, "", 0, 0})
	checkIgnores(t, Token{LESS_THAN_OR_EQUAL, "", 0, 0})
	checkIgnores(t, Token{MORE_THAN, "", 0, 0})
	checkIgnores(t, Token{MORE_THAN_OR_EQUAL, "", 0, 0})
	checkIgnores(t, Token{VOID, "", 0, 0})
	checkIgnores(t, Token{FUNC, "", 0, 0})
}

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

func Test_F1(t *testing.T) {

	in := []Token{
		Token{ID, "", 0, 0},
		Token{NEWLINE, "", 0, 0},
	}

	exp := []Token{
		Token{ID, "", 0, 0},
		Token{TERMINATOR, "", 0, 0},
	}

	checkMany(t, exp, in)
}
