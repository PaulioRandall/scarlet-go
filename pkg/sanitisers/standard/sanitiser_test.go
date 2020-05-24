package standard

import (
	"strconv"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"

	"github.com/stretchr/testify/require"
)

func checkRemovesTerminators(t *testing.T, prev Token) {

	in := []Token{
		prev,
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
	}

	exp := []Token{prev}

	checkMany(t, exp, in)
}

func checkMany(t *testing.T, exp, in []Token) {

	out := SanitiseAll(in)

	expSize := len(exp)
	outSize := len(out)

	for i := 0; i < expSize || i < outSize; i++ {

		require.True(t, i < outSize,
			"Expected ("+tkStr(exp, i)+") but no actual tokens remain")

		require.True(t, i < expSize,
			"Didn't expect any more tokens but got ("+tkStr(out, i)+")")

		checkToken(t, exp[i], out[i])
	}
}

func checkIgnores(t *testing.T, tk Token) {
	out := SanitiseAll([]Token{tk})
	checkSize(t, 1, out)
	checkToken(t, tk, out[0])
}

func checkFormats(t *testing.T, exp, in Token) {
	out := SanitiseAll([]Token{in})
	checkSize(t, 1, out)
	checkToken(t, exp, out[0])
}

func checkRemoves(t *testing.T, tk Token) {
	out := SanitiseAll([]Token{tk})
	checkSize(t, 0, out)
}

func checkToken(t *testing.T, exp, act Token) {
	require.Equal(t, exp, act,
		"Expected ("+ToString(exp)+") but got ("+ToString(act)+")")
}

func checkSize(t *testing.T, exp int, acts []Token) {
	require.Equal(t, exp, len(acts),
		"Expected "+strconv.Itoa(exp)+
			" tokens but got "+strconv.Itoa(len(acts)))
}

func tkStr(tks []Token, i int) (_ string) {
	if i < len(tks) {
		return ToString(tks[i])
	}
	return
}

func Test_F1_Newline(t *testing.T) {

	in := []Token{
		tok{K_IDENTIFIER, M_IDENTIFIER, "", 0, 0},
		tok{K_NEWLINE, M_NEWLINE, "", 0, 0},
	}

	exp := []Token{
		tok{K_IDENTIFIER, M_IDENTIFIER, "", 0, 0},
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
	}

	checkMany(t, exp, in)
}

func Test_F2_String(t *testing.T) {
	checkFormats(t,
		tok{K_LITERAL, M_STRING, "string", 0, 0},
		tok{K_LITERAL, M_STRING, "`string`", 0, 0},
	)
}

func Test_F3_Template(t *testing.T) {
	checkFormats(t,
		tok{K_LITERAL, M_TEMPLATE, `template`, 0, 0},
		tok{K_LITERAL, M_TEMPLATE, `"template"`, 0, 0},
	)
}

func Test_I1_Func(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_FUNC, "", 0, 0})
}

func Test_I2_Fix(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_FIX, "", 0, 0})
}

func Test_I3_ID(t *testing.T) {
	checkIgnores(t, tok{K_IDENTIFIER, M_IDENTIFIER, "", 0, 0})
}

func Test_I4_Delim(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_DELIMITER, "", 0, 0})
}

func Test_I5_Assign(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_ASSIGN, "", 0, 0})
}

func Test_I6_Output(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_OUTPUT, "", 0, 0})
}

func Test_I7_BlockOpen(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_BLOCK_OPEN, "", 0, 0})
}

func Test_I8_BlockClose(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_BLOCK_CLOSE, "", 0, 0})
}

func Test_I9_ParenOpen(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_PAREN_OPEN, "", 0, 0})
}

func Test_I10_ParenClose(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_PAREN_CLOSE, "", 0, 0})
}

func Test_I11_List(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_LIST, "", 0, 0})
}

func Test_I12_Match(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_MATCH, "", 0, 0})
}

func Test_I13_GuardOpen(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_GUARD_OPEN, "", 0, 0})
}

func Test_I14_GuardClose(t *testing.T) {
	checkIgnores(t, tok{K_DELIMITER, M_GUARD_CLOSE, "", 0, 0})
}

