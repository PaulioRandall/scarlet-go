package scan

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"

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

func tok(gen GenType, sub SubType, raw string, line, begin, end int, props ...Prop) token.Tok {
	return token.Tok{
		Gen:      gen,
		Sub:      sub,
		RawProps: props,
		RawStr:   raw,
		Line:     line,
		ColBegin: begin,
		ColEnd:   end,
	}
}

func halfTok(gen GenType, sub SubType, raw string, props ...Prop) token.Tok {
	return token.Tok{
		Gen:      gen,
		Sub:      sub,
		RawProps: props,
		RawStr:   raw,
		ColEnd:   len(raw),
	}
}

func Test_S1(t *testing.T) {

	in := "@Set(x, 1)"

	exp := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Set", 0, 0, 4, PR_CALLABLE, PR_SPELL),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "(", 0, 4, 5, PR_DELIMITER, PR_PARENTHESIS, PR_OPENER),
		tok(GEN_IDENTIFIER, SUB_IDENTIFIER, "x", 0, 5, 6, PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ",", 0, 6, 7, PR_DELIMITER, PR_SEPARATOR),
		tok(GEN_REDUNDANT, SUB_WHITESPACE, " ", 0, 7, 8, PR_REDUNDANT, PR_WHITESPACE),
		tok(GEN_LITERAL, SUB_NUMBER, "1", 0, 8, 9),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")", 0, 9, 10, PR_DELIMITER, PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}

func Test_T0_1(t *testing.T) {
	doErrorTest(t, "~")
}

func Test_T1_1(t *testing.T) {
	doTest(t, " \t\v\f", []token.Token{
		halfTok(GEN_REDUNDANT, SUB_WHITESPACE, " \t\v\f", PR_REDUNDANT, PR_WHITESPACE),
	})
}

func Test_T2_1(t *testing.T) {
	doTest(t, ";", []token.Token{
		halfTok(GEN_TERMINATOR, SUB_TERMINATOR, ";", PR_TERMINATOR),
	})
}

func Test_T2_2(t *testing.T) {
	doTest(t, "\n", []token.Token{
		halfTok(GEN_TERMINATOR, SUB_NEWLINE, "\n", PR_TERMINATOR, PR_NEWLINE),
	})
}

func Test_T2_3(t *testing.T) {
	doTest(t, "\r\n", []token.Token{
		halfTok(GEN_TERMINATOR, SUB_NEWLINE, "\r\n", PR_TERMINATOR, PR_NEWLINE),
	})
}

func Test_T2_4(t *testing.T) {
	doErrorTest(t, "\r")
}

func Test_T3_1(t *testing.T) {
	doTest(t, "false", []token.Token{
		halfTok(GEN_LITERAL, SUB_BOOL, "false", PR_TERM, PR_LITERAL, PR_BOOL),
	})
}

func Test_T3_2(t *testing.T) {
	doTest(t, "true", []token.Token{
		halfTok(GEN_LITERAL, SUB_BOOL, "true", PR_TERM, PR_LITERAL, PR_BOOL),
	})
}

func Test_T4_1(t *testing.T) {
	doTest(t, "1", []token.Token{
		halfTok(GEN_LITERAL, SUB_NUMBER, "1"),
	})
}

func Test_T4_2(t *testing.T) {
	doTest(t, "123", []token.Token{
		halfTok(GEN_LITERAL, SUB_NUMBER, "123"),
	})
}

func Test_T4_3(t *testing.T) {
	doTest(t, "1.0", []token.Token{
		halfTok(GEN_LITERAL, SUB_NUMBER, "1.0"),
	})
}

func Test_T4_4(t *testing.T) {
	doTest(t, "123.456", []token.Token{
		halfTok(GEN_LITERAL, SUB_NUMBER, "123.456"),
	})
}

func Test_T4_5(t *testing.T) {
	doErrorTest(t, "123.")
}

func Test_T5_1(t *testing.T) {
	doTest(t, `""`, []token.Token{
		halfTok(GEN_LITERAL, SUB_STRING, `""`, PR_TERM, PR_LITERAL, PR_STRING),
	})
}

func Test_T5_2(t *testing.T) {
	doTest(t, `"abc"`, []token.Token{
		halfTok(GEN_LITERAL, SUB_STRING, `"abc"`, PR_TERM, PR_LITERAL, PR_STRING),
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
		halfTok(GEN_IDENTIFIER, SUB_IDENTIFIER, "a", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_2(t *testing.T) {
	doTest(t, "abc", []token.Token{
		halfTok(GEN_IDENTIFIER, SUB_IDENTIFIER, "abc", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_3(t *testing.T) {
	doTest(t, "a_b", []token.Token{
		halfTok(GEN_IDENTIFIER, SUB_IDENTIFIER, "a_b", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_4(t *testing.T) {
	doTest(t, "ab_", []token.Token{
		halfTok(GEN_IDENTIFIER, SUB_IDENTIFIER, "ab_", PR_TERM, PR_ASSIGNEE, PR_IDENTIFIER),
	})
}

func Test_T6_5(t *testing.T) {
	doTest(t, "_", []token.Token{
		halfTok(GEN_IDENTIFIER, SUB_VOID, "_", PR_ASSIGNEE, PR_VOID),
	})
}

func Test_T7_1(t *testing.T) {
	doTest(t, "(", []token.Token{
		halfTok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "(", PR_DELIMITER, PR_PARENTHESIS, PR_OPENER),
	})
}

func Test_T7_2(t *testing.T) {
	doTest(t, ")", []token.Token{
		halfTok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")", PR_DELIMITER, PR_PARENTHESIS, PR_CLOSER),
	})
}

func Test_T8_1(t *testing.T) {
	doTest(t, "@abc", []token.Token{
		halfTok(GEN_SPELL, SUB_UNDEFINED, "@abc", PR_CALLABLE, PR_SPELL),
	})
}

func Test_T8_2(t *testing.T) {
	doTest(t, "@abc.xyz", []token.Token{
		halfTok(GEN_SPELL, SUB_UNDEFINED, "@abc.xyz", PR_CALLABLE, PR_SPELL),
	})
}

func Test_T8_3(t *testing.T) {
	doTest(t, "@a.b.c.d", []token.Token{
		halfTok(GEN_SPELL, SUB_UNDEFINED, "@a.b.c.d", PR_CALLABLE, PR_SPELL),
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
		halfTok(GEN_DELIMITER, SUB_VALUE_DELIM, ",", PR_DELIMITER, PR_SEPARATOR),
	})
}

func Test_T10_1(t *testing.T) {
	doTest(t, "# abc", []token.Token{
		halfTok(GEN_REDUNDANT, SUB_COMMENT, "# abc", PR_REDUNDANT, PR_COMMENT),
	})
}
