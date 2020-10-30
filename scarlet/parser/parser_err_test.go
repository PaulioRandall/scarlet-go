package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"

	"github.com/stretchr/testify/require"
)

func doErrTest(t *testing.T, in ...token.Lexeme) {
	_, e := ParseAll(in)
	require.NotNil(t, e, "Expected parse error")
}

func TestParse_FailAssign_1(t *testing.T) {
	// x 1
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailAssign_2(t *testing.T) {
	// x :=
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
	)
}

func TestParse_FailAssign_3(t *testing.T) {
	// := 1
	doErrTest(t,
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailAssign_4(t *testing.T) {
	// x x
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok("x", token.IDENT),
	)
}

func TestParse_FailAssign_5(t *testing.T) {
	// x, y, :=
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok(":=", token.ASSIGN),
	)
}

func TestParse_FailAssign_6(t *testing.T) {
	// x, y := 1,
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
	)
}

func TestParse_FailAssign_7(t *testing.T) {
	// x, y := 1, 2,
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(",", token.DELIM),
	)
}

func TestParse_FailAssign_8(t *testing.T) {
	// x, y, z := 1, 2
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("z", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("2", token.NUMBER),
	)
}

func TestParse_FailAssign_9(t *testing.T) {
	// x, y := 1, 2, 3
	doErrTest(t,
		token.MakeTok("x", token.IDENT),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("3", token.NUMBER),
	)
}

func TestParse_FailBinaryExpr_1(t *testing.T) {
	// 1 + + 2
	doErrTest(t,
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
	)
}

func TestParse_FailParenExpr_1(t *testing.T) {
	// ()
	doErrTest(t,
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_2(t *testing.T) {
	// (1
	doErrTest(t,
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
	)
}

func TestParse_FailParenExpr_3(t *testing.T) {
	// ((1 + 2)
	doErrTest(t,
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_4(t *testing.T) {
	// (1 +)
	doErrTest(t,
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok(")", token.R_PAREN),
	)
}

func TestParse_FailParenExpr_5(t *testing.T) {
	// (+ 2)
	doErrTest(t,
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
	)
}
