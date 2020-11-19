package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp []token.Lexeme) {
	act, e := ScanAll([]rune(in))
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func doErrTest(t *testing.T, in string) {
	_, e := ScanAll([]rune(in))
	require.NotNil(t, e, "Expected an error for input %q", in)
}

func TestBadToken(t *testing.T) {
	doErrTest(t, "¬")
}

func tok(s string, tk token.Token) token.Lexeme {
	l := token.Make(s, tk, token.Snippet{})
	lineCount := 1
	if s == "\n" || s == "\r\n" {
		lineCount++
	}
	l.Range = position.Rng(
		position.Pos("", 0, 0, 0, 0), lineCount, len(s),
	)
	return l
}

func TestNewline_1(t *testing.T) {
	doTest(t, "\n", []token.Lexeme{
		tok("\n", token.NEWLINE),
	})
}

func TestNewline_2(t *testing.T) {
	doTest(t, "\r\n", []token.Lexeme{
		tok("\r\n", token.NEWLINE),
	})
}

func TestSpace_1(t *testing.T) {
	doTest(t, " ", []token.Lexeme{
		tok(" ", token.SPACE),
	})
}

func TestSpace_2(t *testing.T) {
	doTest(t, "\t\v\f ", []token.Lexeme{
		tok("\t\v\f ", token.SPACE),
	})
}

func TestComment_1(t *testing.T) {
	doTest(t, "# :)", []token.Lexeme{
		tok("# :)", token.COMMENT),
	})
}

func TestBool_1(t *testing.T) {
	doTest(t, "true", []token.Lexeme{
		tok("true", token.TRUE),
	})
}

func TestBool_2(t *testing.T) {
	doTest(t, "false", []token.Lexeme{
		tok("false", token.FALSE),
	})
}

func TestLoop_1(t *testing.T) {
	doTest(t, "loop", []token.Lexeme{
		tok("loop", token.LOOP),
	})
}

func TestWhen_1(t *testing.T) {
	doTest(t, "when", []token.Lexeme{
		tok("when", token.WHEN),
	})
}

func TestIdent_1(t *testing.T) {
	doTest(t, "abc", []token.Lexeme{
		tok("abc", token.IDENT),
	})
}

func TestIdent_2(t *testing.T) {
	doTest(t, "abc_xyz", []token.Lexeme{
		tok("abc_xyz", token.IDENT),
	})
}

func TestTerminator_1(t *testing.T) {
	doTest(t, ";", []token.Lexeme{
		tok(";", token.TERMINATOR),
	})
}

func TestAssign_1(t *testing.T) {
	doTest(t, "<-", []token.Lexeme{
		tok("<-", token.ASSIGN),
	})
}

func TestDelim_1(t *testing.T) {
	doTest(t, ",", []token.Lexeme{
		tok(",", token.DELIM),
	})
}

func TestLeftParen_1(t *testing.T) {
	doTest(t, "(", []token.Lexeme{
		tok("(", token.L_PAREN),
	})
}

func TestRightParen_1(t *testing.T) {
	doTest(t, ")", []token.Lexeme{
		tok(")", token.R_PAREN),
	})
}

func TestLeftSquare_1(t *testing.T) {
	doTest(t, "[", []token.Lexeme{
		tok("[", token.L_SQUARE),
	})
}

func TestRightSquare_1(t *testing.T) {
	doTest(t, "]", []token.Lexeme{
		tok("]", token.R_SQUARE),
	})
}

func TestLeftCurly_1(t *testing.T) {
	doTest(t, "{", []token.Lexeme{
		tok("{", token.L_CURLY),
	})
}

func TestRightCurly_1(t *testing.T) {
	doTest(t, "}", []token.Lexeme{
		tok("}", token.R_CURLY),
	})
}

func TestVoid_1(t *testing.T) {
	doTest(t, "_", []token.Lexeme{
		tok("_", token.VOID),
	})
}

func TestAdd_1(t *testing.T) {
	doTest(t, "+", []token.Lexeme{
		tok("+", token.ADD),
	})
}

func TestSub_1(t *testing.T) {
	doTest(t, "-", []token.Lexeme{
		tok("-", token.SUB),
	})
}

func TestMul_1(t *testing.T) {
	doTest(t, "*", []token.Lexeme{
		tok("*", token.MUL),
	})
}

func TestDiv_1(t *testing.T) {
	doTest(t, "/", []token.Lexeme{
		tok("/", token.DIV),
	})
}

func TestRem_1(t *testing.T) {
	doTest(t, "%", []token.Lexeme{
		tok("%", token.REM),
	})
}

func TestAnd_1(t *testing.T) {
	doTest(t, "&&", []token.Lexeme{
		tok("&&", token.AND),
	})
}

func TestOr_1(t *testing.T) {
	doTest(t, "||", []token.Lexeme{
		tok("||", token.OR),
	})
}

func TestLessEqual_1(t *testing.T) {
	doTest(t, "<=", []token.Lexeme{
		tok("<=", token.LTE),
	})
}

func TestLess_1(t *testing.T) {
	doTest(t, "<", []token.Lexeme{
		tok("<", token.LT),
	})
}

func TestMoreEqual_1(t *testing.T) {
	doTest(t, ">=", []token.Lexeme{
		tok(">=", token.MTE),
	})
}

func TestMore_1(t *testing.T) {
	doTest(t, ">", []token.Lexeme{
		tok(">", token.MT),
	})
}

func TestEqual_1(t *testing.T) {
	doTest(t, "==", []token.Lexeme{
		tok("==", token.EQU),
	})
}

func TestNotEqual_1(t *testing.T) {
	doTest(t, "!=", []token.Lexeme{
		tok("!=", token.NEQ),
	})
}

func TestSpell_1(t *testing.T) {
	doTest(t, "@abc", []token.Lexeme{
		tok("@abc", token.SPELL),
	})
}

func TestSpell_2(t *testing.T) {
	doTest(t, "@abc.efg", []token.Lexeme{
		tok("@abc.efg", token.SPELL),
	})
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestSpell_5(t *testing.T) {
	doErrTest(t, "@abc.efg.hij")
}

func TestString_1(t *testing.T) {
	doTest(t, `""`, []token.Lexeme{
		tok(`""`, token.STRING),
	})
}

func TestString_2(t *testing.T) {
	doTest(t, `"abc"`, []token.Lexeme{
		tok(`"abc"`, token.STRING),
	})
}

func TestString_3(t *testing.T) {
	doTest(t, `"\""`, []token.Lexeme{
		tok(`"\""`, token.STRING),
	})
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
	doTest(t, "123", []token.Lexeme{
		tok("123", token.NUMBER),
	})
}

func TestNumber_2(t *testing.T) {
	doTest(t, "123.456", []token.Lexeme{
		tok("123.456", token.NUMBER),
	})
}

func TestNumber_3(t *testing.T) {
	doErrTest(t, "123.")
}

func TestNumber_4(t *testing.T) {
	doErrTest(t, "123.abc")
}

func TestExist_1(t *testing.T) {
	doTest(t, "?", []token.Lexeme{
		tok("?", token.EXIST),
	})
}

func TestComprehensive_1(t *testing.T) {

	in := `x <- 1 + 2
@Println("x = ", x)`

	tm := &position.TextMarker{}
	genLex := func(v string, tk token.Token) token.Lexeme {
		rng := tm.RangeOf(v)
		tm.Adv(v)
		l := tok(v, tk)
		l.Range = rng
		return l
	}

	exp := []token.Lexeme{
		genLex("x", token.IDENT),
		genLex(" ", token.SPACE),
		genLex("<-", token.ASSIGN),
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
	}

	doTest(t, in, exp)
}
