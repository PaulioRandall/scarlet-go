package scanner

import (
	"testing"

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

func TestNewline_1(t *testing.T) {
	doTest(t, "\n", []token.Lexeme{
		token.Make("\n", token.NEWLINE, token.Snippet{
			End: token.UTF8Pos{Offset: 1, Line: 1},
		}),
	})
}

func TestNewline_2(t *testing.T) {
	doTest(t, "\r\n", []token.Lexeme{
		token.Make("\r\n", token.NEWLINE, token.Snippet{
			End: token.UTF8Pos{Offset: 2, Line: 1},
		}),
	})
}

func TestSpace_1(t *testing.T) {
	doTest(t, " ", []token.Lexeme{
		token.MakeTok(" ", token.SPACE),
	})
}

func TestSpace_2(t *testing.T) {
	doTest(t, "\t\r\v\f ", []token.Lexeme{
		token.MakeTok("\t\r\v\f ", token.SPACE),
	})
}

func TestComment_1(t *testing.T) {
	doTest(t, "# :)", []token.Lexeme{
		token.MakeTok("# :)", token.COMMENT),
	})
}

func TestBool_1(t *testing.T) {
	doTest(t, "true", []token.Lexeme{
		token.MakeTok("true", token.TRUE),
	})
}

func TestBool_2(t *testing.T) {
	doTest(t, "false", []token.Lexeme{
		token.MakeTok("false", token.FALSE),
	})
}

func TestLoop_1(t *testing.T) {
	doTest(t, "loop", []token.Lexeme{
		token.MakeTok("loop", token.LOOP),
	})
}

func TestIdent_1(t *testing.T) {
	doTest(t, "abc", []token.Lexeme{
		token.MakeTok("abc", token.IDENT),
	})
}

func TestIdent_2(t *testing.T) {
	doTest(t, "abc_xyz", []token.Lexeme{
		token.MakeTok("abc_xyz", token.IDENT),
	})
}

func TestTerminator_1(t *testing.T) {
	doTest(t, ";", []token.Lexeme{
		token.MakeTok(";", token.TERMINATOR),
	})
}

func TestAssign_1(t *testing.T) {
	doTest(t, ":=", []token.Lexeme{
		token.MakeTok(":=", token.ASSIGN),
	})
}

func TestDelim_1(t *testing.T) {
	doTest(t, ",", []token.Lexeme{
		token.MakeTok(",", token.DELIM),
	})
}

func TestLeftParen_1(t *testing.T) {
	doTest(t, "(", []token.Lexeme{
		token.MakeTok("(", token.L_PAREN),
	})
}

func TestRightParen_1(t *testing.T) {
	doTest(t, ")", []token.Lexeme{
		token.MakeTok(")", token.R_PAREN),
	})
}

func TestLeftSquare_1(t *testing.T) {
	doTest(t, "[", []token.Lexeme{
		token.MakeTok("[", token.L_SQUARE),
	})
}

func TestRightSquare_1(t *testing.T) {
	doTest(t, "]", []token.Lexeme{
		token.MakeTok("]", token.R_SQUARE),
	})
}

func TestLeftCurly_1(t *testing.T) {
	doTest(t, "{", []token.Lexeme{
		token.MakeTok("{", token.L_CURLY),
	})
}

func TestRightCurly_1(t *testing.T) {
	doTest(t, "}", []token.Lexeme{
		token.MakeTok("}", token.R_CURLY),
	})
}

func TestVoid_1(t *testing.T) {
	doTest(t, "_", []token.Lexeme{
		token.MakeTok("_", token.VOID),
	})
}

func TestAdd_1(t *testing.T) {
	doTest(t, "+", []token.Lexeme{
		token.MakeTok("+", token.ADD),
	})
}

func TestSub_1(t *testing.T) {
	doTest(t, "-", []token.Lexeme{
		token.MakeTok("-", token.SUB),
	})
}

func TestMul_1(t *testing.T) {
	doTest(t, "*", []token.Lexeme{
		token.MakeTok("*", token.MUL),
	})
}

func TestDiv_1(t *testing.T) {
	doTest(t, "/", []token.Lexeme{
		token.MakeTok("/", token.DIV),
	})
}

func TestRem_1(t *testing.T) {
	doTest(t, "%", []token.Lexeme{
		token.MakeTok("%", token.REM),
	})
}

func TestAnd_1(t *testing.T) {
	doTest(t, "&&", []token.Lexeme{
		token.MakeTok("&&", token.AND),
	})
}

func TestOr_1(t *testing.T) {
	doTest(t, "||", []token.Lexeme{
		token.MakeTok("||", token.OR),
	})
}

func TestLessEqual_1(t *testing.T) {
	doTest(t, "<=", []token.Lexeme{
		token.MakeTok("<=", token.LESS_EQUAL),
	})
}

func TestLess_1(t *testing.T) {
	doTest(t, "<", []token.Lexeme{
		token.MakeTok("<", token.LESS),
	})
}

func TestMoreEqual_1(t *testing.T) {
	doTest(t, ">=", []token.Lexeme{
		token.MakeTok(">=", token.MORE_EQUAL),
	})
}

func TestMore_1(t *testing.T) {
	doTest(t, ">", []token.Lexeme{
		token.MakeTok(">", token.MORE),
	})
}

func TestEqual_1(t *testing.T) {
	doTest(t, "==", []token.Lexeme{
		token.MakeTok("==", token.EQUAL),
	})
}

func TestNotEqual_1(t *testing.T) {
	doTest(t, "!=", []token.Lexeme{
		token.MakeTok("!=", token.NOT_EQUAL),
	})
}

func TestSpell_1(t *testing.T) {
	doTest(t, "@abc", []token.Lexeme{
		token.MakeTok("@abc", token.SPELL),
	})
}

func TestSpell_2(t *testing.T) {
	doTest(t, "@a.b.c", []token.Lexeme{
		token.MakeTok("@a.b.c", token.SPELL),
	})
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestString_1(t *testing.T) {
	doTest(t, `""`, []token.Lexeme{
		token.MakeTok(`""`, token.STRING),
	})
}

func TestString_2(t *testing.T) {
	doTest(t, `"abc"`, []token.Lexeme{
		token.MakeTok(`"abc"`, token.STRING),
	})
}

func TestString_3(t *testing.T) {
	doTest(t, `"\""`, []token.Lexeme{
		token.MakeTok(`"\""`, token.STRING),
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
		token.MakeTok("123", token.NUMBER),
	})
}

func TestNumber_2(t *testing.T) {
	doTest(t, "123.456", []token.Lexeme{
		token.MakeTok("123.456", token.NUMBER),
	})
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

	tm := &token.TextMarker{}
	genLex := func(v string, tk token.Token) token.Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v)
		return token.Make(v, tk, snip)
	}

	exp := []token.Lexeme{
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
	}

	doTest(t, in, exp)
}
