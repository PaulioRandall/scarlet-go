package sanitiser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/todo/series"
	"github.com/PaulioRandall/scarlet-go/todo/tokentest"
)

func doTest(t *testing.T, in, exp *series.Series) {
	SanitiseAll(in)
	in.JumpToStart()
	tokentest.RequireSeries(t, exp, in)
}

func TestRedundant_1(t *testing.T) {
	in := tokentest.FeignSeries(token.MakeTok(" ", token.SPACE))
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestRedundant_2(t *testing.T) {
	in := tokentest.FeignSeries(token.MakeTok("# Scarlet", token.COMMENT))
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestLeadingTerminators_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok("\n", token.TERMINATOR),
		token.MakeTok(";", token.TERMINATOR),
	)
	exp := tokentest.FeignSeries()
	doTest(t, in, exp)
}

func TestSuccessiveTerminators_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok("x", token.IDENT),
		token.MakeTok("\n", token.TERMINATOR),
		token.MakeTok(";", token.TERMINATOR),
	)
	exp := tokentest.FeignSeries(
		token.MakeTok("x", token.IDENT),
		token.MakeTok("\n", token.TERMINATOR),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterOpener_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("\n", token.NEWLINE),
	)
	exp := tokentest.FeignSeries(
		token.MakeTok("(", token.L_PAREN),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterDelim_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok(",", token.DELIM),
		token.MakeTok("\n", token.NEWLINE),
	)
	exp := tokentest.FeignSeries(
		token.MakeTok(",", token.DELIM),
	)
	doTest(t, in, exp)
}

func TestDelimBeforeRParen_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok(",", token.DELIM),
		token.MakeTok(")", token.R_PAREN),
	)
	exp := tokentest.FeignSeries(
		token.MakeTok(")", token.R_PAREN),
	)
	doTest(t, in, exp)
}

func TestTerminatorBeforeRCurly_1(t *testing.T) {
	in := tokentest.FeignSeries(
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("}", token.R_CURLY),
	)
	exp := tokentest.FeignSeries(
		token.MakeTok("}", token.R_CURLY),
	)
	doTest(t, in, exp)
}

func TestFull_1(t *testing.T) {

	in := tokentest.FeignSeries(
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("@Println", token.SPELL),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(",", token.DELIM),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("\n", token.NEWLINE),
	)

	// @Println(1,1)
	exp := tokentest.FeignSeries(
		token.MakeTok("@Println", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok("\n", token.NEWLINE),
	)

	doTest(t, in, exp)
}

func TestFull_2(t *testing.T) {

	// [true] {
	//   "abc"
	//   "xyz"
	// }
	in := tokentest.FeignSeries(
		token.MakeTok("[", token.L_SQUARE),
		token.MakeTok("true", token.TRUE),
		token.MakeTok("]", token.R_SQUARE),
		token.MakeTok(" ", token.SPACE),
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\t", token.SPACE),
		token.MakeTok(`"abc"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("\t", token.SPACE),
		token.MakeTok(`"xyz"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("}", token.R_CURLY),
	)

	// [true] {"abc"
	// "xyz"}
	exp := tokentest.FeignSeries(
		token.MakeTok("[", token.L_SQUARE),
		token.MakeTok("true", token.TRUE),
		token.MakeTok("]", token.R_SQUARE),
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok(`"abc"`, token.STRING),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok(`"xyz"`, token.STRING),
		token.MakeTok("}", token.R_CURLY),
	)

	doTest(t, in, exp)
}
