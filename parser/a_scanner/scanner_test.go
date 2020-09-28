package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/container/conttest"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"

	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in string, exp *container.Container) {
	act, e := ScanString(in)
	require.Nil(t, e, "%+v", e)
	conttest.RequireEqual(t, exp.Iterator(), act.Iterator())
}

func doErrTest(t *testing.T, in string) {
	_, e := ScanString(in)
	require.NotNil(t, e, "Expected an error for input %q", in)
}

func TestBadToken(t *testing.T) {
	doErrTest(t, "¬")
}

func TestNewline_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("\n", lexeme.NEWLINE))
	doTest(t, "\n", exp)
}

func TestNewline_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("\r\n", lexeme.NEWLINE))
	doTest(t, "\r\n", exp)
}

func TestSpace_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(" ", lexeme.SPACE))
	doTest(t, " ", exp)
}

func TestSpace_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("\t\r\v\f ", lexeme.SPACE))
	doTest(t, "\t\r\v\f ", exp)
}

func TestComment_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("# :)", lexeme.COMMENT))
	doTest(t, "# :)", exp)
}

func TestBool_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("true", lexeme.BOOL))
	doTest(t, "true", exp)
}

func TestBool_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("false", lexeme.BOOL))
	doTest(t, "false", exp)
}

func TestLoop_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("loop", lexeme.LOOP))
	doTest(t, "loop", exp)
}

func TestIdent_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("abc", lexeme.IDENT))
	doTest(t, "abc", exp)
}

func TestIdent_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("abc_xyz", lexeme.IDENT))
	doTest(t, "abc_xyz", exp)
}

func TestTerminator_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(";", lexeme.TERMINATOR))
	doTest(t, ";", exp)
}

func TestAssign_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(":=", lexeme.ASSIGN))
	doTest(t, ":=", exp)
}

func TestDelim_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(",", lexeme.DELIM))
	doTest(t, ",", exp)
}

func TestLeftParen_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("(", lexeme.L_PAREN))
	doTest(t, "(", exp)
}

func TestRightParen_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(")", lexeme.R_PAREN))
	doTest(t, ")", exp)
}

func TestLeftSquare_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("[", lexeme.L_SQUARE))
	doTest(t, "[", exp)
}

func TestRightSquare_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("]", lexeme.R_SQUARE))
	doTest(t, "]", exp)
}

func TestLeftCurly_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("{", lexeme.L_CURLY))
	doTest(t, "{", exp)
}

func TestRightCurly_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("}", lexeme.R_CURLY))
	doTest(t, "}", exp)
}

func TestVoid_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("_", lexeme.VOID))
	doTest(t, "_", exp)
}

func TestAdd_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("+", lexeme.ADD))
	doTest(t, "+", exp)
}

func TestSub_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("-", lexeme.SUB))
	doTest(t, "-", exp)
}

func TestMul_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("*", lexeme.MUL))
	doTest(t, "*", exp)
}

func TestDiv_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("/", lexeme.DIV))
	doTest(t, "/", exp)
}

func TestRem_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("%", lexeme.REM))
	doTest(t, "%", exp)
}

func TestAnd_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("&&", lexeme.AND))
	doTest(t, "&&", exp)
}

func TestOr_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("||", lexeme.OR))
	doTest(t, "||", exp)
}

func TestLessEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("<=", lexeme.LESS_EQUAL))
	doTest(t, "<=", exp)
}

func TestLess_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("<", lexeme.LESS))
	doTest(t, "<", exp)
}

func TestMoreEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(">=", lexeme.MORE_EQUAL))
	doTest(t, ">=", exp)
}

func TestMore_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(">", lexeme.MORE))
	doTest(t, ">", exp)
}

func TestEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("==", lexeme.EQUAL))
	doTest(t, "==", exp)
}

func TestNotEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("!=", lexeme.NOT_EQUAL))
	doTest(t, "!=", exp)
}

func TestSpell_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("@abc", lexeme.SPELL))
	doTest(t, "@abc", exp)
}

func TestSpell_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("@a.b.c", lexeme.SPELL))
	doTest(t, "@a.b.c", exp)
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestString_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`""`, lexeme.STRING))
	doTest(t, `""`, exp)
}

func TestString_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`"abc"`, lexeme.STRING))
	doTest(t, `"abc"`, exp)
}

func TestString_3(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`"\""`, lexeme.STRING))
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
	exp := conttest.Feign(lexeme.Tok("123", lexeme.NUMBER))
	doTest(t, "123", exp)
}

func TestNumber_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("123.456", lexeme.NUMBER))
	doTest(t, "123.456", exp)
}

func TestNumber_3(t *testing.T) {
	doErrTest(t, "123.")
}

func TestNumber_4(t *testing.T) {
	doErrTest(t, "123.abc")
}
