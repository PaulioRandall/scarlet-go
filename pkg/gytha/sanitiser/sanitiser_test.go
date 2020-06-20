package sanitiser

import (
	"strconv"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/stretchr/testify/require"
)

func tok(ty TokenType, v string) Token {
	return NewToken(ty, v, 0, 0)
}

func checkRemovesTerminators(t *testing.T, prev Token) {

	in := []Token{
		prev,
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
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
		tok(TK_IDENTIFIER, ""),
		tok(TK_NEWLINE, ""),
	}

	exp := []Token{
		tok(TK_IDENTIFIER, ""),
		tok(TK_TERMINATOR, ""),
	}

	checkMany(t, exp, in)
}

func Test_F2_String(t *testing.T) {
	checkFormats(t,
		tok(TK_STRING, "string"),
		tok(TK_STRING, "`string`"),
	)
}

func Test_I1_Func(t *testing.T) {
	checkIgnores(t, tok(TK_FUNCTION, ""))
}

func Test_I2_Fix(t *testing.T) {
	checkIgnores(t, tok(TK_DEFINITION, ""))
}

func Test_I3_ID(t *testing.T) {
	checkIgnores(t, tok(TK_IDENTIFIER, ""))
}

func Test_I4_Delim(t *testing.T) {
	checkIgnores(t, tok(TK_DELIMITER, ""))
}

func Test_I5_Assign(t *testing.T) {
	checkIgnores(t, tok(TK_ASSIGNMENT, ""))
}

func Test_I6_Output(t *testing.T) {
	checkIgnores(t, tok(TK_OUTPUT, ""))
}

func Test_I7_BlockOpen(t *testing.T) {
	checkIgnores(t, tok(TK_BLOCK_OPEN, ""))
}

func Test_I8_BlockClose(t *testing.T) {
	checkIgnores(t, tok(TK_BLOCK_CLOSE, ""))
}

func Test_I9_ParenOpen(t *testing.T) {
	checkIgnores(t, tok(TK_PAREN_OPEN, ""))
}

func Test_I10_ParenClose(t *testing.T) {
	checkIgnores(t, tok(TK_PAREN_CLOSE, ""))
}

func Test_I11_List(t *testing.T) {
	checkIgnores(t, tok(TK_LIST, ""))
}

func Test_I12_When(t *testing.T) {
	checkIgnores(t, tok(TK_WHEN, ""))
}

func Test_I13_GuardOpen(t *testing.T) {
	checkIgnores(t, tok(TK_GUARD_OPEN, ""))
}

func Test_I14_GuardClose(t *testing.T) {
	checkIgnores(t, tok(TK_GUARD_CLOSE, ""))
}

func Test_I15_Spell(t *testing.T) {
	checkIgnores(t, tok(TK_SPELL, ""))
}

func Test_I16_Number(t *testing.T) {
	checkIgnores(t, tok(TK_NUMBER, ""))
}

func Test_I17_Bool(t *testing.T) {
	checkIgnores(t, tok(TK_BOOL, ""))
}

func Test_I18_Add(t *testing.T) {
	checkIgnores(t, tok(TK_PLUS, ""))
}

func Test_I19_Subtract(t *testing.T) {
	checkIgnores(t, tok(TK_MINUS, ""))
}

func Test_I20_Multiply(t *testing.T) {
	checkIgnores(t, tok(TK_MULTIPLY, ""))
}

func Test_I21_Divide(t *testing.T) {
	checkIgnores(t, tok(TK_DIVIDE, ""))
}

func Test_I22_Remainder(t *testing.T) {
	checkIgnores(t, tok(TK_REMAINDER, ""))
}

func Test_I23_And(t *testing.T) {
	checkIgnores(t, tok(TK_AND, ""))
}

func Test_I24_Or(t *testing.T) {
	checkIgnores(t, tok(TK_OR, ""))
}

func Test_I25_Equal(t *testing.T) {
	checkIgnores(t, tok(TK_EQUAL, ""))
}

func Test_I26_NotEqual(t *testing.T) {
	checkIgnores(t, tok(TK_NOT_EQUAL, ""))
}

func Test_I27_LessThan(t *testing.T) {
	checkIgnores(t, tok(TK_LESS_THAN, ""))
}

func Test_I28_LessThanOrEqual(t *testing.T) {
	checkIgnores(t, tok(TK_LESS_THAN_OR_EQUAL, ""))
}

func Test_I29_MoreThan(t *testing.T) {
	checkIgnores(t, tok(TK_MORE_THAN, ""))
}

func Test_I30_MoreThanOrEqual(t *testing.T) {
	checkIgnores(t, tok(TK_MORE_THAN_OR_EQUAL, ""))
}

func Test_I31_Void(t *testing.T) {
	checkIgnores(t, tok(TK_VOID, ""))
}

func Test_I32_Func(t *testing.T) {
	checkIgnores(t, tok(TK_FUNCTION, ""))
}

func Test_R1_Newline(t *testing.T) {
	checkRemoves(t, tok(TK_NEWLINE, ""))
}

func Test_R2_Whitespace(t *testing.T) {
	checkRemoves(t, tok(TK_WHITESPACE, ""))
}

func Test_R3_Comment(t *testing.T) {
	checkRemoves(t, tok(TK_COMMENT, ""))
}

func Test_R4_Undefined(t *testing.T) {
	checkRemoves(t, tok(TK_UNDEFINED, ""))
}

func Test_R5_RepeatedTerminators(t *testing.T) {

	in := []Token{
		tok(TK_IDENTIFIER, ""),
		tok(TK_TERMINATOR, ""),
		tok(TK_TERMINATOR, ""),
	}

	exp := []Token{
		tok(TK_IDENTIFIER, ""),
		tok(TK_TERMINATOR, ""),
	}

	checkMany(t, exp, in)
}

func Test_R6_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok(TK_DELIMITER, ""))
}

func Test_R7_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok(TK_BLOCK_OPEN, ""))
}

func Test_R8_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok(TK_BLOCK_CLOSE, ""))
}

func Test_R9_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok(TK_WHEN, ""))
}

func Test_R10_RepeatedTerminators(t *testing.T) {
	checkRemovesTerminators(t, tok(TK_LIST, ""))
}
