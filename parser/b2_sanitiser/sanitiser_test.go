package sanitiser2

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/container/conttest"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

func doTest(t *testing.T, in, exp *container.Container) {
	SanitiseAll(in.Iterator())
	conttest.RequireEqual(t, exp.Iterator(), in.Iterator())
}

func TestRedundant_1(t *testing.T) {
	in := conttest.Feign(lexeme.Tok(" ", lexeme.SPACE))
	exp := conttest.Feign()
	doTest(t, in, exp)
}

func TestRedundant_2(t *testing.T) {
	in := conttest.Feign(lexeme.Tok("# Scarlet", lexeme.COMMENT))
	exp := conttest.Feign()
	doTest(t, in, exp)
}

func TestLeadingTerminators_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok("\n", lexeme.TERMINATOR),
		lexeme.Tok(";", lexeme.TERMINATOR),
	)
	exp := conttest.Feign()
	doTest(t, in, exp)
}

func TestSuccessiveTerminators_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok("x", lexeme.IDENT),
		lexeme.Tok("\n", lexeme.TERMINATOR),
		lexeme.Tok(";", lexeme.TERMINATOR),
	)
	exp := conttest.Feign(
		lexeme.Tok("x", lexeme.IDENT),
		lexeme.Tok("\n", lexeme.TERMINATOR),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterOpener_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok("(", lexeme.L_PAREN),
		lexeme.Tok("\n", lexeme.NEWLINE),
	)
	exp := conttest.Feign(
		lexeme.Tok("(", lexeme.L_PAREN),
	)
	doTest(t, in, exp)
}

func TestNewlineAfterDelim_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok(",", lexeme.DELIM),
		lexeme.Tok("\n", lexeme.NEWLINE),
	)
	exp := conttest.Feign(
		lexeme.Tok(",", lexeme.DELIM),
	)
	doTest(t, in, exp)
}

func TestDelimBeforeRParen_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok(",", lexeme.DELIM),
		lexeme.Tok(")", lexeme.R_PAREN),
	)
	exp := conttest.Feign(
		lexeme.Tok(")", lexeme.R_PAREN),
	)
	doTest(t, in, exp)
}

func TestTerminatorBeforeRCurly_1(t *testing.T) {
	in := conttest.Feign(
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("}", lexeme.R_CURLY),
	)
	exp := conttest.Feign(
		lexeme.Tok("}", lexeme.R_CURLY),
	)
	doTest(t, in, exp)
}

func TestFull_1(t *testing.T) {

	in := conttest.Feign(
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("@Println", lexeme.SPELL),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("(", lexeme.L_PAREN),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("1", lexeme.NUMBER),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok(",", lexeme.DELIM),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("1", lexeme.NUMBER),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok(",", lexeme.DELIM),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok(")", lexeme.R_PAREN),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("\n", lexeme.NEWLINE),
	)

	// @Println(1,1)
	exp := conttest.Feign(
		lexeme.Tok("@Println", lexeme.SPELL),
		lexeme.Tok("(", lexeme.L_PAREN),
		lexeme.Tok("1", lexeme.NUMBER),
		lexeme.Tok(",", lexeme.DELIM),
		lexeme.Tok("1", lexeme.NUMBER),
		lexeme.Tok(")", lexeme.R_PAREN),
		lexeme.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func TestFull_2(t *testing.T) {

	// [true] {
	//   "abc"
	//   "xyz"
	// }
	in := conttest.Feign(
		lexeme.Tok("[", lexeme.L_SQUARE),
		lexeme.Tok("true", lexeme.BOOL),
		lexeme.Tok("]", lexeme.R_SQUARE),
		lexeme.Tok(" ", lexeme.SPACE),
		lexeme.Tok("{", lexeme.L_CURLY),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("\t", lexeme.SPACE),
		lexeme.Tok(`"abc"`, lexeme.STRING),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("\t", lexeme.SPACE),
		lexeme.Tok(`"xyz"`, lexeme.STRING),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok("}", lexeme.R_CURLY),
	)

	// [true] {"abc"
	// "xyz"}
	exp := conttest.Feign(
		lexeme.Tok("[", lexeme.L_SQUARE),
		lexeme.Tok("true", lexeme.BOOL),
		lexeme.Tok("]", lexeme.R_SQUARE),
		lexeme.Tok("{", lexeme.L_CURLY),
		lexeme.Tok(`"abc"`, lexeme.STRING),
		lexeme.Tok("\n", lexeme.NEWLINE),
		lexeme.Tok(`"xyz"`, lexeme.STRING),
		lexeme.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}
