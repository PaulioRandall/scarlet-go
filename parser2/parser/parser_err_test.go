package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/token"
	"github.com/PaulioRandall/scarlet-go/token/tokentest"

	"github.com/stretchr/testify/require"
)

func doErrTest(t *testing.T, in ...lexeme.Lexeme) {
	tokenItr := tokentest.FeignSeries(in...)
	_, e := ParseAll(tokenItr)
	require.NotNil(t, e, "Expected parse error")
}

func TestParse_FailAssign_1(t *testing.T) {
	// x 1
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailAssign_2(t *testing.T) {
	// x :=
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
	)
}

func TestParse_FailAssign_3(t *testing.T) {
	// := 1
	doErrTest(t,
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailAssign_4(t *testing.T) {
	// x x
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok("x", token.IDENT),
	)
}

func TestParse_FailAssign_5(t *testing.T) {
	// x, y, :=
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok(":=", token.ASSIGN),
	)
}

func TestParse_FailAssign_6(t *testing.T) {
	// x, y := 1,
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
	)
}

func TestParse_FailAssign_7(t *testing.T) {
	// x, y := 1, 2,
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
	)
}

func TestParse_FailAssign_8(t *testing.T) {
	// x, y, z := 1, 2
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("z", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("2", token.NUMBER),
	)
}

func TestParse_FailAssign_9(t *testing.T) {
	// x, y := 1, 2, 3
	doErrTest(t,
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("3", token.NUMBER),
	)
}

func TestParse_FailBinaryExpr_1(t *testing.T) {
	// 1 + + 2
	doErrTest(t,
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
	)
}

func TestParse_FailParenExpr_1(t *testing.T) {
	// ()
	doErrTest(t,
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_2(t *testing.T) {
	// (1
	doErrTest(t,
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailParenExpr_3(t *testing.T) {
	// ((1 + 2)
	doErrTest(t,
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_4(t *testing.T) {
	// (1 +)
	doErrTest(t,
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_5(t *testing.T) {
	// (+ 2)
	doErrTest(t,
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
	)
}
