package sanitise

import (
	"fmt"
	"testing"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/testutils"
	"github.com/stretchr/testify/require"
)

func doTest(t *testing.T, in, exps []token.Token) {

	var (
		tk     token.Token
		e      error
		stream = token.NewStream(in)
		acts   = []token.Token{}
	)

	for f := New(stream); f != nil; {
		if tk, f, e = f(); e != nil {
			require.NotNil(t, fmt.Sprintf("%+v", e))
		}

		acts = append(acts, tk)
	}

	testutils.RequireTokenSlice(t, exps, acts)
}

func tok(gen GenType, sub SubType, raw string) token.Tok {
	return token.Tok{
		Gen:    gen,
		Sub:    sub,
		RawStr: raw,
		ColEnd: len(raw),
	}
}

func Test1_1(t *testing.T) {

	// WHEN sanitising a statement containing redudant whitespace
	// @Println (  )
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_WHITESPACE, SUB_UNDEFINED, " "),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_WHITESPACE, SUB_UNDEFINED, "  "),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	// THEN the whitespace is removed
	exp := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test2_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after opening parenthesis
	// @Println(
	// )
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_1(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// @Println(1,
	// 1)
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed
	exp := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}

func Test3_2(t *testing.T) {

	// WHEN sanitising a spell call containing a newline after a value delimiter
	// AND the next line only contains the closing parenthesis
	// @Println(1,
	// )
	in := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_DELIMITER, SUB_VALUE_DELIM, ","),
		tok(GEN_TERMINATOR, SUB_NEWLINE, "\n"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	// THEN the newline is removed along with the value delimiter
	exp := []token.Token{
		tok(GEN_SPELL, SUB_UNDEFINED, "@Print"),
		tok(GEN_PARENTHESIS, SUB_PAREN_OPEN, "("),
		tok(GEN_LITERAL, SUB_NUMBER, "1"),
		tok(GEN_PARENTHESIS, SUB_PAREN_CLOSE, ")"),
	}

	doTest(t, in, exp)
}
