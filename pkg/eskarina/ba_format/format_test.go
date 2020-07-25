package format

import (
	"testing"

	//"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme/lextest"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func Test1_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign()

	act := trimLeadingSpace(given)
	lextest.Equal(t, exp, act)
}

func Test1_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := trimLeadingSpace(given)
	lextest.Equal(t, exp, act)
}

func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test2_4(t *testing.T) {

	given := lextest.Feign(
		lextest.Lex(2, 4, ",", prop.PR_SEPARATOR),
		lextest.Lex(2, 5, "1", prop.PR_LITERAL),
	)

	exp := lextest.Feign(
		lextest.Lex(2, 4, ",", prop.PR_SEPARATOR),
		lextest.Lex(2, 5, " ", prop.PR_REDUNDANT, prop.PR_WHITESPACE),
		lextest.Lex(2, 5, "1", prop.PR_LITERAL),
	)

	act := insertSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	exp := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("(", prop.PR_OPENER),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

func Test3_3(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
	)

	exp := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
	)

	act := trimSpaces(given)
	lextest.Equal(t, exp, act)
}

/*
func Test1_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\r\n", prop.PR_NEWLINE),
	)

	exps := lextest.Feign(
		lextest.Tok("\r\n", prop.PR_NEWLINE),
	)

	acts := FormatAll(given, "\r\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test2_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		)

	exps := lextest.Feign(
		lextest.Tok("\n", prop.PR_NEWLINE),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_5(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_SEPARATOR),
		)

	exps := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*

/*
func Test2_8(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(")", prop.PR_CLOSER),
		)

	exps := lextest.Feign(
		lextest.Tok(")", prop.PR_CLOSER),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test2_9(t *testing.T) {

	// " @Println ( 1 , 1 , \n 1 ) \n "
	given := lextest.Feign(
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok(")", prop.PR_CLOSER),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		)

	// "@Println(1, 1,\n1)\n"
	exps := lextest.Feign(
		lextest.Tok("@Println", prop.PR_SPELL),
		lextest.Tok("(", prop.PR_OPENER),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(")", prop.PR_CLOSER),
		lextest.Tok("\n", prop.PR_NEWLINE),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test3_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("   ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	exps := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test3_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok("\t", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	exps := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok(",", prop.PR_SEPARATOR),
		lextest.Tok(" ", prop.PR_WHITESPACE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test4_1(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	exps := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
/*
func Test4_2(t *testing.T) {

	given := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	exps := lextest.Feign(
		lextest.Tok("1", prop.PR_LITERAL),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("\n", prop.PR_NEWLINE),
		lextest.Tok("1", prop.PR_LITERAL),
		)

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
*/
