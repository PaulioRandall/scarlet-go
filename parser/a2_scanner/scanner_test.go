package scanner2

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
