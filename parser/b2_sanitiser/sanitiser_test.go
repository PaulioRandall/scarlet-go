package sanitiser2

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/container/conttest"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

func doTest(t *testing.T, in, exp *container.Container) {
	SanitiseAll(in)
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

/*
func Test1_3(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("", lexeme.UNDEFINED),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("", lexeme.UNDEFINED),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test1_4(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("(", lexeme.L_PAREN),
	)

	doTest(t, in, exp)
}

func Test1_5(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
	)

	doTest(t, in, exp)
}

func Test1_6(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(")", lexeme.R_PAREN),
	)

	exp := lextest.Feign(
		lextest.Tok(")", lexeme.R_PAREN),
	)

	doTest(t, in, exp)
}

func Test1_7(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("{", lexeme.L_CURLY),
	)

	doTest(t, in, exp)
}

func Test1_8(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	exp := lextest.Feign(
		lextest.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}

func Test99_1(t *testing.T) {

	in := lextest.Feign(
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	// @Println(1,1)
	exp := lextest.Feign(
		lextest.Tok("@Println", lexeme.SPELL),
		lextest.Tok("(", lexeme.L_PAREN),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(",", lexeme.DELIM),
		lextest.Tok("1", lexeme.NUMBER),
		lextest.Tok(")", lexeme.R_PAREN),
		lextest.Tok("\n", lexeme.NEWLINE),
	)

	doTest(t, in, exp)
}

func Test99_2(t *testing.T) {

	// [true] {
	//   "abc"
	//   "xyz"
	// }
	in := lextest.Feign(
		lextest.Tok("[", lexeme.L_SQUARE),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok(" ", lexeme.SPACE),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.SPACE),
		lextest.Tok(`"abc"`, lexeme.STRING),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("\t", lexeme.SPACE),
		lextest.Tok(`"xyz"`, lexeme.STRING),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	// [true] {"abc"
	// "xyz"}
	exp := lextest.Feign(
		lextest.Tok("[", lexeme.L_SQUARE),
		lextest.Tok("true", lexeme.BOOL),
		lextest.Tok("]", lexeme.R_SQUARE),
		lextest.Tok("{", lexeme.L_CURLY),
		lextest.Tok(`"abc"`, lexeme.STRING),
		lextest.Tok("\n", lexeme.NEWLINE),
		lextest.Tok(`"xyz"`, lexeme.STRING),
		lextest.Tok("}", lexeme.R_CURLY),
	)

	doTest(t, in, exp)
}
*/
