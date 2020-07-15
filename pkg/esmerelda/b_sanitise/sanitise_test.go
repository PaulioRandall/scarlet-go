package sanitise

import (
	"fmt"
	"testing"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exps []token.Token) {

	acts, e := SanitiseAll(in)
	if e != nil {
		require.Nil(t, fmt.Sprintf("%+v", e))
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func tok(raw string, props ...Prop) token.Tok {
	return token.Tok{
		RawProps: props,
		RawStr:   raw,
		ColEnd:   len(raw),
	}
}

func Test1_1(t *testing.T) {

	// WHEN sanitising a statement containing redudant whitespace and comments
	// @Println (  )
	in := []token.Token{
		tok("@Print"),
		tok(" ", PR_REDUNDANT),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("  ", PR_REDUNDANT),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
		tok("# abc", PR_REDUNDANT),
	}

	// THEN the whitespace is removed
	exp := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after opening parenthesis
	// @Println(
	// )
	in := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("\n", PR_NEWLINE),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// @Println(1,
	// 1)
	in := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("1"),
		tok(",", PR_SEPARATOR),
		tok("\n", PR_NEWLINE),
		tok("1"),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("1"),
		tok(",", PR_SEPARATOR),
		tok("1"),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// AND the next line only contains the closing parenthesis
	// @Println(1,
	// )
	in := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("1"),
		tok(",", PR_SEPARATOR),
		tok("\n", PR_NEWLINE),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	// THEN the newline is removed along with the value delimiter
	exp := []token.Token{
		tok("@Print"),
		tok("(", PR_PARENTHESIS, PR_OPENER),
		tok("1"),
		tok(")", PR_PARENTHESIS, PR_CLOSER),
	}

	doTest(t, in, exp)
}
