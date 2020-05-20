package matching

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func Test_F1_Newline(t *testing.T) {

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

func Test_F2_String(t *testing.T) {
	checkFormats(t,
		Token{STRING, "string", 0, 0},
		Token{STRING, "`string`", 0, 0},
	)
}

func Test_F3_Template(t *testing.T) {
	checkFormats(t,
		Token{TEMPLATE, `template`, 0, 0},
		Token{TEMPLATE, `"template"`, 0, 0},
	)
}

func Test_I1_Func(t *testing.T) {
	checkIgnores(t, Token{FUNC, "", 0, 0})
}

func Test_I2_Fix(t *testing.T) {
	checkIgnores(t, Token{FIX, "", 0, 0})
}

func Test_I3_ID(t *testing.T) {
	checkIgnores(t, Token{ID, "", 0, 0})
}

func Test_I4_Delim(t *testing.T) {
	checkIgnores(t, Token{DELIM, "", 0, 0})
}

func Test_I5_Assign(t *testing.T) {
	checkIgnores(t, Token{ASSIGN, "", 0, 0})
}

func Test_I6_Output(t *testing.T) {
	checkIgnores(t, Token{OUTPUT, "", 0, 0})
}

func Test_I7_BlockOpen(t *testing.T) {
	checkIgnores(t, Token{BLOCK_OPEN, "", 0, 0})
}

func Test_I8_BlockClose(t *testing.T) {
	checkIgnores(t, Token{BLOCK_CLOSE, "", 0, 0})
}

func Test_I9_ParenOpen(t *testing.T) {
	checkIgnores(t, Token{PAREN_OPEN, "", 0, 0})
}

func Test_I10_ParenClose(t *testing.T) {
	checkIgnores(t, Token{PAREN_CLOSE, "", 0, 0})
}

func Test_I11_List(t *testing.T) {
	checkIgnores(t, Token{LIST, "", 0, 0})
}

func Test_I12_Match(t *testing.T) {
	checkIgnores(t, Token{MATCH, "", 0, 0})
}

func Test_I13_GuardOpen(t *testing.T) {
	checkIgnores(t, Token{GUARD_OPEN, "", 0, 0})
}

func Test_I14_GuardClose(t *testing.T) {
	checkIgnores(t, Token{GUARD_CLOSE, "", 0, 0})
}

func Test_I15_Spell(t *testing.T) {
	checkIgnores(t, Token{SPELL, "", 0, 0})
}

func Test_I16_Number(t *testing.T) {
	checkIgnores(t, Token{NUMBER, "", 0, 0})
}

func Test_I17_Bool(t *testing.T) {
	checkIgnores(t, Token{BOOL, "", 0, 0})
}

func Test_I18_Add(t *testing.T) {
	checkIgnores(t, Token{ADD, "", 0, 0})
}

func Test_I19_Subtract(t *testing.T) {
	checkIgnores(t, Token{SUBTRACT, "", 0, 0})
}

func Test_I20_Multiply(t *testing.T) {
	checkIgnores(t, Token{MULTIPLY, "", 0, 0})
}

func Test_I21_Divide(t *testing.T) {
	checkIgnores(t, Token{DIVIDE, "", 0, 0})
}

func Test_I22_Remainder(t *testing.T) {
	checkIgnores(t, Token{REMAINDER, "", 0, 0})
}

func Test_I23_And(t *testing.T) {
	checkIgnores(t, Token{AND, "", 0, 0})
}

func Test_I24_Or(t *testing.T) {
	checkIgnores(t, Token{OR, "", 0, 0})
}

func Test_I25_Equal(t *testing.T) {
	checkIgnores(t, Token{EQUAL, "", 0, 0})
}

func Test_I26_NotEqual(t *testing.T) {
	checkIgnores(t, Token{NOT_EQUAL, "", 0, 0})
}

func Test_I27_LessThan(t *testing.T) {
	checkIgnores(t, Token{LESS_THAN, "", 0, 0})
}

func Test_I28_LessThanOrEqual(t *testing.T) {
	checkIgnores(t, Token{LESS_THAN_OR_EQUAL, "", 0, 0})
}

func Test_I29_MoreThan(t *testing.T) {
	checkIgnores(t, Token{MORE_THAN, "", 0, 0})
}

func Test_I30_MoreThanOrEqual(t *testing.T) {
	checkIgnores(t, Token{MORE_THAN_OR_EQUAL, "", 0, 0})
}

func Test_I31_Void(t *testing.T) {
	checkIgnores(t, Token{VOID, "", 0, 0})
}

func Test_I32_Func(t *testing.T) {
	checkIgnores(t, Token{FUNC, "", 0, 0})
}

func Test_R1_Newline(t *testing.T) {
	checkRemoves(t, Token{NEWLINE, "", 0, 0})
}

func Test_R2_Whitespace(t *testing.T) {
	checkRemoves(t, Token{WHITESPACE, "", 0, 0})
}

func Test_R3_Comment(t *testing.T) {
	checkRemoves(t, Token{COMMENT, "", 0, 0})
}

func Test_R4_Undefined(t *testing.T) {
	checkRemoves(t, Token{UNDEFINED, "", 0, 0})
}

func Test_R5_RepeatedTerminators(t *testing.T) {

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

func Test_R6_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, Token{DELIM, "", 0, 0})
}

func Test_R7_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, Token{BLOCK_OPEN, "", 0, 0})
}

func Test_R8_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, Token{BLOCK_CLOSE, "", 0, 0})
}

func Test_R9_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, Token{MATCH, "", 0, 0})
}

func Test_R10_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, Token{LIST, "", 0, 0})
}
