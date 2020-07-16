package check

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in []token.Token) {

	acts, e := CheckAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, in, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {
	_, e := CheckAll(in)
	require.NotNil(t, e, "Expected an error")
}

func tok(raw string, props ...Prop) token.Tok {
	return token.Tok{
		RawProps: props,
		RawStr:   raw,
		ColEnd:   len(raw),
	}
}

func Test1_1(t *testing.T) {

	// WHEN checking a spell with no arguments
	// THEN no errors should be returned
	// @Println()
	in := []token.Token{
		tok("@Print", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in)
}

func Test1_2(t *testing.T) {

	// WHEN checking a spell with one argument
	// THEN no errors should be returned
	// @Println(x)
	in := []token.Token{
		tok("@Print", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in)
}

func Test1_3(t *testing.T) {

	// WHEN checking a spell with multiple arguments
	// THEN no errors should be returned
	// @Println(x, 1, true)
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok(",", PR_SEPARATOR),
		tok("1", PR_TERM),
		tok(",", PR_SEPARATOR),
		tok("true", PR_TERM),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in)
}

func Test2_1(t *testing.T) {

	// WHEN checking a spell with missing opening parenthesis
	// THEN an error should be returned
	// @Println)
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doErrorTest(t, in)
}

func Test2_2(t *testing.T) {

	// WHEN checking a spell with missing closing parenthesis
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("\n", PR_TERMINATOR),
	}

	doErrorTest(t, in)
}

func Test2_3(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(",", PR_SEPARATOR),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doErrorTest(t, in)
}

func Test2_4(t *testing.T) {

	// WHEN checking a spell with a stray value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok(",", PR_SEPARATOR),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doErrorTest(t, in)
}

func Test2_5(t *testing.T) {

	// WHEN checking a spell with a missing value delimiter
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok("y", PR_TERM),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	doErrorTest(t, in)
}

func Test2_6(t *testing.T) {

	// WHEN checking a spell with a missing final terminator
	// THEN an error should be returned
	// @Println(
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	doErrorTest(t, in)
}