func Test_I15_Spell(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_SPELL, "", 0, 0})
}

func Test_I16_Number(t *testing.T) {
	checkIgnores(t, tok{K_LITERAL, M_NUMBER, "", 0, 0})
}

func Test_I17_Bool(t *testing.T) {
	checkIgnores(t, tok{K_LITERAL, M_BOOL, "", 0, 0})
}

func Test_I18_Add(t *testing.T) {
	checkIgnores(t, tok{K_ARITHMETIC, M_ADD, "", 0, 0})
}

func Test_I19_Subtract(t *testing.T) {
	checkIgnores(t, tok{K_ARITHMETIC, M_SUBTRACT, "", 0, 0})
}

func Test_I20_Multiply(t *testing.T) {
	checkIgnores(t, tok{K_ARITHMETIC, M_MULTIPLY, "", 0, 0})
}

func Test_I21_Divide(t *testing.T) {
	checkIgnores(t, tok{K_ARITHMETIC, M_DIVIDE, "", 0, 0})
}

func Test_I22_Remainder(t *testing.T) {
	checkIgnores(t, tok{K_ARITHMETIC, M_REMAINDER, "", 0, 0})
}

func Test_I23_And(t *testing.T) {
	checkIgnores(t, tok{K_LOGIC, M_AND, "", 0, 0})
}

func Test_I24_Or(t *testing.T) {
	checkIgnores(t, tok{K_LOGIC, M_OR, "", 0, 0})
}

func Test_I25_Equal(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_EQUAL, "", 0, 0})
}

func Test_I26_NotEqual(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_NOT_EQUAL, "", 0, 0})
}

func Test_I27_LessThan(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_LESS_THAN, "", 0, 0})
}

func Test_I28_LessThanOrEqual(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_LESS_THAN_OR_EQUAL, "", 0, 0})
}

func Test_I29_MoreThan(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_MORE_THAN, "", 0, 0})
}

func Test_I30_MoreThanOrEqual(t *testing.T) {
	checkIgnores(t, tok{K_COMPARISON, M_MORE_THAN_OR_EQUAL, "", 0, 0})
}

func Test_I31_Void(t *testing.T) {
	checkIgnores(t, tok{K_IDENTIFIER, M_VOID, "", 0, 0})
}

func Test_I32_Func(t *testing.T) {
	checkIgnores(t, tok{K_KEYWORD, M_FUNC, "", 0, 0})
}

func Test_R1_Newline(t *testing.T) {
	checkRemoves(t, tok{K_NEWLINE, M_NEWLINE, "", 0, 0})
}

func Test_R2_Whitespace(t *testing.T) {
	checkRemoves(t, tok{K_REDUNDANT, M_WHITESPACE, "", 0, 0})
}

func Test_R3_Comment(t *testing.T) {
	checkRemoves(t, tok{K_REDUNDANT, M_COMMENT, "", 0, 0})
}

func Test_R4_Undefined(t *testing.T) {
	checkRemoves(t, tok{K_UNDEFINED, M_UNDEFINED, "", 0, 0})
}

func Test_R5_RepeatedTerminators(t *testing.T) {

	in := []Token{
		tok{K_IDENTIFIER, M_IDENTIFIER, "", 0, 0},
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
	}

	exp := []Token{
		tok{K_IDENTIFIER, M_IDENTIFIER, "", 0, 0},
		tok{K_DELIMITER, M_TERMINATOR, "", 0, 0},
	}

	checkMany(t, exp, in)
}

func Test_R6_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok{K_DELIMITER, M_DELIMITER, "", 0, 0})
}

func Test_R7_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok{K_DELIMITER, M_BLOCK_OPEN, "", 0, 0})
}

func Test_R8_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok{K_DELIMITER, M_BLOCK_CLOSE, "", 0, 0})
}

func Test_R9_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok{K_KEYWORD, M_MATCH, "", 0, 0})
}

func Test_R10_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok{K_KEYWORD, M_LIST, "", 0, 0})
}
