package scan

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exps []token.Token) {

	acts, e := ScanAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func doErrorTest(t *testing.T, in string) {
	_, e := ScanAll(in)
	require.NotNil(t, e, "Expected an error")
}

func tok(raw string, line, begin, end int, props ...Prop) token.Tok {
	return token.Tok{
		RawProps: props,
		RawStr:   raw,
		Line:     line,
		ColBegin: begin,
		ColEnd:   end,
	}
}

func halfTok(raw string, props ...Prop) token.Tok {
	return token.Tok{
		RawProps: props,
		RawStr:   raw,
		ColEnd:   len(raw),
	}
}

func Test_S1(t *testing.T) {

	in := "@Set(x, 1)"

	exp := []token.Token{
		tok("@Set", 0, 0, 4, PR_CALLABLE, PR_SPELL),
		tok("(", 0, 4, 5, PR_DELIMITER, PR_PARENTHESIS, PR_OPENER),
		tok("x", 0, 5, 6, PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
		tok(",", 0, 6, 7, PR_DELIMITER, PR_SEPARATOR),
		tok(" ", 0, 7, 8, PR_REDUNDANT, PR_WHITESPACE),
		tok("1", 0, 8, 9, PR_TERM, PR_LITERAL, PR_NUMBER),
		tok(")", 0, 9, 10, PR_DELIMITER, PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}

func Test_T0_1(t *testing.T) {
	doErrorTest(t, "~")
}

func Test_T1_1(t *testing.T) {
	doTest(t, " \t\v\f", []token.Token{
		halfTok(" \t\v\f", PR_REDUNDANT, PR_WHITESPACE),
	})
}

func Test_T2_1(t *testing.T) {
	doTest(t, ";", []token.Token{
		halfTok(";", PR_TERMINATOR),
	})
}

func Test_T2_2(t *testing.T) {
	doTest(t, "\n", []token.Token{
		halfTok("\n", PR_TERMINATOR, PR_NEWLINE),
	})
}

func Test_T2_3(t *testing.T) {
	doTest(t, "\r\n", []token.Token{
		halfTok("\r\n", PR_TERMINATOR, PR_NEWLINE),
	})
}

func Test_T2_4(t *testing.T) {
	doErrorTest(t, "\r")
}

func Test_T3_1(t *testing.T) {
	doTest(t, "false", []token.Token{
		halfTok("false", PR_TERM, PR_LITERAL, PR_BOOL),
	})
}

func Test_T3_2(t *testing.T) {
	doTest(t, "true", []token.Token{
		halfTok("true", PR_TERM, PR_LITERAL, PR_BOOL),
	})
}

func Test_T4_1(t *testing.T) {
	doTest(t, "1", []token.Token{
		halfTok("1", PR_TERM, PR_LITERAL, PR_NUMBER),
	})
}

func Test_T4_2(t *testing.T) {
	doTest(t, "123", []token.Token{
		halfTok("123", PR_TERM, PR_LITERAL, PR_NUMBER),
	})
}

func Test_T4_3(t *testing.T) {
	doTest(t, "1.0", []token.Token{
		halfTok("1.0", PR_TERM, PR_LITERAL, PR_NUMBER),
	})
}

func Test_T4_4(t *testing.T) {
	doTest(t, "123.456", []token.Token{
		halfTok("123.456", PR_TERM, PR_LITERAL, PR_NUMBER),
	})
}

func Test_T4_5(t *testing.T) {
	doErrorTest(t, "123.")
}

func Test_T5_1(t *testing.T) {
	doTest(t, `""`, []token.Token{
		halfTok(`""`, PR_TERM, PR_LITERAL, PR_STRING),
	})
}

func Test_T5_2(t *testing.T) {
	doTest(t, `"abc"`, []token.Token{
		halfTok(`"abc"`, PR_TERM, PR_LITERAL, PR_STRING),
	})
}

func Test_T5_3(t *testing.T) {
	doErrorTest(t, `"`)
}

func Test_T5_4(t *testing.T) {
	doErrorTest(t, `"abc`)
}

func Test_T5_5(t *testing.T) {
	doErrorTest(t, `"\"`)
}

func Test_T5_6(t *testing.T) {
	doErrorTest(t, `"\"\"abc\"\"`)
}

func Test_T6_1(t *testing.T) {
	doTest(t, "a", []token.Token{
		halfTok("a", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_2(t *testing.T) {
	doTest(t, "abc", []token.Token{
		halfTok("abc", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_3(t *testing.T) {
	doTest(t, "a_b", []token.Token{
		halfTok("a_b", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_4(t *testing.T) {
	doTest(t, "ab_", []token.Token{
		halfTok("ab_", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_5(t *testing.T) {
	doTest(t, "_", []token.Token{
		halfTok("_", PR_ASSIGNEE, PR_VOID),
	})
}

func Test_T7_1(t *testing.T) {
	doTest(t, "(", []token.Token{
		halfTok("(", PR_DELIMITER, PR_PARENTHESIS, PR_OPENER),
	})
}

func Test_T7_2(t *testing.T) {
	doTest(t, ")", []token.Token{
		halfTok(")", PR_DELIMITER, PR_PARENTHESIS, PR_CLOSER),
	})
}

func Test_T8_1(t *testing.T) {
	doTest(t, "@abc", []token.Token{
		halfTok("@abc", PR_CALLABLE, PR_SPELL),
	})
}

func Test_T8_2(t *testing.T) {
	doTest(t, "@abc.xyz", []token.Token{
		halfTok("@abc.xyz", PR_CALLABLE, PR_SPELL),
	})
}

func Test_T8_3(t *testing.T) {
	doTest(t, "@a.b.c.d", []token.Token{
		halfTok("@a.b.c.d", PR_CALLABLE, PR_SPELL),
	})
}

func Test_T8_4(t *testing.T) {
	doErrorTest(t, "@")
}

func Test_T8_5(t *testing.T) {
	doErrorTest(t, "@.")
}

func Test_T8_6(t *testing.T) {
	doErrorTest(t, "@a.")
}

func Test_T8_7(t *testing.T) {
	doErrorTest(t, "@a..a")
}

func Test_T9_1(t *testing.T) {
	doTest(t, ",", []token.Token{
		halfTok(",", PR_DELIMITER, PR_SEPARATOR),
	})
}

func Test_T10_1(t *testing.T) {
	doTest(t, "# abc", []token.Token{
		halfTok("# abc", PR_REDUNDANT, PR_COMMENT),
	})
}
