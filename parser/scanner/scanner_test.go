package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/position"
	"github.com/PaulioRandall/scarlet-go/token/token"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp []lexeme.Lexeme) {
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
	doTest(t, "\n", []lexeme.Lexeme{
		lexeme.Make("\n", token.NEWLINE, position.Snippet{
			End: position.UTF8Pos{Offset: 1, Line: 1},
		}),
	})
}

func TestNewline_2(t *testing.T) {
	doTest(t, "\r\n", []lexeme.Lexeme{
		lexeme.Make("\r\n", token.NEWLINE, position.Snippet{
			End: position.UTF8Pos{Offset: 2, Line: 1},
		}),
	})
}

func TestSpace_1(t *testing.T) {
	doTest(t, " ", []lexeme.Lexeme{
		lexeme.MakeTok(" ", token.SPACE),
	})
}

func TestSpace_2(t *testing.T) {
	doTest(t, "\t\r\v\f ", []lexeme.Lexeme{
		lexeme.MakeTok("\t\r\v\f ", token.SPACE),
	})
}

func TestComment_1(t *testing.T) {
	doTest(t, "# :)", []lexeme.Lexeme{
		lexeme.MakeTok("# :)", token.COMMENT),
	})
}

func TestBool_1(t *testing.T) {
	doTest(t, "true", []lexeme.Lexeme{
		lexeme.MakeTok("true", token.TRUE),
	})
}

func TestBool_2(t *testing.T) {
	doTest(t, "false", []lexeme.Lexeme{
		lexeme.MakeTok("false", token.FALSE),
	})
}

func TestLoop_1(t *testing.T) {
	doTest(t, "loop", []lexeme.Lexeme{
		lexeme.MakeTok("loop", token.LOOP),
	})
}

func TestIdent_1(t *testing.T) {
	doTest(t, "abc", []lexeme.Lexeme{
		lexeme.MakeTok("abc", token.IDENT),
	})
}

func TestIdent_2(t *testing.T) {
	doTest(t, "abc_xyz", []lexeme.Lexeme{
		lexeme.MakeTok("abc_xyz", token.IDENT),
	})
}

func TestTerminator_1(t *testing.T) {
	doTest(t, ";", []lexeme.Lexeme{
		lexeme.MakeTok(";", token.TERMINATOR),
	})
}

func TestAssign_1(t *testing.T) {
	doTest(t, ":=", []lexeme.Lexeme{
		lexeme.MakeTok(":=", token.ASSIGN),
	})
}

func TestDelim_1(t *testing.T) {
	doTest(t, ",", []lexeme.Lexeme{
		lexeme.MakeTok(",", token.DELIM),
	})
}

func TestLeftParen_1(t *testing.T) {
	doTest(t, "(", []lexeme.Lexeme{
		lexeme.MakeTok("(", token.L_PAREN),
	})
}

func TestRightParen_1(t *testing.T) {
	doTest(t, ")", []lexeme.Lexeme{
		lexeme.MakeTok(")", token.R_PAREN),
	})
}

func TestLeftSquare_1(t *testing.T) {
	doTest(t, "[", []lexeme.Lexeme{
		lexeme.MakeTok("[", token.L_SQUARE),
	})
}

func TestRightSquare_1(t *testing.T) {
	doTest(t, "]", []lexeme.Lexeme{
		lexeme.MakeTok("]", token.R_SQUARE),
	})
}

func TestLeftCurly_1(t *testing.T) {
	doTest(t, "{", []lexeme.Lexeme{
		lexeme.MakeTok("{", token.L_CURLY),
	})
}

func TestRightCurly_1(t *testing.T) {
	doTest(t, "}", []lexeme.Lexeme{
		lexeme.MakeTok("}", token.R_CURLY),
	})
}

func TestVoid_1(t *testing.T) {
	doTest(t, "_", []lexeme.Lexeme{
		lexeme.MakeTok("_", token.VOID),
	})
}

func TestAdd_1(t *testing.T) {
	doTest(t, "+", []lexeme.Lexeme{
		lexeme.MakeTok("+", token.ADD),
	})
}

func TestSub_1(t *testing.T) {
	doTest(t, "-", []lexeme.Lexeme{
		lexeme.MakeTok("-", token.SUB),
	})
}

func TestMul_1(t *testing.T) {
	doTest(t, "*", []lexeme.Lexeme{
		lexeme.MakeTok("*", token.MUL),
	})
}

func TestDiv_1(t *testing.T) {
	doTest(t, "/", []lexeme.Lexeme{
		lexeme.MakeTok("/", token.DIV),
	})
}

func TestRem_1(t *testing.T) {
	doTest(t, "%", []lexeme.Lexeme{
		lexeme.MakeTok("%", token.REM),
	})
}

func TestAnd_1(t *testing.T) {
	doTest(t, "&&", []lexeme.Lexeme{
		lexeme.MakeTok("&&", token.AND),
	})
}

func TestOr_1(t *testing.T) {
	doTest(t, "||", []lexeme.Lexeme{
		lexeme.MakeTok("||", token.OR),
	})
}

func TestLessEqual_1(t *testing.T) {
	doTest(t, "<=", []lexeme.Lexeme{
		lexeme.MakeTok("<=", token.LESS_EQUAL),
	})
}

func TestLess_1(t *testing.T) {
	doTest(t, "<", []lexeme.Lexeme{
		lexeme.MakeTok("<", token.LESS),
	})
}

func TestMoreEqual_1(t *testing.T) {
	doTest(t, ">=", []lexeme.Lexeme{
		lexeme.MakeTok(">=", token.MORE_EQUAL),
	})
}

func TestMore_1(t *testing.T) {
	doTest(t, ">", []lexeme.Lexeme{
		lexeme.MakeTok(">", token.MORE),
	})
}

func TestEqual_1(t *testing.T) {
	doTest(t, "==", []lexeme.Lexeme{
		lexeme.MakeTok("==", token.EQUAL),
	})
}

func TestNotEqual_1(t *testing.T) {
	doTest(t, "!=", []lexeme.Lexeme{
		lexeme.MakeTok("!=", token.NOT_EQUAL),
	})
}

func TestSpell_1(t *testing.T) {
	doTest(t, "@abc", []lexeme.Lexeme{
		lexeme.MakeTok("@abc", token.SPELL),
	})
}

func TestSpell_2(t *testing.T) {
	doTest(t, "@a.b.c", []lexeme.Lexeme{
		lexeme.MakeTok("@a.b.c", token.SPELL),
	})
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestString_1(t *testing.T) {
	doTest(t, `""`, []lexeme.Lexeme{
		lexeme.MakeTok(`""`, token.STRING),
	})
}

func TestString_2(t *testing.T) {
	doTest(t, `"abc"`, []lexeme.Lexeme{
		lexeme.MakeTok(`"abc"`, token.STRING),
	})
}

func TestString_3(t *testing.T) {
	doTest(t, `"\""`, []lexeme.Lexeme{
		lexeme.MakeTok(`"\""`, token.STRING),
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
	doTest(t, "123", []lexeme.Lexeme{
		lexeme.MakeTok("123", token.NUMBER),
	})
}

func TestNumber_2(t *testing.T) {
	doTest(t, "123.456", []lexeme.Lexeme{
		lexeme.MakeTok("123.456", token.NUMBER),
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

	tm := &position.TextMarker{}
	genLex := func(v string, tk token.Token) lexeme.Lexeme {
		snip := tm.Snippet(v)
		tm.Advance(v)
		return lexeme.Make(v, tk, snip)
	}

	exp := []lexeme.Lexeme{
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
