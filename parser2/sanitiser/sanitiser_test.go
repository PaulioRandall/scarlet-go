package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/lexeme"
	"github.com/PaulioRandall/scarlet-go/token/series"
	"github.com/PaulioRandall/scarlet-go/token/token"
	"github.com/PaulioRandall/scarlet-go/token/tokentest"
)

func doTest(t *testing.T, in, exp *series.Series) {
	SanitiseAll(in)
	in.JumpToStart()
	tokentest.RequireSeries(t, exp, in)
}

func TestRedundant_1(t *testing.T) {
	in := tokentest.FeignSeries(lexeme.MakeTok(" ", token.SPACE))
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestRedundant_2(t *testing.T) {
	in := tokentest.FeignSeries(lexeme.MakeTok("# Scarlet", token.COMMENT))
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestLeadingTerminators_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok("\n", token.TERMINATOR),
		lexeme.MakeTok(";", token.TERMINATOR),
	)
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestSuccessiveTerminators_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok("\n", token.TERMINATOR),
		lexeme.MakeTok(";", token.TERMINATOR),
	)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok("\n", token.TERMINATOR),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterOpener_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("\n", token.NEWLINE),
	)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok("(", token.L_PAREN),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterDelim_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("\n", token.NEWLINE),
	)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok(",", token.DELIM),
	)
	doTest(t, in, exp)
}

func TestDelimBeforeRParen_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok(")", token.R_PAREN),
	)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok(")", token.R_PAREN),
	)
	doTest(t, in, exp)
}

func TestTerminatorBeforeRCurly_1(t *testing.T) {
	in := tokentest.FeignSeries(
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("}", token.R_CURLY),
	)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok("}", token.R_CURLY),
	)
	doTest(t, in, exp)
}

func TestFull_1(t *testing.T) {

	in := tokentest.FeignSeries(
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("@Println", token.SPELL),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok(")", token.R_PAREN),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("\n", token.NEWLINE),
	)

	// @Println(1,1)
	exp := tokentest.FeignSeries(
		lexeme.MakeTok("@Println", token.SPELL),
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
		lexeme.MakeTok("\n", token.NEWLINE),
	)

	doTest(t, in, exp)
}

func TestFull_2(t *testing.T) {

	// [true] {
	//   "abc"
	//   "xyz"
	// }
	in := tokentest.FeignSeries(
		lexeme.MakeTok("[", token.L_SQUARE),
		lexeme.MakeTok("true", token.TRUE),
		lexeme.MakeTok("]", token.R_SQUARE),
		lexeme.MakeTok(" ", token.SPACE),
		lexeme.MakeTok("{", token.L_CURLY),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("\t", token.SPACE),
		lexeme.MakeTok(`"abc"`, token.STRING),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("\t", token.SPACE),
		lexeme.MakeTok(`"xyz"`, token.STRING),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok("}", token.R_CURLY),
	)

	// [true] {"abc"
	// "xyz"}
	exp := tokentest.FeignSeries(
		lexeme.MakeTok("[", token.L_SQUARE),
		lexeme.MakeTok("true", token.TRUE),
		lexeme.MakeTok("]", token.R_SQUARE),
		lexeme.MakeTok("{", token.L_CURLY),
		lexeme.MakeTok(`"abc"`, token.STRING),
		lexeme.MakeTok("\n", token.NEWLINE),
		lexeme.MakeTok(`"xyz"`, token.STRING),
		lexeme.MakeTok("}", token.R_CURLY),
	)

	doTest(t, in, exp)
}
