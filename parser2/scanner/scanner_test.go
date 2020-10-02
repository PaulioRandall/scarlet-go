package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/series"
	"github.com/PaulioRandall/scarlet-go/token2/token"

	"github.com/PaulioRandall/scarlet-go/token2/tokentest"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp *series.Series) {
	act, e := ScanString(in)
	require.Nil(t, e, "%+v", e)
	tokentest.RequireSeries(t, exp, act)
}

func doErrTest(t *testing.T, in string) {
	_, e := ScanString(in)
	require.NotNil(t, e, "Expected an error for input %q", in)
}

func TestBadToken(t *testing.T) {
	doErrTest(t, "¬")
}

func TestNewline_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("\n", token.NEWLINE))
	doTest(t, "\n", exp)
}

func TestNewline_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("\r\n", token.NEWLINE))
	doTest(t, "\r\n", exp)
}

func TestSpace_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(" ", token.SPACE))
	doTest(t, " ", exp)
}

func TestSpace_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("\t\r\v\f ", token.SPACE))
	doTest(t, "\t\r\v\f ", exp)
}

func TestComment_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("# :)", token.COMMENT))
	doTest(t, "# :)", exp)
}

func TestBool_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("true", token.TRUE))
	doTest(t, "true", exp)
}

func TestBool_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("false", token.FALSE))
	doTest(t, "false", exp)
}

func TestLoop_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("loop", token.LOOP))
	doTest(t, "loop", exp)
}

func TestIdent_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("abc", token.IDENT))
	doTest(t, "abc", exp)
}

func TestIdent_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("abc_xyz", token.IDENT))
	doTest(t, "abc_xyz", exp)
}

func TestTerminator_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(";", token.TERMINATOR))
	doTest(t, ";", exp)
}

func TestAssign_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(":=", token.ASSIGN))
	doTest(t, ":=", exp)
}

func TestDelim_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(",", token.DELIM))
	doTest(t, ",", exp)
}

func TestLeftParen_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("(", token.L_PAREN))
	doTest(t, "(", exp)
}

func TestRightParen_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(")", token.R_PAREN))
	doTest(t, ")", exp)
}

func TestLeftSquare_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("[", token.L_SQUARE))
	doTest(t, "[", exp)
}

func TestRightSquare_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("]", token.R_SQUARE))
	doTest(t, "]", exp)
}

func TestLeftCurly_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("{", token.L_CURLY))
	doTest(t, "{", exp)
}

func TestRightCurly_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("}", token.R_CURLY))
	doTest(t, "}", exp)
}

func TestVoid_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("_", token.VOID))
	doTest(t, "_", exp)
}

func TestAdd_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("+", token.ADD))
	doTest(t, "+", exp)
}

func TestSub_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("-", token.SUB))
	doTest(t, "-", exp)
}

func TestMul_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("*", token.MUL))
	doTest(t, "*", exp)
}

func TestDiv_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("/", token.DIV))
	doTest(t, "/", exp)
}

func TestRem_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("%", token.REM))
	doTest(t, "%", exp)
}

func TestAnd_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("&&", token.AND))
	doTest(t, "&&", exp)
}

func TestOr_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("||", token.OR))
	doTest(t, "||", exp)
}

func TestLessEqual_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("<=", token.LESS_EQUAL))
	doTest(t, "<=", exp)
}

func TestLess_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("<", token.LESS))
	doTest(t, "<", exp)
}

func TestMoreEqual_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(">=", token.MORE_EQUAL))
	doTest(t, ">=", exp)
}

func TestMore_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(">", token.MORE))
	doTest(t, ">", exp)
}

func TestEqual_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("==", token.EQUAL))
	doTest(t, "==", exp)
}

func TestNotEqual_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("!=", token.NOT_EQUAL))
	doTest(t, "!=", exp)
}

func TestSpell_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("@abc", token.SPELL))
	doTest(t, "@abc", exp)
}

func TestSpell_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("@a.b.c", token.SPELL))
	doTest(t, "@a.b.c", exp)
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestString_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(`""`, token.STRING))
	doTest(t, `""`, exp)
}

func TestString_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(`"abc"`, token.STRING))
	doTest(t, `"abc"`, exp)
}

func TestString_3(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok(`"\""`, token.STRING))
	doTest(t, `"\""`, exp)
}

func TestString_4(t *testing.T) {
	doErrTest(t, `"`)
}

func TestString_5(t *testing.T) {
	doErrTest(t, `"abc`)
}

func TestString_6(t *testing.T) {
	doErrTest(t, `"\"`)
}

func TestString_7(t *testing.T) {
	doErrTest(t, "\"\n\"")
}

func TestNumber_1(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("123", token.NUMBER))
	doTest(t, "123", exp)
}

func TestNumber_2(t *testing.T) {
	exp := tokentest.FeignSeries(lexeme.MakeTok("123.456", token.NUMBER))
	doTest(t, "123.456", exp)
}

func TestNumber_3(t *testing.T) {
	doErrTest(t, "123.")
}

func TestNumber_4(t *testing.T) {
	doErrTest(t, "123.abc")
}

func TestComprehensive_1(t *testing.T) {

	in := `x := 1 + 2
@Println("x = ", x)`

	tm := &position.TextMarker{}
	genLex := func(v string, tk token.Token) lexeme.Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v, v == "\n")
		return lexeme.Make(v, tk, snip)
	}

	exp := tokentest.FeignSeries(
		genLex("x", token.IDENT),
		genLex(" ", token.SPACE),
		genLex(":=", token.ASSIGN),
		genLex(" ", token.SPACE),
		genLex("1", token.NUMBER),
		genLex(" ", token.SPACE),
		genLex("+", token.ADD),
		genLex(" ", token.SPACE),
		genLex("2", token.NUMBER),
		genLex("\n", token.NEWLINE),
		genLex("@Println", token.SPELL),
		genLex("(", token.L_PAREN),
		genLex(`"x = "`, token.STRING),
		genLex(",", token.DELIM),
		genLex(" ", token.SPACE),
		genLex("x", token.IDENT),
		genLex(")", token.R_PAREN),
	)

	doTest(t, in, exp)
}
