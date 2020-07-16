package shunt

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exps []token.Token) {

	acts, e := ShuntAll(in)
	if e != nil {
		require.Nil(t, e, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func doErrorTest(t *testing.T, in []token.Token) {
	_, e := ShuntAll(in)
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

	// WHEN refixing a spell with no arguments
	// @Println()
	in := []token.Token{
		tok("@Print", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	// THEN parenthesis are removed
	// AND magic token indicating parameter start inserted before spell
	exp := []token.Token{
		tok("(", PR_PARENTHESIS, PR_OPENER, PR_PARAMETERS),
		tok("@Print", PR_SPELL),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in, exp)
}

func Test1_2(t *testing.T) {

	// WHEN refixing a spell with one argument
	// @Println(x)
	in := []token.Token{
		tok("@Print", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	// THEN parenthesis are removed
	// AND the argument is placed before the spell
	// AND magic token indicating parameter start inserted before the argument
	exp := []token.Token{
		tok("(", PR_PARENTHESIS, PR_OPENER, PR_PARAMETERS),
		tok("x", PR_TERM),
		tok("@Print", PR_SPELL),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in, exp)
}

func Test1_3(t *testing.T) {

	// WHEN refixing a spell with multiple arguments
	// @Println(x, y, z)
	in := []token.Token{
		tok("@Println", PR_SPELL),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("x", PR_TERM),
		tok(",", PR_SEPARATOR),
		tok("y", PR_TERM),
		tok(",", PR_SEPARATOR),
		tok("z", PR_TERM),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("\n", PR_TERMINATOR),
	}

	// THEN parenthesis are removed
	// AND the delimiters are removed
	// AND the arguments are placed before the spell
	// AND magic token indicating parameter start inserted before the arguments
	exp := []token.Token{
		tok("(", PR_PARENTHESIS, PR_OPENER, PR_PARAMETERS),
		tok("x", PR_TERM),
		tok("y", PR_TERM),
		tok("z", PR_TERM),
		tok("@Println", PR_SPELL),
		tok("\n", PR_TERMINATOR),
	}

	doTest(t, in, exp)
}
