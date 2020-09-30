package scanner

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token2/container"
	"github.com/PaulioRandall/scarlet-go/token2/container/conttest"
	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/token"

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
	exp := conttest.Feign(lexeme.Tok("\n", token.NEWLINE))
	doTest(t, "\n", exp)
}

func TestNewline_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("\r\n", token.NEWLINE))
	doTest(t, "\r\n", exp)
}

func TestSpace_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(" ", token.SPACE))
	doTest(t, " ", exp)
}

func TestSpace_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("\t\r\v\f ", token.SPACE))
	doTest(t, "\t\r\v\f ", exp)
}

func TestComment_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("# :)", token.COMMENT))
	doTest(t, "# :)", exp)
}

func TestBool_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("true", token.TRUE))
	doTest(t, "true", exp)
}

func TestBool_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("false", token.FALSE))
	doTest(t, "false", exp)
}

func TestLoop_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("loop", token.LOOP))
	doTest(t, "loop", exp)
}

func TestIdent_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("abc", token.IDENT))
	doTest(t, "abc", exp)
}

func TestIdent_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("abc_xyz", token.IDENT))
	doTest(t, "abc_xyz", exp)
}

func TestTerminator_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(";", token.TERMINATOR))
	doTest(t, ";", exp)
}

func TestAssign_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(":=", token.ASSIGN))
	doTest(t, ":=", exp)
}

func TestDelim_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(",", token.DELIM))
	doTest(t, ",", exp)
}

func TestLeftParen_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("(", token.L_PAREN))
	doTest(t, "(", exp)
}

func TestRightParen_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(")", token.R_PAREN))
	doTest(t, ")", exp)
}

func TestLeftSquare_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("[", token.L_SQUARE))
	doTest(t, "[", exp)
}

func TestRightSquare_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("]", token.R_SQUARE))
	doTest(t, "]", exp)
}

func TestLeftCurly_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("{", token.L_CURLY))
	doTest(t, "{", exp)
}

func TestRightCurly_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("}", token.R_CURLY))
	doTest(t, "}", exp)
}

func TestVoid_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("_", token.VOID))
	doTest(t, "_", exp)
}

func TestAdd_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("+", token.ADD))
	doTest(t, "+", exp)
}

func TestSub_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("-", token.SUB))
	doTest(t, "-", exp)
}

func TestMul_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("*", token.MUL))
	doTest(t, "*", exp)
}

func TestDiv_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("/", token.DIV))
	doTest(t, "/", exp)
}

func TestRem_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("%", token.REM))
	doTest(t, "%", exp)
}

func TestAnd_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("&&", token.AND))
	doTest(t, "&&", exp)
}

func TestOr_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("||", token.OR))
	doTest(t, "||", exp)
}

func TestLessEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("<=", token.LESS_EQUAL))
	doTest(t, "<=", exp)
}

func TestLess_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("<", token.LESS))
	doTest(t, "<", exp)
}

func TestMoreEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(">=", token.MORE_EQUAL))
	doTest(t, ">=", exp)
}

func TestMore_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(">", token.MORE))
	doTest(t, ">", exp)
}

func TestEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("==", token.EQUAL))
	doTest(t, "==", exp)
}

func TestNotEqual_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("!=", token.NOT_EQUAL))
	doTest(t, "!=", exp)
}

func TestSpell_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("@abc", token.SPELL))
	doTest(t, "@abc", exp)
}

func TestSpell_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("@a.b.c", token.SPELL))
	doTest(t, "@a.b.c", exp)
}

func TestSpell_3(t *testing.T) {
	doErrTest(t, "@")
}

func TestSpell_4(t *testing.T) {
	doErrTest(t, "@abc.")
}

func TestString_1(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`""`, token.STRING))
	doTest(t, `""`, exp)
}

func TestString_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`"abc"`, token.STRING))
	doTest(t, `"abc"`, exp)
}

func TestString_3(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok(`"\""`, token.STRING))
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
	exp := conttest.Feign(lexeme.Tok("123", token.NUMBER))
	doTest(t, "123", exp)
}

func TestNumber_2(t *testing.T) {
	exp := conttest.Feign(lexeme.Tok("123.456", token.NUMBER))
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

	exp := conttest.Feign(
		lexeme.New("x", token.IDENT, 0, 0),
		lexeme.New(" ", token.SPACE, 0, 1),
		lexeme.New(":=", token.ASSIGN, 0, 2),
		lexeme.New(" ", token.SPACE, 0, 4),
		lexeme.New("1", token.NUMBER, 0, 5),
		lexeme.New(" ", token.SPACE, 0, 6),
		lexeme.New("+", token.ADD, 0, 7),
		lexeme.New(" ", token.SPACE, 0, 8),
		lexeme.New("2", token.NUMBER, 0, 9),
		lexeme.New("\n", token.NEWLINE, 0, 10),
		lexeme.New("@Println", token.SPELL, 1, 0),
		lexeme.New("(", token.L_PAREN, 1, 8),
		lexeme.New(`"x = "`, token.STRING, 1, 9),
		lexeme.New(",", token.DELIM, 1, 15),
		lexeme.New(" ", token.SPACE, 1, 16),
		lexeme.New("x", token.IDENT, 1, 17),
		lexeme.New(")", token.R_PAREN, 1, 18),
	)

	doTest(t, in, exp)
}
