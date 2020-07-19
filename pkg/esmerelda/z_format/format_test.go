package format

import (
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func tok(raw string, props ...Prop) token.Tok {
	return token.Tok{
		RawStr:   raw,
		RawProps: props,
	}
}

func Test1_1(t *testing.T) {

	given := []token.Token{
		tok("\r\n", PR_NEWLINE),
	}

	exps := []token.Token{
		tok("\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test1_2(t *testing.T) {

	given := []token.Token{
		tok("\r\n", PR_NEWLINE),
	}

	exps := []token.Token{
		tok("\r\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\r\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_1(t *testing.T) {

	given := []token.Token{
		tok("\n", PR_NEWLINE),
		tok(" ", PR_WHITESPACE),
	}

	exps := []token.Token{
		tok("\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_2(t *testing.T) {

	given := []token.Token{
		tok(" ", PR_WHITESPACE),
		tok("\n", PR_NEWLINE),
	}

	exps := []token.Token{
		tok("\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_3(t *testing.T) {

	given := []token.Token{
		tok("@Println", PR_SPELL),
		tok(" ", PR_WHITESPACE),
		tok("(", PR_OPENER),
	}

	exps := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_OPENER),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_4(t *testing.T) {

	given := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_OPENER),
		tok(" ", PR_WHITESPACE),
	}

	exps := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_OPENER),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_5(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok(" ", PR_WHITESPACE),
		tok(",", PR_SEPARATOR),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_6(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_7(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("\n", PR_NEWLINE),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok("\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_8(t *testing.T) {

	given := []token.Token{
		tok(" ", PR_WHITESPACE),
		tok(")", PR_CLOSER),
	}

	exps := []token.Token{
		tok(")", PR_CLOSER),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test2_9(t *testing.T) {

	// " @Println ( 1 , 1 , \n 1 ) \n "
	given := []token.Token{
		tok(" ", PR_WHITESPACE),
		tok("@Println", PR_SPELL),
		tok(" ", PR_WHITESPACE),
		tok("(", PR_OPENER),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
		tok(" ", PR_WHITESPACE),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
		tok(" ", PR_WHITESPACE),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("\n", PR_NEWLINE),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
		tok(" ", PR_WHITESPACE),
		tok(")", PR_CLOSER),
		tok(" ", PR_WHITESPACE),
		tok("\n", PR_NEWLINE),
		tok(" ", PR_WHITESPACE),
	}

	// "@Println(1, 1,\n1)\n"
	exps := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_OPENER),
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok("\n", PR_NEWLINE),
		tok("1", PR_LITERAL),
		tok(")", PR_CLOSER),
		tok("\n", PR_NEWLINE),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test3_1(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok("   ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test3_2(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok("\t", PR_WHITESPACE),
		tok("1", PR_LITERAL),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok(",", PR_SEPARATOR),
		tok(" ", PR_WHITESPACE),
		tok("1", PR_LITERAL),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test4_1(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok("\n", PR_NEWLINE),
		tok("\n", PR_NEWLINE),
		tok("1", PR_LITERAL),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok("\n", PR_NEWLINE),
		tok("\n", PR_NEWLINE),
		tok("1", PR_LITERAL),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}

func Test4_2(t *testing.T) {

	given := []token.Token{
		tok("1", PR_LITERAL),
		tok("\n", PR_NEWLINE),
		tok("\n", PR_NEWLINE),
		tok("\n", PR_NEWLINE),
		tok("1", PR_LITERAL),
	}

	exps := []token.Token{
		tok("1", PR_LITERAL),
		tok("\n", PR_NEWLINE),
		tok("\n", PR_NEWLINE),
		tok("1", PR_LITERAL),
	}

	acts := FormatAll(given, "\n")
	testutils.RequireTokenSlice(t, exps, acts)
}
